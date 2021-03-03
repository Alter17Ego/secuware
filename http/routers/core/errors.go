package core

import "github.com/Alter17Ego/generic-app/errors/codederrors"

var ErrForbiddenAccess = codederrors.New("err.access.forbidden", "you don't have permission to access this url or method")
var ErrClientRolesNotFound = codederrors.New("err.client.roles.notfound", "client roles not found in this token")
var ErrNoAuthorizationInToken = codederrors.New("err.access.noauthorization", "no authorizations provided in the access_token")
var ErrNoPermissionInToken = codederrors.New("err.access.nopermission", "no permissions provided in the access_token")
