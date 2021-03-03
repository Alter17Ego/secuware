package jwt

import (
	"net/http"
	"regexp"
	"strings"
)

func isUrlOnList(r *http.Request, urls ...string) (bool, error) {
	for _, url := range urls {
		if match, err := regexp.MatchString(url, r.RequestURI); err != nil {
			return false, err
		} else if match {
			return true, nil
		}
	}
	return false, nil
}

func getTokenFromHeader(r *http.Request) string {
	tokenHeader := r.Header.Get("Authorization")
	tokenStr := strings.Replace(tokenHeader, "Bearer ", "", 1)
	return tokenStr
}

func hasAuthorizationHeader(r *http.Request) bool {
	authorizationKey := r.Header.Get("Authorization")
	return len(strings.TrimSpace(authorizationKey)) > 0
}
