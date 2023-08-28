package menuitem

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type openWindow struct {
	*tray.MenuItem
}

// Icon implements tray.IMenuItem.
func (*openWindow) Icon() []byte {
	return app.ASSETS().GetBytes("tray.ico")
}

// Key implements tray.IMenuItem.
func (*openWindow) Key() string {
	return "open.window"
}

// OnClick implements tray.IMenuItem.
func (w *openWindow) OnClick() (quit bool) {
	runtime.WindowShow(w.Ctx)
	return false
}

// Title implements tray.IMenuItem.
func (*openWindow) Title() string {
	return app.T().OpenWindow
}

// Tooltip implements tray.IMenuItem.
func (*openWindow) Tooltip() string {
	return app.T().OpenWindow
}

func newOpenWindow(ctx context.Context) tray.IMenuItem {
	return &openWindow{
		MenuItem: &tray.MenuItem{
			MenuItemBase: &tray.MenuItemBase{
				Ctx: ctx,
			},
		},
	}
}
