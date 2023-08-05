package menuitem

import (
	"context"
	"my-app/backend/pkg/tray"
)

var _root tray.IMenuItemBase

func init() {
	_root = newRoot(context.Background())
}

func Root() tray.IMenuItemBase {
	return _root
}
