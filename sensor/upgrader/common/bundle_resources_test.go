package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TestEnsureBundleResourcesTypesAreCorrect exists to make accidental modifications of BundleResourcesTypes less likely.
// There is of course no protection against somebody intentionally modifying both lists.
func TestEnsureBundleResourcesTypesAreCorrect(t *testing.T) {
	t.Parallel()

	assert.ElementsMatch(t, BundleResourceTypes, []schema.GroupVersionKind{
		{Version: "v1", Kind: "Service"},
		{Version: "v1", Kind: "ServiceAccount"},
		{Version: "v1", Kind: "Secret"},
		{Version: "v1", Kind: "ConfigMap"},
		{Group: "apps", Version: "v1beta2", Kind: "DaemonSet"},
		{Group: "extensions", Version: "v1beta1", Kind: "Deployment"},
		{Group: "extensions", Version: "v1beta1", Kind: "PodSecurityPolicy"},
		{Group: "admissionregistration.k8s.io", Version: "v1beta1", Kind: "ValidatingWebhookConfiguration"},
		{Group: "networking.k8s.io", Version: "v1", Kind: "NetworkPolicy"},
		{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "ClusterRole"},
		{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "ClusterRoleBinding"},
		{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "Role"},
		{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "RoleBinding"},
		{Group: "security.openshift.io", Version: "v1", Kind: "SecurityContextConstraints"},
	})
}
