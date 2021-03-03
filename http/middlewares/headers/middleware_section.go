package headers

import (
	"github.com/unrolled/secure"
)

var DefaultSecureOption = secure.Options{

	STSSeconds:            31536000,
	STSIncludeSubdomains:  true,
	STSPreload:            true,
	FrameDeny:             true,
	ContentTypeNosniff:    true,
	BrowserXssFilter:      true,
	ContentSecurityPolicy: "script-src $NONCE",
}

func NewSecurityMiddleware(options secure.Options) *secure.Secure {
	return secure.New(options)
}

func NewDefaultSecurityMiddleware() *secure.Secure {
	return secure.New(DefaultSecureOption)
}
