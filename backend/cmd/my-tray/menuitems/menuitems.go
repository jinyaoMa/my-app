package menuitems

import (
	"context"
	"my-app/backend/pkg/tray"
)

var _root tray.IMenuItemBase

func init() {
	_root = newRoot(context.TODO())
}

func Root() tray.IMenuItemBase {
	return _root
}
