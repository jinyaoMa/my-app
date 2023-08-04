package menuitem

import (
	"context"
	"my-app/backend/pkg/tray"
	"sync"
)

var (
	_root tray.Interface
	once  sync.Once
)

func Root(ctx ...context.Context) tray.Interface {
	once.Do(func() {
		if len(ctx) > 0 {
			_root = newRoot(ctx[0])
		} else {
			_root = newRoot(context.Background())
		}
	})
	return _root
}

func BindContext(ctx context.Context) {
	_root.SetContext(ctx)
}
