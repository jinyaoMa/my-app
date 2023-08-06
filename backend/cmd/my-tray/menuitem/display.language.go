package menuitem

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type displayLanguage struct {
	*tray.MenuItem
}

// Title implements tray.IMenuItem.
func (*displayLanguage) Title() string {
	return app.T().DisplayLanguage.Title
}

// Tooltip implements tray.IMenuItem.
func (*displayLanguage) Tooltip() string {
	return app.T().DisplayLanguage.Title
}

func newdisplayLanguage(ctx context.Context) tray.IMenuItem {
	menuitems := make([]tray.IMenuItem, 0)

	return &displayLanguage{
		MenuItem: &tray.MenuItem{
			MenuItemBase: &tray.MenuItemBase{
				Ctx:       ctx,
				MenuItems: []tray.IMenuItem{},
			},
		},
	}
}
