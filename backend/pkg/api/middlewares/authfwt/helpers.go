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
	keyUserData
)

func AttachFwt[T fwt.IdentityGetter](ctx huma.Context, f fwt.IFwt[T]) huma.Context {
	if ctx.Context().Value(keyFwt) != nil {
		panic("authfwt fwt already attached, please check your middlewares")
	}
	return huma.WithValue(ctx, keyFwt, f)
}

func GetFwt[T fwt.IdentityGetter](ctx huma.Context) fwt.IFwt[T] {
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

func AttachMemcache(ctx huma.Context, m memcache.IMemcache) huma.Context {
	if ctx.Context().Value(keyMemcache) != nil {
		panic("authfwt memcache already attached, please check your middlewares")
	}
	return huma.WithValue(ctx, keyMemcache, m)
}

func GetMemcache(ctx huma.Context) memcache.IMemcache {
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

func AttachUserData[T fwt.IdentityGetter](ctx huma.Context, userdata T) huma.Context {
	if ctx.Context().Value(keyUserData) != nil {
		panic("authfwt userdata already attached, please check your middlewares")
	}
	return huma.WithValue(ctx, keyUserData, userdata)
}

func GetUserData[T fwt.IdentityGetter](ctx huma.Context) T {
	userdata, _ := ctx.Context().Value(keyUserData).(T)
	return userdata
}

func GetUserDataFromContext[T fwt.IdentityGetter](ctx context.Context) T {
	userdata, _ := ctx.Value(keyUserData).(T)
	return userdata
}
