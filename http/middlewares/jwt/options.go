package jwt

import (
	"github.com/Alter17Ego/generic-app/errors/codederrors"
	"github.com/Alter17Ego/generic-http/responses/errorresp"
	"net/http"

	"github.com/go-chi/render"
)

type JwtValidationOptions struct {
	JwksUrl                  string
	WhiteListUrl             []string
	OnValidationErrorHandler JwtValidationErrorHandler
}

func DefaultValidationErrorHandler(w http.ResponseWriter, r *http.Request, err *codederrors.CodedError) {
	response := errorresp.Unauthorized(err)
	render.Render(w, r, response)
}

type JwtValidationOptionsBuilder struct {
	jwksUrl                  string
	whiteListUrl             []string
	onValidationErrorHandler JwtValidationErrorHandler
}

func (b *JwtValidationOptionsBuilder) WithDefaultOpts() *JwtValidationOptionsBuilder {
	b.onValidationErrorHandler = DefaultValidationErrorHandler
	return b
}

func (b *JwtValidationOptionsBuilder) SetJwksUrl(jwksUrl string) *JwtValidationOptionsBuilder {
	b.jwksUrl = jwksUrl
	return b
}

// deprecated
func (b *JwtValidationOptionsBuilder) AddWhitelist(urls ...string) *JwtValidationOptionsBuilder {
	return b.SetProtectedResources(urls...)
}

func (b *JwtValidationOptionsBuilder) SetProtectedResources(resourceUrls ...string) *JwtValidationOptionsBuilder {
	b.whiteListUrl = append(b.whiteListUrl, resourceUrls...)
	return b
}

func (b *JwtValidationOptionsBuilder) SetOnValidationErrorHandler(validationErrorHandler JwtValidationErrorHandler) *JwtValidationOptionsBuilder {
	b.onValidationErrorHandler = validationErrorHandler
	return b
}

func (b *JwtValidationOptionsBuilder) Build() *JwtValidationOptions {
	options := JwtValidationOptions{
		JwksUrl:                  b.jwksUrl,
		WhiteListUrl:             b.whiteListUrl,
		OnValidationErrorHandler: b.onValidationErrorHandler,
	}
	return &options
}

func CreateJwtValidationOptionsBuilder() *JwtValidationOptionsBuilder {
	return &JwtValidationOptionsBuilder{}
}
