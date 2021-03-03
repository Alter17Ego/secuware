package rolebased

import (
	"github.com/Alter17Ego/secuware/http/routers/core"
	"net/http"
)

func RouteWithCurrentSessionClientRoles(roles []string, route http.HandlerFunc) http.HandlerFunc {
	return core.RouteWithPolicy(core.MatchAll, core.Policies{HasCurrentSessionClientRoles(roles...)}, route)
}

func RouteWithClientRoles(clientName string, roles []string, route http.HandlerFunc) http.HandlerFunc {
	return core.RouteWithPolicy(core.MatchAll, core.Policies{HasClientRoles(clientName, roles...)}, route)
}

func RouteWithRealmRoles(roles []string, route http.HandlerFunc) http.HandlerFunc {
	return core.RouteWithPolicy(core.MatchAll, core.Policies{HasRealmRoles(roles...)}, route)
}
