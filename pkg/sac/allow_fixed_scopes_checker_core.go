package sac

import (
	"context"

	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/sac/effectiveaccessscope"
)

type scopeKeySet map[ScopeKey]struct{}

type allowFixedScopesCheckerCore []scopeKeySet

// AllowFixedScopes returns a scope checker core that allows those scopes that
// are in the cross product of all individual scope key lists. I.e.,
// AllowFixedScopes(
//   AccessModeScopeKeys(storage.Access_READ, storage.Access_READ_WRITE),
//   ResourceScopeKeys(resources.CLUSTER),
// )
// returns a scope checker core that allows read and write access to all cluster resources.
func AllowFixedScopes(keyLists ...[]ScopeKey) ScopeCheckerCore {
	sets := make(allowFixedScopesCheckerCore, len(keyLists))
	for i, keyList := range keyLists {
		set := make(map[ScopeKey]struct{}, len(keyList))
		for _, key := range keyList {
			set[key] = struct{}{}
		}
		sets[i] = set
	}
	return sets
}

func (c allowFixedScopesCheckerCore) TryAllowed() TryAllowedResult {
	if len(c) == 0 {
		return Allow
	}
	return Deny
}

func (c allowFixedScopesCheckerCore) PerformChecks(_ context.Context) error {
	return nil
}

func (c allowFixedScopesCheckerCore) SubScopeChecker(key ScopeKey) ScopeCheckerCore {
	if len(c) == 0 {
		return c
	}
	if _, ok := c[0][key]; ok {
		return c[1:]
	}
	return denyAllScopeCheckerCore
}

func (c allowFixedScopesCheckerCore) EffectiveAccessScope(resource permissions.ResourceWithAccess) (*effectiveaccessscope.ScopeTree, error) {
	if len(c) == 0 {
		return effectiveaccessscope.UnrestrictedEffectiveAccessScope(), nil
	}
	// Build the effective access scope tree based on the top level ScopeKey type
	// TODO (ROX-9981): Change the structure of allowFixedScopesCheckerCore to not rely on magic index values
	for key := range c[0] {
		switch key.(type) {
		case AccessModeScopeKey:
			return c.getAccessModeEffectiveAccessScope(resource)
		case ResourceScopeKey:
			return c.getResourceEffectiveAccessScope(resource)
		case ClusterScopeKey:
			return c.getClusterEffectiveAccessScope()
		}
		break
	}
	return effectiveaccessscope.DenyAllEffectiveAccessScope(), nil
}

func (c allowFixedScopesCheckerCore) getAccessModeEffectiveAccessScope(resource permissions.ResourceWithAccess) (*effectiveaccessscope.ScopeTree, error) {
	if len(c) == 0 {
		return effectiveaccessscope.UnrestrictedEffectiveAccessScope(), nil
	}
	_, accessAllowed := c[0][AccessModeScopeKey(resource.Access)]
	if !accessAllowed {
		return effectiveaccessscope.DenyAllEffectiveAccessScope(), nil
	}
	return c.next().getResourceEffectiveAccessScope(resource)
}

func (c allowFixedScopesCheckerCore) getResourceEffectiveAccessScope(resource permissions.ResourceWithAccess) (*effectiveaccessscope.ScopeTree, error) {
	if len(c) == 0 {
		return effectiveaccessscope.UnrestrictedEffectiveAccessScope(), nil
	}
	_, resourceAllowed := c[0][ResourceScopeKey(resource.Resource.GetResource())]
	if !resourceAllowed {
		return effectiveaccessscope.DenyAllEffectiveAccessScope(), nil
	}
	return c.next().getClusterEffectiveAccessScope()
}

func (c allowFixedScopesCheckerCore) getClusterEffectiveAccessScope() (*effectiveaccessscope.ScopeTree, error) {
	if len(c) == 0 {
		return effectiveaccessscope.UnrestrictedEffectiveAccessScope(), nil
	}
	clusterIDs := make([]string, 0, len(c[0]))
	for clusterID := range c[0] {
		clusterIDs = append(clusterIDs, clusterID.String())
	}
	if len(c) == 1 {
		return effectiveaccessscope.FromClustersAndNamespacesMap(clusterIDs, nil), nil
	}
	namespaces := make([]string, 0, len(c[1]))
	for namespace := range c[1] {
		namespaces = append(namespaces, namespace.String())
	}
	clusterNamespaceMap := make(map[string][]string, 0)
	for clusterIx := range clusterIDs {
		clusterID := clusterIDs[clusterIx]
		clusterNamespaceMap[clusterID] = namespaces
	}
	return effectiveaccessscope.FromClustersAndNamespacesMap(nil, clusterNamespaceMap), nil
}

func (c allowFixedScopesCheckerCore) next() allowFixedScopesCheckerCore {
	return c[1:]
}
