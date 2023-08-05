package menuitem

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type colorThemeLight struct {
	*tray.MenuItem
}

// CanCheck implements tray.IMenuItem.
func (*colorThemeLight) CanCheck() bool {
	return true
}

// Checked implements tray.IMenuItem.
func (*colorThemeLight) Checked() bool {
	return false
}

// Key implements tray.IMenuItem.
func (*colorThemeLight) Key() string {
	return ""
}

// OnClick implements tray.IMenuItem.
func (t *colorThemeLight) OnClick() (quit bool) {
	return false
}

// Title implements tray.IMenuItem.
func (*colorThemeLight) Title() string {
	return app.T().ColorTheme.Light
}

// Tooltip implements tray.IMenuItem.
func (*colorThemeLight) Tooltip() string {
	return app.T().ColorTheme.Light
}

func newColorThemeLight(ctx context.Context) tray.IMenuItem {
	return &colorThemeLight{
		MenuItem: &tray.MenuItem{
			MenuItemBase: &tray.MenuItemBase{
				Ctx: ctx,
			},
		},
	}
}
