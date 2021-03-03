package resourcebased

type Scopes []string

func (s Scopes) MatchOneOf(checkerScopes Scopes) bool {
	bucket := map[string]int{}
	for _, scope := range s {
		bucket[scope] = 1
	}

	for _, scope := range checkerScopes {
		if bucket[scope] == 1 {
			return true
		}
	}
	return false
}

func (s Scopes) IncludedOn(checkerScopes Scopes) bool {
	bucket := map[string]int{}
	for _, scope := range s {
		bucket[scope] = 1
	}
	foundScopeCount := 0
	for _, scope := range checkerScopes {
		if bucket[scope] == 1 {
			foundScopeCount++
		}
	}
	return len(s) == foundScopeCount
}

type Resource struct {
	Name  string
	Scope Scopes
}
