package role

import (
	"github.com/stackrox/rox/central/role/resources"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/grpc/authn"
	"github.com/stackrox/rox/pkg/set"
)

// All builtin, immutable role names are declared in the block below.
const (
	// Admin is a role that's, well, authorized to do anything.
	Admin = "Admin"

	// Analyst is a role that has read access to all resources.
	Analyst = "Analyst"

	// None role has no access.
	None = authn.NoneRole

	// ContinuousIntegration is for CI pipelines.
	ContinuousIntegration = "Continuous Integration"

	// SensorCreator is a role that has the minimal privileges required to create a sensor.
	SensorCreator = "Sensor Creator"

	// VulnMgmtApprover is a role that has the minimal privileges required to approve vulnerability deferrals or false positive requests.
	VulnMgmtApprover = "Vulnerability Management Approver"

	// VulnMgmtRequester is a role that has the minimal privileges required to request vulnerability deferrals or false positives.
	VulnMgmtRequester = "Vulnerability Management Requester"

	// VulnReporter is a role that has the minimal privileges required to create and manage vulnerability reporting configurations.
	VulnReporter = "Vulnerability Report Creator"

	// ScopeManager is a role that has the minimal privileges to view and modify scopes for use in access control, vulnerability reporting etc.
	ScopeManager = "Scope Manager"
)

// DefaultRoleNames is a string set containing the names of all default (built-in) Roles.
var DefaultRoleNames = set.NewStringSet(Admin, Analyst, None, ContinuousIntegration, ScopeManager, SensorCreator, VulnMgmtApprover, VulnMgmtRequester, VulnReporter)

// AccessScopeExcludeAll has empty rules and hence excludes all
// scoped resources. Global resources must be unaffected.
var AccessScopeExcludeAll = &storage.SimpleAccessScope{
	Id:          EnsureValidAccessScopeID("denyall"),
	Name:        "Deny All",
	Description: "No access to scoped resources",
	Rules:       &storage.SimpleAccessScope_Rules{},
}

// IsDefaultRoleName checks if a given role name corresponds to a default role.
func IsDefaultRoleName(name string) bool {
	return DefaultRoleNames.Contains(name)
}

// IsDefaultAccessScope checks if a given access scope name corresponds to the
// default access scope.
func IsDefaultAccessScope(name string) bool {
	return AccessScopeExcludeAll.GetName() == name
}

// GetAnalystPermissions returns permissions for `Analyst` role.
func GetAnalystPermissions() []permissions.ResourceWithAccess {
	resourceToAccess := resources.AllResourcesViewPermissions()
	for i, resourceWithAccess := range resourceToAccess {
		if resourceWithAccess.Resource.GetResource() == resources.DebugLogs.GetResource() {
			return append(resourceToAccess[:i], resourceToAccess[i+1:]...)
		}
	}
	panic("DebugLogs resource was not found amongst all resources.")
}
