package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"

	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type colorThemeDark struct {
	*tray.MenuItem
}

// CanCheck implements tray.IMenuItem.
func (*colorThemeDark) CanCheck() bool {
	return true
}

// Checked implements tray.IMenuItem.
func (*colorThemeDark) Checked() bool {
	return false
}

// Key implements tray.IMenuItem.
func (*colorThemeDark) Key() string {
	return "color.theme.dark"
}

// CanClick implements tray.IMenuItem.
func (*colorThemeDark) CanClick() bool {
	return true
}

// OnClick implements tray.IMenuItem.
func (t *colorThemeDark) OnClick() (quit bool) {
	app.THEME(windows.Dark)
	runtime.WindowSetDarkTheme(t.Ctx)
	return false
}

// Title implements tray.IMenuItem.
func (*colorThemeDark) Title() string {
	return app.T().ColorTheme.Dark
}

// Tooltip implements tray.IMenuItem.
func (*colorThemeDark) Tooltip() string {
	return app.T().ColorTheme.Dark
}

func newColorThemeDark(ctx context.Context) tray.IMenuItem {
	return &colorThemeDark{
		MenuItem: tray.NewMenuItem(ctx),
	}
}
