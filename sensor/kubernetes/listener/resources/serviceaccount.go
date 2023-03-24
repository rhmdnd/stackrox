package resources

import (
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/protoconv"
	"github.com/stackrox/rox/sensor/common/store/resolver"
	"github.com/stackrox/rox/sensor/kubernetes/eventpipeline/component"
	v1 "k8s.io/api/core/v1"
)

// serviceAccountDispatcher handles service account events
type serviceAccountDispatcher struct {
	serviceAccountStore *ServiceAccountStore
}

// newServiceAccountDispatcher creates and returns a new service account dispatcher.
func newServiceAccountDispatcher(serviceAccountStore *ServiceAccountStore) *serviceAccountDispatcher {
	return &serviceAccountDispatcher{
		serviceAccountStore: serviceAccountStore,
	}
}

// ProcessEvent processes a service account resource event, and returns the sensor events to emit in response.
func (s *serviceAccountDispatcher) ProcessEvent(obj, _ interface{}, action central.ResourceAction) *component.ResourceEvent {
	serviceAccount := obj.(*v1.ServiceAccount)

	var serviceAccountSecrets []string
	for _, secret := range serviceAccount.Secrets {
		serviceAccountSecrets = append(serviceAccountSecrets, secret.Name)
	}

	var serviceAccountImagePullSecrets []string
	for _, ipSecret := range serviceAccount.ImagePullSecrets {
		serviceAccountImagePullSecrets = append(serviceAccountImagePullSecrets, ipSecret.Name)
	}

	sa := &central.SensorEvent_ServiceAccount{
		ServiceAccount: &storage.ServiceAccount{
			Id:               string(serviceAccount.GetUID()),
			Name:             serviceAccount.GetName(),
			Namespace:        serviceAccount.GetNamespace(),
			CreatedAt:        protoconv.ConvertTimeToTimestamp(serviceAccount.GetCreationTimestamp().Time),
			AutomountToken:   true,
			Labels:           serviceAccount.GetLabels(),
			Annotations:      serviceAccount.GetAnnotations(),
			Secrets:          serviceAccountSecrets,
			ImagePullSecrets: serviceAccountImagePullSecrets,
		},
	}

	if serviceAccount.AutomountServiceAccountToken != nil {
		sa.ServiceAccount.AutomountToken = *serviceAccount.AutomountServiceAccountToken
	}

	var deploymentReference resolver.DeploymentReference
	if features.SourcedAutogeneratedIntegrations.Enabled() {
		switch action {
		case central.ResourceAction_REMOVE_RESOURCE:
			s.serviceAccountStore.Remove(sa.ServiceAccount)
		default:
			s.serviceAccountStore.Add(sa.ServiceAccount)
		}

		// Since SourcedAutogeneratedIntegrations influences deployment detection, with re-sync disabled we have to make
		// sure that deployments matching the ServiceAccount get reprocessed whenever a ServiceAccount is created / updated.
		deploymentReference = resolver.ResolveDeploymentsByMultipleServiceAccounts([]resolver.NamespaceServiceAccount{
			{
				Namespace:      sa.ServiceAccount.GetNamespace(),
				ServiceAccount: sa.ServiceAccount.GetName(),
			},
		})
	}

	events := []*central.SensorEvent{
		{
			Id:       string(serviceAccount.UID),
			Action:   action,
			Resource: sa,
		},
	}
	componentMessage := component.NewEvent(events...)
	if env.ResyncDisabled.BooleanSetting() && deploymentReference != nil {
		// Service Accounts can influence the `image` on AugmentedObject instance. Meaning the storage.Deployment object
		// won't be changed based on Service Account properties, but the alerts might. So the detection has to be forced.
		componentMessage.AddDeploymentReference(deploymentReference, central.ResourceAction_UPDATE_RESOURCE, true)
	}

	return componentMessage
}
