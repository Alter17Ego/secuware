package resourcebased

import (
	"github.com/Alter17Ego/secuware/http/routers/core"
	"github.com/dgrijalva/jwt-go"
)

func HasScopedResources(targetResources ...Resource) core.PolicyFunc {
	return func(claims jwt.MapClaims) (bool, error) {
		authorizations, ok := claims["authorization"].(map[string]interface{})
		if !ok {
			return false, nil
		}
		permissions, ok := authorizations["permissions"].([]interface{})

		if !ok {
			return false, nil
		}
		resourceNamesMap := make(map[string]*Resource)
		for _, resource := range targetResources {
			resourceNamesMap[resource.Name] = &resource
		}
		for _, permission := range permissions {
			permissionAsMap := permission.(map[string]interface{})
			resourceName := permissionAsMap["rsname"].(string)
			if target := resourceNamesMap[resourceName]; target != nil {
				scopes := permissionAsMap["scopes"].(Scopes)
				if target.Scope.IncludedOn(scopes) {
					return true, nil
				}
				return false, nil
			}
		}
		return false, nil
	}
}

func HasResources(resources ...Resource) core.PolicyFunc {
	return func(claims jwt.MapClaims) (bool, error) {
		authorizations, ok := claims["authorization"].(map[string]interface{})
		if !ok {
			return false, nil
		}
		permissions, ok := authorizations["permissions"].([]interface{})
		if !ok {
			return false, nil
		}
		resourceNamesMap := make(map[string]int)
		for _, resource := range resources {
			resourceNamesMap[resource.Name] = 1
		}
		for _, permission := range permissions {
			permissionAsMap := permission.(map[string]interface{})
			resourceName := permissionAsMap["rsname"].(string)
			if resourceNamesMap[resourceName] == 1 {
				return true, nil
			}
		}
		return false, nil
	}
}
