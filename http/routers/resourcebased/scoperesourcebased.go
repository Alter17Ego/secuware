package resourcebased

import (
	"github.com/Alter17Ego/secuware/http/routers/core"
	"net/http"
)

func RouteWithScopeResources(resources []Resource, routeFunc http.HandlerFunc) http.HandlerFunc {
	return core.RouteWithPolicy(core.MatchAll, core.Policies{HasScopedResources(resources...)}, routeFunc)
}
