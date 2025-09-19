package authfwt

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/fwt"
	"majinyao.cn/my-app/backend/pkg/memcache"
)

type key int

const (
	keyFwt key = iota + 1
	keyMemcache
	keyClaims
)

func AttachFwtToHumaContext[T fwt.IdentityGetter](ctx huma.Context, f fwt.IFwt[T]) huma.Context {
	if ctx.Context().Value(keyFwt) != nil {
		panic("authfwt fwt already attached, please check your middlewares")
	}
	return huma.WithValue(ctx, keyFwt, f)
}

func GetFwtFromHumaContext[T fwt.IdentityGetter](ctx huma.Context) fwt.IFwt[T] {
	f, ok := ctx.Context().Value(keyFwt).(fwt.IFwt[T])
	if !ok {
		panic("authfwt fwt not found, please check your middlewares")
	}
	return f
}

func GetFwtFromContext[T fwt.IdentityGetter](ctx context.Context) fwt.IFwt[T] {
	f, ok := ctx.Value(keyFwt).(fwt.IFwt[T])
	if !ok {
		panic("authfwt fwt not found, please check your middlewares")
	}
	return f
}

func AttachMemcacheToHumaContext(ctx huma.Context, m memcache.IMemcache) huma.Context {
	if ctx.Context().Value(keyMemcache) != nil {
		panic("authfwt memcache already attached, please check your middlewares")
	}
	return huma.WithValue(ctx, keyMemcache, m)
}

func GetMemcacheFromHumaContext(ctx huma.Context) memcache.IMemcache {
	m, ok := ctx.Context().Value(keyMemcache).(memcache.IMemcache)
	if !ok {
		panic("authfwt memcache not found, please check your middlewares")
	}
	return m
}

func GetMemcacheFromContext(ctx context.Context) memcache.IMemcache {
	m, ok := ctx.Value(keyMemcache).(memcache.IMemcache)
	if !ok {
		panic("authfwt memcache not found, please check your middlewares")
	}
	return m
}

func AttachClaimsToHumaContext[T fwt.IdentityGetter](ctx huma.Context, claims fwt.Claims[T]) huma.Context {
	if ctx.Context().Value(keyClaims) != nil {
		panic("authfwt fwt claims already attached, please check your middlewares")
	}
	return huma.WithValue(ctx, keyClaims, claims)
}

func GetClaimsFromHumaContext[T fwt.IdentityGetter](ctx huma.Context) fwt.Claims[T] {
	claims, _ := ctx.Context().Value(keyClaims).(fwt.Claims[T])
	return claims
}

func GetClaimsFromContext[T fwt.IdentityGetter](ctx context.Context) fwt.Claims[T] {
	claims, _ := ctx.Value(keyClaims).(fwt.Claims[T])
	return claims
}
