package menuitems

import (
	"context"
	"my-app/backend/pkg/tray"
)

type separator struct {
	*tray.MenuItem
}

// Separator implements tray.IMenuItem.
func (*separator) Separator() bool {
	return true
}

func newSeparator(ctx context.Context) tray.IMenuItem {
	return &separator{
		MenuItem: tray.NewMenuItem(ctx),
	}
}
