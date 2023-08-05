package menuitem

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type colorTheme struct {
	*tray.MenuItem
}

// Title implements tray.IMenuItem.
func (*colorTheme) Title() string {
	return app.T().ColorTheme.Title
}

// Tooltip implements tray.IMenuItem.
func (*colorTheme) Tooltip() string {
	return app.T().ColorTheme.Title
}

func newColorTheme(ctx context.Context) tray.IMenuItem {
	return &colorTheme{
		MenuItem: &tray.MenuItem{
			MenuItemBase: &tray.MenuItemBase{
				Ctx: ctx,
				MenuItems: []tray.IMenuItem{
					newColorThemeSystem(ctx),
					newColorThemeLight(ctx),
					newColorThemeDark(ctx),
				},
			},
		},
	}
}
