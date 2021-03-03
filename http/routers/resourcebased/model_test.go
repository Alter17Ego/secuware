package resourcebased_test

import (
	"github.com/Alter17Ego/secuware/http/routers/resourcebased"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScopes_MatchOneOfShouldBeTrue(t *testing.T) {
	scopes := resourcebased.Scopes{"A", "B", "C"}
	checkerScopes := resourcebased.Scopes{"Z", "B", "X"}
	result := scopes.MatchOneOf(checkerScopes)
	assert.True(t, result)
}

func TestScopes_MatchOneOfShouldBeFalse(t *testing.T) {
	scopes := resourcebased.Scopes{"A", "B", "C"}
	checkerScopes := resourcebased.Scopes{"X", "Y", "Z"}
	result := scopes.MatchOneOf(checkerScopes)
	assert.False(t, result)
}

func TestScopes_IncludedOnShouldBeTrue(t *testing.T) {
	scopes := resourcebased.Scopes{"A", "B", "C"}
	checkerScopes := resourcebased.Scopes{"A", "B", "C"}
	result := scopes.IncludedOn(checkerScopes)
	assert.True(t, result)
}

func TestScopes_IncludedOnShouldBeFalse(t *testing.T) {
	scopes := resourcebased.Scopes{"A", "B", "C"}
	checkerScopes := resourcebased.Scopes{"A", "B"}
	result := scopes.IncludedOn(checkerScopes)
	assert.False(t, result)
}
