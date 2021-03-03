package headers

import "github.com/unrolled/secure"

type SecureOptionsBuilder struct {
	BrowserXssFilter *bool

	ContentTypeNosniff *bool

	ForceSTSHeader *bool

	FrameDeny *bool

	IsDevelopment *bool

	nonceEnabled *bool

	SSLRedirect *bool

	SSLForceHost *bool

	SSLTemporaryRedirect *bool

	STSIncludeSubdomains *bool

	STSPreload *bool

	ContentSecurityPolicy *string

	ContentSecurityPolicyReportOnly *string

	CustomBrowserXssValue *string

	CustomFrameOptionsValue *string

	PublicKey *string

	ReferrerPolicy *string

	FeaturePolicy *string

	SSLHost *string

	AllowedHosts *[]string

	AllowedHostsAreRegex *bool

	HostsProxyHeaders *[]string

	SSLHostFunc *secure.SSLHostFunc

	SSLProxyHeaders *map[string]string

	STSSeconds *int64

	ExpectCTHeader    *string
	withDefaultOption bool
}

func (b *SecureOptionsBuilder) AddSSLHostFunc(value *secure.SSLHostFunc) *SecureOptionsBuilder {
	b.SSLHostFunc = value
	return b
}

func (b *SecureOptionsBuilder) AddSSLRedirect(value bool) *SecureOptionsBuilder {
	b.SSLRedirect = &value
	return b
}
func (b *SecureOptionsBuilder) AddBrowserXssFilter(value bool) *SecureOptionsBuilder {
	b.BrowserXssFilter = &value
	return b
}
func (b *SecureOptionsBuilder) AddContentTypeNosniff(value bool) *SecureOptionsBuilder {
	b.ContentTypeNosniff = &value
	return b
}
func (b *SecureOptionsBuilder) AddForceSTSHeader(value bool) *SecureOptionsBuilder {
	b.SSLRedirect = &value
	return b
}
func (b *SecureOptionsBuilder) AddFrameDeny(value bool) *SecureOptionsBuilder {
	b.FrameDeny = &value
	return b
}
func (b *SecureOptionsBuilder) AddIsDevelopment(value bool) *SecureOptionsBuilder {
	b.IsDevelopment = &value
	return b
}
func (b *SecureOptionsBuilder) AddSSLForceHost(value bool) *SecureOptionsBuilder {
	b.SSLForceHost = &value
	return b
}
func (b *SecureOptionsBuilder) AddSSLTemporaryRedirect(value bool) *SecureOptionsBuilder {
	b.SSLTemporaryRedirect = &value
	return b
}
func (b *SecureOptionsBuilder) AddSTSIncludeSubdomains(value bool) *SecureOptionsBuilder {
	b.STSIncludeSubdomains = &value
	return b
}
func (b *SecureOptionsBuilder) AddSTSPreload(value bool) *SecureOptionsBuilder {
	b.STSPreload = &value
	return b
}
func (b *SecureOptionsBuilder) AddContentSecurityPolicy(value string) *SecureOptionsBuilder {
	b.ContentSecurityPolicy = &value
	return b
}
func (b *SecureOptionsBuilder) AddContentSecurityPolicyReportOnly(value string) *SecureOptionsBuilder {
	b.ContentSecurityPolicyReportOnly = &value
	return b
}
func (b *SecureOptionsBuilder) AddCustomBrowserXssValue(value string) *SecureOptionsBuilder {
	b.CustomBrowserXssValue = &value
	return b
}
func (b *SecureOptionsBuilder) AddCustomFrameOptionsValue(value string) *SecureOptionsBuilder {
	b.CustomFrameOptionsValue = &value
	return b
}
func (b *SecureOptionsBuilder) AddPublicKey(value string) *SecureOptionsBuilder {
	b.PublicKey = &value
	return b
}
func (b *SecureOptionsBuilder) AddReferrerPolicy(value string) *SecureOptionsBuilder {
	b.ReferrerPolicy = &value
	return b
}
func (b *SecureOptionsBuilder) AddFeaturePolicy(value string) *SecureOptionsBuilder {
	b.FeaturePolicy = &value
	return b
}
func (b *SecureOptionsBuilder) AddSSLHost(value string) *SecureOptionsBuilder {
	b.SSLHost = &value
	return b
}
func (b *SecureOptionsBuilder) AddAllowedHosts(value []string) *SecureOptionsBuilder {
	b.AllowedHosts = &value
	return b
}
func (b *SecureOptionsBuilder) AddAllowedHostsAreRegex(value bool) *SecureOptionsBuilder {
	b.AllowedHostsAreRegex = &value
	return b
}
func (b *SecureOptionsBuilder) AddHostsProxyHeaders(value []string) *SecureOptionsBuilder {
	b.HostsProxyHeaders = &value
	return b
}
func (b *SecureOptionsBuilder) AddSSLProxyHeaders(value map[string]string) *SecureOptionsBuilder {
	b.SSLProxyHeaders = &value
	return b
}
func (b *SecureOptionsBuilder) AddSTSSeconds(value int64) *SecureOptionsBuilder {
	b.STSSeconds = &value
	return b
}
func (b *SecureOptionsBuilder) AddExpectCTHeader(value string) *SecureOptionsBuilder {
	b.ExpectCTHeader = &value
	return b
}

