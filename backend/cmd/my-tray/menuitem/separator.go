package menuitem

import "my-app/backend/pkg/tray"

type separator struct {
	tray.IMenuItem
}

// Separator implements tray.IMenuItem.
func (*separator) Separator() bool {
	return true
}

func newSeparator() tray.IMenuItem {
	return &separator{}
}
