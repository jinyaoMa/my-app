package menuitem

import (
	"context"
	"my-app/backend/pkg/tray"
)

var _root tray.Interface

func init() {
	_root = newRoot(context.Background())
}

func Root() tray.Interface {
	return _root
}
