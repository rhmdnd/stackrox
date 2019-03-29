package connection

import (
	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/sensor/service/pipeline"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/sync"
)

type manager struct {
	connectionsByClusterID      map[string]*sensorConnection
	connectionsByClusterIDMutex sync.RWMutex
}

func newManager() *manager {
	return &manager{
		connectionsByClusterID: make(map[string]*sensorConnection),
	}
}

func (m *manager) GetConnection(clusterID string) SensorConnection {
	m.connectionsByClusterIDMutex.RLock()
	defer m.connectionsByClusterIDMutex.RUnlock()

	conn := m.connectionsByClusterID[clusterID]
	if conn == nil {
		return nil
	}
	return conn
}

func (m *manager) HandleConnection(clusterID string, pf pipeline.Factory, server central.SensorService_CommunicateServer, recorder checkInRecorder) error {
	conn, err := newConnection(clusterID, pf, recorder)

	if err != nil {
		return errors.Wrap(err, "creating sensor connection")
	}

	var oldConnection *sensorConnection
	concurrency.WithLock(&m.connectionsByClusterIDMutex, func() {
		oldConnection = m.connectionsByClusterID[clusterID]
		m.connectionsByClusterID[clusterID] = conn
	})

	if oldConnection != nil {
		oldConnection.Terminate(errors.New("replaced by new connection"))
	}

	return conn.Run(server)
}
