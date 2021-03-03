# Secuware Common Security
## Description
 Common helper module for http security standard ,such us :
  * Security Header 
  * Cryptography [In Progress]
  * etc 

## 3rd Party Mod
 - [Secure v1.0.8](https://github.com/unrolled/secure)


## Package

| Package                                 | Description                                                                                                                                                  |
| --------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| secuware/http/middlewares/headers | Middleware for add security standard header, this is useful for prevent security attack such us `XXS`, `CFS` also for enable `Strict-Transport-Security` etc |


## Usage

### Header Security


* Default Option
  
```
  headerMiddleware := header.NewDefaultSecurityMiddleware()

  //chi router
  r.Use(headerMiddleware.Handle)
```


* Default options with few customize option


```
  secOptions := headers.NewSecureOptionsBuilder().
		WithDefaultOpts().
		AddForceSTSHeader(false).
		AddIsDevelopment(true).
		AddSSLRedirect(false).
    // etc
		Build()
  headerMiddleware := header.NewSecurityMiddleware(secOptions)

  //chi router
  r.Use(headerMiddleware.Handle)
```

* Customize option

```
  secOptions := headers.NewSecureOptionsBuilder().
		AddForceSTSHeader(false).
		AddIsDevelopment(true).
		AddSSLRedirect(false).
    // etc
		Build()
  headerMiddleware := header.NewSecurityMiddleware(secOptions)

  //chi router
  r.Use(headerMiddleware.Handle)
```

### Jwt validation with Jwks

* Add `JwksValidationMiddleware` middleware to your chi router config
  
```
  r.Use(secuware.JwksValidationMiddleware(commonsecurity.
		JwksValidationMiddlewareOpts().
		WithDefaultOpts().
		SetJwksUrl("<ExternalIP>>/auth/realms/<NameRealms>/protocol/openid-connect/certs").
		AddWhitelist("/api/*").
		Build()))

```

#### Filter access by user session grant access

* Filter by client roles
```
import (
    secuwareroutes "github.com/Alter17Ego/secuware/http/routers"
)

// then on your router add this HOF(Higher Order Function)
router.Get("/ping/role-based",
        secuwareroutes.RouteWithClientRoles([]string{"ping-client"}, delivery.Ping))
```


* Filter by realm roles
```
import (
    secuwareroutes "github.com/Alter17Ego/secuware/http/routers"
)

// then on your router add this HOF(Higher Order Function)
router.Get("/ping/role-based",
        secuwareroutes.RouteWithRealmRoles([]string{"health-check"}, delivery.Ping))
```

* Filter by resource names
```
import (
    secuwareroutes "github.com/Alter17Ego/secuware/http/routers"
    "github.com/Alter17Ego/secuware/http/routers/resourcebased"
)

// then on your router add this HOF(Higher Order Function)
router.Get("/ping/resource-based", secuwareroutes.RouteWithResources([]resourcebased.Resource{
        resourcebased.Resource{
            Name:  "user-management",
        },
    }, delivery.Ping))
```

* Filter by scoped resource
```
import (
    secuwareroutes "github.com/Alter17Ego/secuware/http/routers"
    "github.com/Alter17Ego/secuware/http/routers/resourcebased"
)

// then on your router add this HOF(Higher Order Function)
router.Get("/ping/resource-based", secuwareroutes.RouteWithScopedResources([]resourcebased.Resource{
        resourcebased.Resource{
            Name:  "user-management",
            Scope: resourcebased.Scopes{
                "ping",
            },
        },
    }, delivery.Ping))
```

* Advance filtering by composing policies
```
import (
    secuwareroutes "github.com/Alter17Ego/secuware/http/routers"
    "github.com/Alter17Ego/secuware/http/routers/resourcebased"
    "github.com/Alter17Ego/secuware/http/routers/rolebased"
)

router.Get("/ping/resource-based", secuwareroutes.RouteWithPolicy(core.MatchAll,core.Policies{
        rolebased.HasClientRoles("healthcheck","monitoring"),
        rolebased.HasClientRoles("ping-my-client"),
        resourcebased.HasScopedResources(resourcebased.Resource{
            Name:  "user-management",
            Scope: resourcebased.Scopes{"ping"},
        }),
    },delivery.Ping))
```