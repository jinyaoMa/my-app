package menuitem

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type openWindow struct {
	ctx context.Context
}

// SetContext implements tray.IMenuItem.
func (w *openWindow) SetContext(ctx context.Context) {
	w.ctx = ctx
}

// CanCheck implements tray.IMenuItem.
func (*openWindow) CanCheck() bool {
	return false
}

// Checked implements tray.IMenuItem.
func (*openWindow) Checked() bool {
	return false
}

// Enabled implements tray.IMenuItem.
func (*openWindow) Enabled() bool {
	return true
}

// Icon implements tray.IMenuItem.
func (*openWindow) Icon() []byte {
	return app.App().Assets().GetBytes("tray.ico")
}

// Items implements tray.IMenuItem.
func (*openWindow) Items() []tray.IMenuItem {
	return nil
}

// Key implements tray.IMenuItem.
func (*openWindow) Key() string {
	return "open.window"
}

// OnClick implements tray.IMenuItem.
func (w *openWindow) OnClick() (quit bool) {
	runtime.WindowShow(w.ctx)
	return false
}

// Separator implements tray.IMenuItem.
func (*openWindow) Separator() bool {
	return false
}

// Title implements tray.IMenuItem.
func (*openWindow) Title() string {
	return app.App().T().OpenWindow
}

// Tooltip implements tray.IMenuItem.
func (*openWindow) Tooltip() string {
	return app.App().T().OpenWindow
}

// Visible implements tray.IMenuItem.
func (*openWindow) Visible() bool {
	return true
}

func newOpenWindow(ctx context.Context) tray.IMenuItem {
	return &openWindow{
		ctx: ctx,
	}
}
