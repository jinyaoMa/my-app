package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"

	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	return app.THEME() == windows.Light
}

// Key implements tray.IMenuItem.
func (*colorThemeLight) Key() string {
	return "color.theme.light"
}

// CanClick implements tray.IMenuItem.
func (*colorThemeLight) CanClick() bool {
	return true
}

// OnClick implements tray.IMenuItem.
func (t *colorThemeLight) OnClick() (quit bool) {
	if app.THEME(windows.Light) == windows.Light {
		runtime.WindowSetLightTheme(t.Ctx)
		tray.Update(_root)
	}
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
		MenuItem: tray.NewMenuItem(ctx),
	}
}
