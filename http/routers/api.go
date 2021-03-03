package routers

import (
	"github.com/Alter17Ego/secuware/http/routers/core"
	"github.com/Alter17Ego/secuware/http/routers/resourcebased"
	"github.com/Alter17Ego/secuware/http/routers/rolebased"
	"net/http"
)

func RouteWithClientRoles(clientName string, roles []string, routeFunc http.HandlerFunc) http.HandlerFunc {
	return rolebased.RouteWithClientRoles(clientName, roles, routeFunc)
}

func RouteWithCurrentSessionClientRoles(roles []string, routeFunc http.HandlerFunc) http.HandlerFunc {
	return rolebased.RouteWithCurrentSessionClientRoles(roles, routeFunc)
}

func RouteWithRealmRoles(roles []string, routeFunc http.HandlerFunc) http.HandlerFunc {
	return rolebased.RouteWithRealmRoles(roles, routeFunc)
}

func RouteWithResources(resources []resourcebased.Resource, routeFunc http.HandlerFunc) http.HandlerFunc {
	return resourcebased.RouteWithResource(resources, routeFunc)
}

func RouteWithScopedResources(resources []resourcebased.Resource, routeFunc http.HandlerFunc) http.HandlerFunc {
	return resourcebased.RouteWithScopeResources(resources, routeFunc)
}

func RouteWithPolicy(strategy core.PolicyMatcherStrategy, policies core.Policies, routeFunc http.HandlerFunc) http.HandlerFunc {
	return core.RouteWithPolicy(strategy, policies, routeFunc)
}

// Deprecated: Use RouteWithResource
func RouteWithNamedResources(resourceNames []string, routeFunc http.HandlerFunc) http.HandlerFunc {
	var resources []resourcebased.Resource
	for _, resource := range resourceNames {
		resources = append(resources, resourcebased.Resource{Name: resource})
	}
	return RouteWithResources(resources, routeFunc)
}
