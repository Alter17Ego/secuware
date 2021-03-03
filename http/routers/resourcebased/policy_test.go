package resourcebased_test

import (
	"encoding/json"
	"github.com/Alter17Ego/secuware/http/routers/resourcebased"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_HasResource_ShouldReturnTrue_WhenResourceAreExists(t *testing.T) {

	accessTokenBody := `{
		"exp": 1595578496,
		"iat": 1595578196,
		"auth_time": 1595577597,
		"jti": "356c7b72-52f3-456e-a4b4-089239e13788",
		"iss": "https://dev-ctricks.com/auth/realms/ct-devs",
		"aud": "k-chat-app",
		"sub": "25892586-80a9-4b1b-b7e2-8dbc61f7c63b",
		"typ": "Bearer",
		"azp": "k-chat-app",
		"session_state": "0b9b9e26-6332-4b71-be37-6c284ce09d42",
		"acr": "0",
		"realm_access": {
		  "roles": [
			"offline_access",
			"uma_authorization"
		  ]
		},
		"resource_access": {
		  "account": {
			"roles": [
			  "manage-account",
			  "manage-account-links",
			  "view-profile"
			]
		  }
		},
		"authorization": {
		  "permissions": [
			{
			  "rsid": "60bf1a91-b4cd-4011-81ca-60b02ad154d8",
			  "rsname": "Default Resource"
			},
			{
			  "rsid": "de8ef7c5-0828-4990-9cce-be11f6c64c55",
			  "rsname": "book"
			}
		  ]
		},
		"scope": "openid email profile",
		"email_verified": true,
		"preferred_username": "its.xan",
		"email": "its.xan@outlook.com"
	  }`

	var claims jwt.MapClaims

	if err := json.Unmarshal([]byte(accessTokenBody), &claims); err != nil {
		t.Error(err)
	}

	checker := resourcebased.HasResources(resourcebased.Resource{Name: "book"})

	isOk, err := checker(claims)

	if err != nil {
		t.Error(err)
	}

	assert.True(t, isOk, "checker should return true")
}
