package headers_test

import (
	"reflect"
	"testing"

	"github.com/unrolled/secure"

	"github.com/Alter17Ego/secuware/http/middlewares/headers"
)

func TestDefaultOption(t *testing.T) {
	expectedOptions := headers.DefaultSecureOption
	secOptions := headers.NewSecureOptionsBuilder().WithDefaultOpts().Build()
	if !reflect.DeepEqual(secOptions, expectedOptions) {
		t.Errorf("%v is not equals to %v", secOptions, expectedOptions)
	}
}

func TestSetOptionsWithDefaultOptions(t *testing.T) {
	expectedOptions := secure.Options{
		BrowserXssFilter:                headers.DefaultSecureOption.BrowserXssFilter,
		ContentTypeNosniff:              headers.DefaultSecureOption.ContentTypeNosniff,
		ForceSTSHeader:                  false,
		FrameDeny:                       headers.DefaultSecureOption.FrameDeny,
		IsDevelopment:                   true,
		SSLRedirect:                     false,
		SSLForceHost:                    false,
		SSLTemporaryRedirect:            true,
		STSIncludeSubdomains:            headers.DefaultSecureOption.STSIncludeSubdomains,
		STSPreload:                      headers.DefaultSecureOption.STSPreload,
		ContentSecurityPolicy:           headers.DefaultSecureOption.ContentSecurityPolicy,
		ContentSecurityPolicyReportOnly: "content-b",
		CustomBrowserXssValue:           "content-c",
		CustomFrameOptionsValue:         "content-d",
		PublicKey:                       "lorem",
		ReferrerPolicy:                  "lorem-2",
		FeaturePolicy:                   "feature-2",
		SSLHost:                         "http://test.go",
		AllowedHosts:                    []string{"http://test.go"},
		AllowedHostsAreRegex:            true,
		HostsProxyHeaders:               []string{"Method"},
		SSLHostFunc:                     nil,
		SSLProxyHeaders:                 map[string]string{"hello": "hello"},
		STSSeconds:                      headers.DefaultSecureOption.STSSeconds,
		ExpectCTHeader:                  "ctheader",
	}
	secOptions := headers.NewSecureOptionsBuilder().
		WithDefaultOpts().
		AddForceSTSHeader(false).
		AddIsDevelopment(true).
		AddSSLRedirect(false).
		AddSSLForceHost(false).
		AddSSLTemporaryRedirect(true).
		AddContentSecurityPolicyReportOnly("content-b").
		AddCustomBrowserXssValue("content-c").
		AddCustomFrameOptionsValue("content-d").
		AddPublicKey("lorem").
		AddReferrerPolicy("lorem-2").
		AddFeaturePolicy("feature-2").
		AddAllowedHosts([]string{"http://test.go"}).
		AddAllowedHostsAreRegex(true).
		AddHostsProxyHeaders([]string{"Method"}).
		AddSSLHostFunc(nil).
		AddSSLProxyHeaders(map[string]string{"hello": "hello"}).
		AddExpectCTHeader("ctheader").
		AddSSLHost("http://test.go").
		Build()
	if !reflect.DeepEqual(secOptions, expectedOptions) {
		t.Errorf("%v is not equals to %v", secOptions, expectedOptions)
	}
}

func TestSetOptions(t *testing.T) {
	expectedOptions := secure.Options{
		BrowserXssFilter:                false,
		ContentTypeNosniff:              false,
		ForceSTSHeader:                  false,
		FrameDeny:                       true,
		IsDevelopment:                   true,
		SSLRedirect:                     false,
		SSLForceHost:                    false,
		SSLTemporaryRedirect:            true,
		STSIncludeSubdomains:            true,
		STSPreload:                      true,
		ContentSecurityPolicy:           "content-a",
		ContentSecurityPolicyReportOnly: "content-b",
		CustomBrowserXssValue:           "content-c",
		CustomFrameOptionsValue:         "content-d",
		PublicKey:                       "lorem",
		ReferrerPolicy:                  "lorem-2",
		FeaturePolicy:                   "feature-2",
		SSLHost:                         "http://test.go",
		AllowedHosts:                    []string{"http://test.go"},
		AllowedHostsAreRegex:            true,
		HostsProxyHeaders:               []string{"Method"},
		SSLHostFunc:                     nil,
		SSLProxyHeaders:                 map[string]string{"hello": "hello"},
		STSSeconds:                      3000,
		ExpectCTHeader:                  "ctheader",
	}
	secOptions := headers.NewSecureOptionsBuilder().
		AddBrowserXssFilter(false).
		AddContentTypeNosniff(false).
		AddForceSTSHeader(false).
		AddFrameDeny(true).
		AddIsDevelopment(true).
		AddSSLRedirect(false).
		AddSSLForceHost(false).
		AddSSLTemporaryRedirect(true).
		AddSTSIncludeSubdomains(true).
		AddSTSPreload(true).
		AddContentSecurityPolicy("content-a").
		AddContentSecurityPolicyReportOnly("content-b").
		AddCustomBrowserXssValue("content-c").
		AddCustomFrameOptionsValue("content-d").
		AddPublicKey("lorem").
		AddReferrerPolicy("lorem-2").
		AddFeaturePolicy("feature-2").
		AddAllowedHosts([]string{"http://test.go"}).
		AddAllowedHostsAreRegex(true).
		AddHostsProxyHeaders([]string{"Method"}).
		AddSSLHostFunc(nil).
		AddSSLProxyHeaders(map[string]string{"hello": "hello"}).
		AddSTSSeconds(3000).
		AddExpectCTHeader("ctheader").
		AddSSLHost("http://test.go").
		Build()
	if !reflect.DeepEqual(secOptions, expectedOptions) {
		t.Errorf("%v is not equals to %v", secOptions, expectedOptions)
	}
}
