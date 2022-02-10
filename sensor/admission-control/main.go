package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/pkg/clientconn"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/devmode"
	"github.com/stackrox/rox/pkg/env"
	pkgGRPC "github.com/stackrox/rox/pkg/grpc"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/mtls"
	"github.com/stackrox/rox/pkg/mtls/verifier"
	"github.com/stackrox/rox/pkg/namespaces"
	"github.com/stackrox/rox/pkg/safe"
	"github.com/stackrox/rox/pkg/satoken"
	"github.com/stackrox/rox/pkg/utils"
	"github.com/stackrox/rox/pkg/version"
	"github.com/stackrox/rox/sensor/admission-control/alerts"
	"github.com/stackrox/rox/sensor/admission-control/manager"
	"github.com/stackrox/rox/sensor/admission-control/service"
	"github.com/stackrox/rox/sensor/admission-control/settingswatch"
	"golang.org/x/sys/unix"
)

const (
	webhookEndpoint = ":8443"

	internalGracePeriod = 15 * time.Second
)

var (
	log = logging.LoggerForModule()
)

func main() {
	log.Infof("StackRox Sensor Admission Control Service, version %s", version.GetMainVersion())

	utils.Must(mainCmd())
}

func mainCmd() error {
	devmode.StartOnDevBuilds("bin/admission-control")

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, unix.SIGTERM, unix.SIGINT)

	namespace, err := satoken.LoadNamespaceFromFile()
	if err != nil {
		namespace = namespaces.StackRox
		log.Warnf("Failed to read namespace from service account secret: %v. Assuming default '%s' namespace...", err, namespace)
	}

	if err := safe.RunE(func() error {
		if err := configureCA(); err != nil {
			return err
		}
		if err := configureCerts(namespace); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Errorf("Failed to configure certificates: %v. Connection to sensor might fail.", err)
	}

	// Note that the following call returns immediately (connecting happens in the background), hence this does not
	// delay readiness of the admission-control service even if sensor is unavailable.
	sensorConn, err := clientconn.AuthenticatedGRPCConnection(env.SensorEndpoint.Setting(), mtls.SensorSubject)
	if err != nil {
		log.Errorf("Could not establish a gRPC connection to Sensor: %v. Some features, including recording"+
			" violations generated by admission control, will not work.", err)
	}

	mgr := manager.New(sensorConn, namespace)
	if err := mgr.Start(); err != nil {
		return errors.Wrap(err, "starting admission control manager")
	}

	if err := settingswatch.WatchK8sForSettingsUpdatesAsync(mgr.Stopped(), mgr.SettingsUpdateC(), namespace); err != nil {
		log.Errorf("Could not watch Kubernetes for settings updates: %v. Functionality might be impacted", err)
	}
	if err := settingswatch.WatchMountPathForSettingsUpdateAsync(mgr.Stopped(), mgr.SettingsUpdateC()); err != nil {
		log.Errorf("Could not watch mount path for settings updates: %v. Functionality might be impacted", err)
	}
	if err := settingswatch.RunSettingsPersister(mgr); err != nil {
		log.Errorf("Could not run settings persister: %v. Admission control service might take longer to become ready after container restarts", err)
	}
	if sensorConn != nil {
		settingswatch.WatchSensorMessagePush(mgr, sensorConn)
		alerts.NewAlertSender(sensorConn, mgr.Alerts()).Start(concurrency.AsContext(mgr.Stopped()))
	}

	serverConfig := pkgGRPC.Config{
		Endpoints: []*pkgGRPC.EndpointConfig{
			{
				ListenEndpoint: webhookEndpoint,
				TLS:            verifier.NonCA{},
				ServeHTTP:      true,
				NoHTTP2:        true,
			},
		},
	}

	apiServer := pkgGRPC.NewAPI(serverConfig)
	apiServer.Register(service.New(mgr))

	apiServer.Start()

	// Graceful shutdown logic:
	// Upon first SIGTERM, keep running normally until either the internal grace period (20s) passes, or another
	// SIGTERM is received. However, mark the container as not ready immediately, to make sure we no longer receive
	// new requests once the API server has picked up the non-readiness.
	sigTermCounter := 0
	var gracePeriodTimer <-chan time.Time
	for {
		select {
		case sig := <-sigC:
			log.Infof("Received signal %v", sig)
			if sig == unix.SIGTERM {
				sigTermCounter++

				if sigTermCounter == 1 {
					log.Infof("First SIGTERM. Marking as not ready, will exit in %v", internalGracePeriod)
					mgr.Stop()
					gracePeriodTimer = time.After(internalGracePeriod)
				} else {
					log.Info("Second SIGTERM. Exiting immediately.")
					return nil
				}
			} else {
				log.Info("Received signal other than SIGTERM. Exiting immediately.")
				return nil
			}

		case <-gracePeriodTimer:
			log.Infof("Grace period of %v has expired. Exiting ...", internalGracePeriod)
			return nil
		}
	}
}
