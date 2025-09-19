package authfwt

import (
	"net/http"
	"strings"

	"github.com/apache/fory/go/fory"
	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/fwt"
	"majinyao.cn/my-app/backend/pkg/memcache"
)

func MustNew[T fwt.IdentityGetter](
	api huma.API,
	scheme string,
	options Options,
	scopesValidator func(ctx huma.Context, scopes []string) error,
	registers ...func(f *fory.Fory) error,
) func(ctx huma.Context, next func(huma.Context)) {

	handler, err := New[T](api, scheme, options, scopesValidator, registers...)
	if err != nil {
		panic(err)
	}
	return handler
}

// should not use multiple authfwt middlewares in the same request handler chain
// because the first middleware will set the user data in the context, and the second middleware will set the user data to the same context,
// which will override the user data set by the first middleware.
func New[T fwt.IdentityGetter](
	api huma.API,
	scheme string,
	options Options,
	scopesValidator func(ctx huma.Context, scopes []string) error,
	registers ...func(f *fory.Fory) error,
) (func(ctx huma.Context, next func(huma.Context)), error) {

	f, err := fwt.New[T](options.Fwt, registers...)
	if err != nil {
		return nil, err
	}

	m := memcache.New(options.CacheLimit)

	if api.OpenAPI().Components.SecuritySchemes == nil {
		api.OpenAPI().Components.SecuritySchemes = make(map[string]*huma.SecurityScheme)
	}
	api.OpenAPI().Components.SecuritySchemes[scheme] = &huma.SecurityScheme{
		Type:         "http",
		Description:  "Fory Web Token",
		Name:         "Authorization",
		In:           "header",
		Scheme:       "bearer",
		BearerFormat: "Fwt",
	}

	return func(ctx huma.Context, next func(huma.Context)) {
		var scopes []string
		isAuthorizationRequired := false
		for _, s := range ctx.Operation().Security {
			if scopes, isAuthorizationRequired = s[scheme]; isAuthorizationRequired {
				break
			}
		}

		ctx = AttachFwtToHumaContext(ctx, f)
		ctx = AttachMemcacheToHumaContext(ctx, m)
		if !isAuthorizationRequired {
			next(ctx)
			return
		}

		token := strings.TrimPrefix(ctx.Header("Authorization"), "Bearer ")
		if len(token) == 0 {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "token is empty")
			return
		}

		var claims fwt.Claims[T]
		if cacheItem, err := m.Get(token); err == nil {
			claims = cacheItem.(fwt.Claims[T])
		} else if claims, err = f.ParseAccessToken(token); err == nil {
			if err := f.ValidateClaims(claims); err == nil {
				m.Set(token, claims, claims.ExpiredAt)
			} else if scopes != nil { // if scopes is not nil, claims must be valid
				huma.WriteErr(api, ctx, http.StatusForbidden, err.Error())
				return
			}
		} else {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "token parse failed")
			return
		}

		if scopes != nil {
			if err := scopesValidator(ctx, scopes); err != nil {
				huma.WriteErr(api, ctx, http.StatusForbidden, err.Error())
				return
			}
		}

		ctx = AttachClaimsToHumaContext(ctx, claims)
		next(ctx)
	}, nil
}
