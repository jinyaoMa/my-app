package menuitem

import (
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type openWindow struct{}

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
	return app.Assets().GetBytes("tray.ico")
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
func (*openWindow) OnClick() (quit bool) {
	return false
}

// Separator implements tray.IMenuItem.
func (*openWindow) Separator() bool {
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

// Visible implements tray.IMenuItem.
func (*openWindow) Visible() bool {
	return true
}

func newOpenWindow() tray.IMenuItem {
	return &openWindow{}
}
