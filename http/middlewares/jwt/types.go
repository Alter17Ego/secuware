package jwt

import (
	"github.com/Alter17Ego/generic-app/errors/codederrors"
	"net/http"
)

type Middleware func(next http.Handler) http.Handler
type JwtValidationErrorHandler func(w http.ResponseWriter, r *http.Request, err *codederrors.CodedError)
