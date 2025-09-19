package memcache

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/memcache"
)

type key int

const (
	key_ key = iota + 1
)

func AttachToHumaContext(ctx huma.Context, m memcache.IMemcache) huma.Context {
	if ctx.Context().Value(key_) != nil {
		panic("memcache already attached, please check your middlewares")
	}
	return huma.WithValue(ctx, key_, m)
}

func GetFromHumaContext(ctx huma.Context) memcache.IMemcache {
	m, ok := ctx.Context().Value(key_).(memcache.IMemcache)
	if !ok {
		panic("memcache not found, please check your middlewares")
	}
	return m
}

func GetFromContext(ctx context.Context) memcache.IMemcache {
	m, ok := ctx.Value(key_).(memcache.IMemcache)
	if !ok {
		panic("memcache not found, please check your middlewares")
	}
	return m
}
