package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type colorTheme struct {
	*tray.MenuItem
}

// Key implements tray.IMenuItem.
func (*colorTheme) Key() string {
	return "color.theme"
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
		MenuItem: tray.NewMenuItem(ctx,
			newColorThemeSystem(ctx),
			newColorThemeLight(ctx),
			newColorThemeDark(ctx)),
	}
}
