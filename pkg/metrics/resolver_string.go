// Code generated by "stringer -type=Resolver"; DO NOT EDIT.

package metrics

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Cluster-0]
	_ = x[Compliance-1]
	_ = x[ComlianceControl-2]
	_ = x[CVEs-3]
	_ = x[Deployments-4]
	_ = x[Groups-5]
	_ = x[Images-6]
	_ = x[ImageComponents-7]
	_ = x[K8sRoles-8]
	_ = x[Namespaces-9]
	_ = x[Nodes-10]
	_ = x[Notifiers-11]
	_ = x[PermissionSets-12]
	_ = x[Policies-13]
	_ = x[Roles-14]
	_ = x[Root-15]
	_ = x[Secrets-16]
	_ = x[ServiceAccounts-17]
	_ = x[Subjects-18]
	_ = x[Tokens-19]
	_ = x[Violations-20]
	_ = x[Pods-21]
	_ = x[ContainerInstances-22]
	_ = x[ImageCVEs-23]
	_ = x[NodeCVEs-24]
	_ = x[ClusterCVEs-25]
	_ = x[NodeComponents-26]
	_ = x[ImageCVECore-27]
	_ = x[PlatformCVECore-28]
	_ = x[NodeCVECore-29]
}

const _Resolver_name = "ClusterComplianceComlianceControlCVEsDeploymentsGroupsImagesImageComponentsK8sRolesNamespacesNodesNotifiersPermissionSetsPoliciesRolesRootSecretsServiceAccountsSubjectsTokensViolationsPodsContainerInstancesImageCVEsNodeCVEsClusterCVEsNodeComponentsImageCVECorePlatformCVECoreNodeCVECore"

var _Resolver_index = [...]uint16{0, 7, 17, 33, 37, 48, 54, 60, 75, 83, 93, 98, 107, 121, 129, 134, 138, 145, 160, 168, 174, 184, 188, 206, 215, 223, 234, 248, 260, 275, 286}

func (i Resolver) String() string {
	if i < 0 || i >= Resolver(len(_Resolver_index)-1) {
		return "Resolver(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Resolver_name[_Resolver_index[i]:_Resolver_index[i+1]]
}
