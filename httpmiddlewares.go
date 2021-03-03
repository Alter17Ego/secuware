package secuware

import (
	"github.com/Alter17Ego/secuware/http/middlewares/headers"
	"github.com/Alter17Ego/secuware/http/middlewares/jwt"
	"github.com/unrolled/secure"
	"net/http"
)

type chiMiddleware func(http.Handler) http.Handler

func SecurityHeaderMiddleware(options secure.Options) chiMiddleware {
	return headers.NewSecurityMiddleware(options).Handler
}

func JwksValidationMiddleware(options *jwt.JwtValidationOptions) jwt.Middleware {
	return jwt.JwtValidationJwks(options)
}

func JwksValidationMiddlewareOpts() *jwt.JwtValidationOptionsBuilder {
	return jwt.CreateJwtValidationOptionsBuilder()
}

func DefaultSecurityMiddleware() chiMiddleware {
	return headers.NewDefaultSecurityMiddleware().Handler
}
