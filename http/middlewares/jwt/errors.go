package jwt

import "github.com/Alter17Ego/generic-app/errors/codederrors"

var KeyIdNotFoundError error = codederrors.New("token.validation.error", "Key Id not found")
var KeyIdNotMatchError error = codederrors.New("token.validation.error", "Key Id not match")
var TokenNotfoundError error = codederrors.New("token.error.notfound", "No token was provided")
var TokenInvalidError error = codederrors.New("token.error.notfound", "No token was provided")
