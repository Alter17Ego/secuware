package resourcebased

import (
	"github.com/Alter17Ego/secuware/http/routers/core"
	"net/http"
)

func RouteWithResource(resources []Resource, routeFunc http.HandlerFunc) http.HandlerFunc {
	return core.RouteWithPolicy(core.MatchAll, core.Policies{HasResources(resources...)}, routeFunc)
}
