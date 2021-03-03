package jwt

import (
	"context"
	"github.com/Alter17Ego/generic-app/errors/codederrors"
	goJWT "github.com/dgrijalva/jwt-go"
	"net/http"
)

const (
	AccessToken = "access_token"
)

func GetAccessToken(ctx context.Context) *goJWT.Token {
	return ctx.Value(AccessToken).(*goJWT.Token)
}

func GetDecodedAccessToken(ctx context.Context) (map[string]interface{}, error) {
	token := GetAccessToken(ctx)
	if token != nil {
		if claims, ok := token.Claims.(goJWT.MapClaims); ok && token.Valid {
			return claims, nil
		} else {
			codedErr := codederrors.Wrap(token.Claims.Valid(), "error.access.internalerr", "token was invalid")
			return nil, codedErr
		}
	}
	return nil, TokenNotfoundError
}

func AppendAccessToken(ctx context.Context, token *goJWT.Token) context.Context {
	return context.WithValue(ctx, AccessToken, token)
}

func AppendAccessTokenToReq(req *http.Request, token *goJWT.Token) *http.Request {
	ctx := AppendAccessToken(req.Context(), token)
	return req.WithContext(ctx)
}