func (b *SecureOptionsBuilder) WithDefaultOpts() *SecureOptionsBuilder {
	b.withDefaultOption = true
	return b
}

func (b *SecureOptionsBuilder) Build() secure.Options {
	// TODO: merge object automatically
	var options secure.Options
	if b.withDefaultOption {
		options = DefaultSecureOption
	} else {
		options = secure.Options{}
	}

	if b.SSLRedirect != nil {
		options.SSLRedirect = *b.SSLRedirect
	}
	if b.BrowserXssFilter != nil {
		options.BrowserXssFilter = *b.BrowserXssFilter
	}
	if b.ContentTypeNosniff != nil {
		options.ContentTypeNosniff = *b.ContentTypeNosniff
	}
	if b.ForceSTSHeader != nil {
		options.ForceSTSHeader = *b.SSLRedirect
	}
	if b.FrameDeny != nil {
		options.FrameDeny = *b.FrameDeny
	}
	if b.IsDevelopment != nil {
		options.IsDevelopment = *b.IsDevelopment
	}
	if b.SSLForceHost != nil {
		options.SSLForceHost = *b.SSLForceHost
	}
	if b.SSLTemporaryRedirect != nil {
		options.SSLTemporaryRedirect = *b.SSLTemporaryRedirect
	}
	if b.STSIncludeSubdomains != nil {
		options.STSIncludeSubdomains = *b.STSIncludeSubdomains
	}
	if b.STSPreload != nil {
		options.STSPreload = *b.STSPreload
	}
	if b.ContentSecurityPolicy != nil {
		options.ContentSecurityPolicy = *b.ContentSecurityPolicy
	}
	if b.ContentSecurityPolicyReportOnly != nil {
		options.ContentSecurityPolicyReportOnly = *b.ContentSecurityPolicyReportOnly
	}
	if b.CustomBrowserXssValue != nil {
		options.CustomBrowserXssValue = *b.CustomBrowserXssValue
	}
	if b.CustomFrameOptionsValue != nil {
		options.CustomFrameOptionsValue = *b.CustomFrameOptionsValue
	}
	if b.PublicKey != nil {
		options.PublicKey = *b.PublicKey
	}
	if b.ReferrerPolicy != nil {
		options.ReferrerPolicy = *b.ReferrerPolicy
	}
	if b.FeaturePolicy != nil {
		options.FeaturePolicy = *b.FeaturePolicy
	}
	if b.SSLHost != nil {
		options.SSLHost = *b.SSLHost
	}
	if b.AllowedHosts != nil {
		options.AllowedHosts = *b.AllowedHosts
	}
	if b.AllowedHostsAreRegex != nil {
		options.AllowedHostsAreRegex = *b.AllowedHostsAreRegex
	}
	if b.HostsProxyHeaders != nil {
		options.HostsProxyHeaders = *b.HostsProxyHeaders
	}
	if b.SSLProxyHeaders != nil {
		options.SSLProxyHeaders = *b.SSLProxyHeaders
	}
	if b.STSSeconds != nil {
		options.STSSeconds = *b.STSSeconds
	}
	if b.ExpectCTHeader != nil {
		options.ExpectCTHeader = *b.ExpectCTHeader
	}
	if b.SSLHostFunc != nil {
		options.SSLHostFunc = b.SSLHostFunc
	}

	return options
}

func NewSecureOptionsBuilder() *SecureOptionsBuilder {
	return &SecureOptionsBuilder{}
}
