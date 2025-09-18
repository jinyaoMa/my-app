package memcache

import (
	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/memcache"
)

func New(limit int) func(ctx huma.Context, next func(huma.Context)) {
	m := memcache.New(limit)
	return func(ctx huma.Context, next func(huma.Context)) {
		ctx = Attach(ctx, m)
		next(ctx)
	}
}
