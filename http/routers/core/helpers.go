package core

import (
	"github.com/Alter17Ego/generic-app/errors/codederrors"
	"github.com/Alter17Ego/generic-http/responses/errorresp"
	"github.com/Alter17Ego/secuware/http/middlewares/jwt"
	"github.com/go-chi/render"
	"net/http"
)

func RenderErrorFromDecodedAccessToken(err error, w http.ResponseWriter, r *http.Request) {
	switch e := err.(type) {
	case *codederrors.CodedError:
		if e.Error() == jwt.TokenNotfoundError.Error() {
			render.Render(w, r, errorresp.Forbidden(e))
			return
		}
		render.Render(w, r, errorresp.InternalServerError(e))
	default:
		codedErr := codederrors.Wrap(err, "error.access.internalerr", "token was invalid")
		render.Render(w, r, errorresp.InternalServerError(codedErr))
	}
}
