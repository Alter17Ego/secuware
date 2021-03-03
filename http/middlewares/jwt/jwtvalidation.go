package jwt

import (
	"github.com/Alter17Ego/generic-app/errors/codederrors"
	"net/http"

	goJWT "github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

func createTokenValidator(keySet *jwk.Set) func(*goJWT.Token) (interface{}, error) {
	return func(token *goJWT.Token) (interface{}, error) {
		keyID, ok := token.Header["kid"].(string)
		if !ok {
			return nil, KeyIdNotFoundError
		}

		if key := keySet.LookupKeyID(keyID); len(key) == 1 {
			return key[0].Materialize()
		}
		return nil, KeyIdNotMatchError
	}
}

func JwtValidationJwks(options *JwtValidationOptions) Middleware {
	keySet, err := jwk.FetchHTTP(options.JwksUrl)
	tokenValidation := createTokenValidator(keySet)
	if err != nil {
		panic(err)
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			urlMatch, err := isUrlOnList(r, options.WhiteListUrl...)
			if err != nil {
				panic(err)
			}
			if !urlMatch {
				next.ServeHTTP(w, r)
				return
			}
			if !hasAuthorizationHeader(r) {
				options.OnValidationErrorHandler(w, r, TokenNotfoundError.(*codederrors.CodedError))
				return
			}
			tokenStr := getTokenFromHeader(r)

			token, err := goJWT.Parse(tokenStr, tokenValidation)
			if err != nil {
				if codedErr, ok := err.(*codederrors.CodedError); ok {
					options.OnValidationErrorHandler(w, r, codedErr)
				} else {
					codedErr := codederrors.Wrap(err, "token.validation.err", err.Error())
					options.OnValidationErrorHandler(w, r, codedErr)
				}

			} else {
				next.ServeHTTP(w, AppendAccessTokenToReq(r, token))
			}
		})
	}

}
