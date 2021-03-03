package core

import (
	"github.com/Alter17Ego/generic-app/errors/codederrors"
	"github.com/Alter17Ego/generic-http/responses/errorresp"
	commonSec "github.com/Alter17Ego/secuware/http/middlewares/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
	"net/http"
)

type PolicyMatcherStrategy int

const (
	MatchAny PolicyMatcherStrategy = iota
	MatchAll
)

type PolicyFunc func(claims jwt.MapClaims) (bool, error)

type Policies []PolicyFunc

func (p Policies) evaluateMatchAny(claims jwt.MapClaims) (bool, error) {
	for _, policy := range p {
		isOk, err := policy(claims)
		if err != nil {
			return false, err
		}
		if isOk {
			return true, nil
		}
	}
	return false, nil
}

func (p Policies) evaluateMatchAll(claims jwt.MapClaims) (bool, error) {
	for _, policy := range p {
		isOk, err := policy(claims)
		if err != nil {
			return false, err
		}
		if !isOk {
			return false, nil
		}
	}
	return true, nil
}

func (p Policies) Evaluate(strategy PolicyMatcherStrategy, claims jwt.MapClaims) (bool, error) {
	if strategy == MatchAll {
		return p.evaluateMatchAll(claims)
	} else if strategy == MatchAny {
		return p.evaluateMatchAny(claims)
	}
	return false, codederrors.New("err.evalute.policy", "Invalid strategy")
}

func RouteWithPolicy(strategy PolicyMatcherStrategy, policies Policies, route http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		claims, err := commonSec.GetDecodedAccessToken(request.Context())
		if err != nil {
			RenderErrorFromDecodedAccessToken(err, writer, request)
			return
		}
		isOk, err := policies.Evaluate(strategy, claims)
		if err != nil {
			switch e := err.(type) {
			case *codederrors.CodedError:
				render.Render(writer, request, errorresp.InternalServerError(e))
				return
			default:
				render.Render(writer, request, errorresp.InternalServerError(codederrors.Wrap(err, "err.unkown", "Unkown Error Ocurred")))
				return
			}
		}
		if !isOk {
			render.Render(writer, request, errorresp.Forbidden(ErrForbiddenAccess))
			return
		}
		route(writer, request)
	}
}
