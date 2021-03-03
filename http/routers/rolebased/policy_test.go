package rolebased_test

import (
	"github.com/Alter17Ego/secuware/http/routers/rolebased"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasClientRole__ShouldReturnTrueWhenRoleMatches(t *testing.T) {
	claims := jwt.MapClaims{
		"resource_access": map[string]interface{}{
			"secuware-web": map[string]interface{}{
				"roles": []interface{}{
					"secuware-user",
				},
			},
		},
	}
	checker := rolebased.HasClientRoles("secuware-web", "secuware-user")
	found, err := checker(claims)
	if err != nil {
		panic(err)
	}
	assert.True(t, found)
}

func TestHasClientRole__ShouldReturnFalseWhenRoleNotMatches(t *testing.T) {
	claims := jwt.MapClaims{
		"resource_access": map[string]interface{}{
			"secuware-web": map[string]interface{}{
				"roles": []interface{}{
					"secuware-admin",
				},
			},
		},
	}
	checker := rolebased.HasClientRoles("secuware-web", "secuware-user")
	found, err := checker(claims)
	if err != nil {
		panic(err)
	}
	assert.False(t, found)
}

func TestHasClientRole__ShouldReturnFalseWhenResourceNotMatches(t *testing.T) {
	claims := jwt.MapClaims{
		"resource_access": map[string]interface{}{
			"secuware-mobile": map[string]interface{}{
				"roles": []interface{}{
					"secuware-user",
				},
			},
		},
	}
	checker := rolebased.HasClientRoles("secuware-web", "secuware-user")
	found, err := checker(claims)
	if err != nil {
		panic(err)
	}
	assert.False(t, found)
}

func TestHasCurrentSessionClientRole__ShouldReturnTrueWhenRoleMatches(t *testing.T) {
	claims := jwt.MapClaims{
		"azp": "secuware-web",
		"resource_access": map[string]interface{}{
			"secuware-web": map[string]interface{}{
				"roles": []interface{}{
					"secuware-user",
				},
			},
		},
	}
	checker := rolebased.HasCurrentSessionClientRoles("secuware-user")
	found, err := checker(claims)
	if err != nil {
		panic(err)
	}
	assert.True(t, found)
}

func TestHasCurrentSessionClientRole__ShouldReturnFalseWhenRoleNotMatches(t *testing.T) {
	claims := jwt.MapClaims{
		"azp": "secuware-web",
		"resource_access": map[string]interface{}{
			"secuware-web": map[string]interface{}{
				"roles": []interface{}{
					"secuware-admin",
				},
			},
		},
	}
	checker := rolebased.HasCurrentSessionClientRoles("secuware-user")
	found, err := checker(claims)
	if err != nil {
		panic(err)
	}
	assert.False(t, found)
}

func TestCurrentSessionClient__ShouldReturnFalseWhenResourceNotMatches(t *testing.T) {
	claims := jwt.MapClaims{
		"azp": "secuware-web",
		"resource_access": map[string]interface{}{
			"secuware-mobile": map[string]interface{}{
				"roles": []interface{}{
					"secuware-user",
				},
			},
		},
	}
	checker := rolebased.HasCurrentSessionClientRoles("secuware-user")
	found, err := checker(claims)
	if err != nil {
		t.Error(err)
	}
	assert.False(t, found)
}

func TestHasRealmRoles__ShouldReturnTrueWhenRoleMatches(t *testing.T) {
	claims := jwt.MapClaims{
		"realm_access": map[string]interface{}{
			"roles": []interface{}{
				"update_user",
			},
		},
	}

	checker := rolebased.HasRealmRoles("update_user")
	ok, err := checker(claims)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, ok)
}

func TestHasRealmRoles__ShouldReturnTrueWhenMultipleRolesMatches(t *testing.T) {
	claims := jwt.MapClaims{
		"realm_access": map[string]interface{}{
			"roles": []interface{}{
				"update_user",
				"delete_user",
			},
		},
	}

	checker := rolebased.HasRealmRoles("update_user", "delete_user")
	ok, err := checker(claims)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, ok)
}

func TestHasRealmRoles__ShouldReturnFalseWhenNoRoleMatches(t *testing.T) {
	claims := jwt.MapClaims{
		"realm_access": map[string]interface{}{
			"roles": []interface{}{
				"insert_user",
			},
		},
	}

	checker := rolebased.HasRealmRoles("update_user")
	ok, err := checker(claims)
	if err != nil {
		t.Error(err)
	}
	assert.False(t, ok)
}
