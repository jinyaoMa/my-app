package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"

	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type colorThemeSystem struct {
	*tray.MenuItem
}

// CanCheck implements tray.IMenuItem.
func (*colorThemeSystem) CanCheck() bool {
	return true
}

// Checked implements tray.IMenuItem.
func (*colorThemeSystem) Checked() bool {
	return false
}

// Key implements tray.IMenuItem.
func (*colorThemeSystem) Key() string {
	return "color.theme.system"
}

// CanClick implements tray.IMenuItem.
func (*colorThemeSystem) CanClick() bool {
	return true
}

// OnClick implements tray.IMenuItem.
func (t *colorThemeSystem) OnClick() (quit bool) {
	app.THEME(windows.SystemDefault)
	runtime.WindowSetSystemDefaultTheme(t.Ctx)
	return false
}

// Title implements tray.IMenuItem.
func (*colorThemeSystem) Title() string {
	return app.T().ColorTheme.System
}

// Tooltip implements tray.IMenuItem.
func (*colorThemeSystem) Tooltip() string {
	return app.T().ColorTheme.System
}

func newColorThemeSystem(ctx context.Context) tray.IMenuItem {
	return &colorThemeSystem{
		MenuItem: tray.NewMenuItem(ctx),
	}
}
