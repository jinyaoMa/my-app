package menuitem

import "my-app/backend/pkg/tray"

var _root tray.Interface

func init() {
	_root = newRoot()
}

func Root() tray.Interface {
	return _root
}
