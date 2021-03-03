package rolebased

import (
	"github.com/Alter17Ego/secuware/http/routers/core"
	"github.com/dgrijalva/jwt-go"
)

func HasClientRoles(client string, targetClientRoles ...string) core.PolicyFunc {
	return func(claims jwt.MapClaims) (bool, error) {
		resourceAccess, ok := claims["resource_access"].(map[string]interface{})
		if !ok {
			return false, nil
		}
		clientResource, ok := resourceAccess[client].(map[string]interface{})
		if !ok {
			return false, nil
		}
		clientRoles := clientResource["roles"].([]interface{})
		clientRoleMap := make(map[string]int)
		for _, role := range clientRoles {
			clientRoleMap[role.(string)] = 1
		}
		for _, role := range targetClientRoles {
			if clientRoleMap[role] == 1 {
				return true, nil
			}
		}
		return false, nil
	}
}

func HasCurrentSessionClientRoles(targetClientRoles ...string) core.PolicyFunc {
	return func(claims jwt.MapClaims) (bool, error) {
		resourceAccess, ok := claims["resource_access"].(map[string]interface{})
		if !ok {
			return false, nil
		}
		client := claims["azp"].(string)
		clientResource, ok := resourceAccess[client].(map[string]interface{})
		if !ok {
			return false, nil
		}
		clientRoles := clientResource["roles"].([]interface{})
		clientRoleMap := make(map[string]int)
		for _, role := range clientRoles {
			clientRoleMap[role.(string)] = 1
		}
		for _, role := range targetClientRoles {
			if clientRoleMap[role] == 1 {
				return true, nil
			}
		}
		return false, nil
	}
}

func HasRealmRoles(targetRealmRoles ...string) core.PolicyFunc {
	return func(claims jwt.MapClaims) (bool, error) {

		resourceAccess, ok := claims["realm_access"].(map[string]interface{})
		if !ok {
			return false, nil
		}
		clientRoles := resourceAccess["roles"].([]interface{})
		clientRoleMap := make(map[string]int)
		for _, role := range clientRoles {
			clientRoleMap[role.(string)] = 1
		}
		for _, role := range targetRealmRoles {
			if clientRoleMap[role] == 1 {
				return true, nil
			}
		}
		return false, nil
	}
}
