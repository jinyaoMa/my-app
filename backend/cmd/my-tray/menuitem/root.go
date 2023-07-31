package menuitem

import (
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type root struct{}

// Icon implements tray.Interface.
func (*root) Icon() []byte {
	if app.Web().IsStopping() {
		return app.Assets().GetBytes("tray.orange.ico")
	}
	if app.Web().IsRunning() {
		return app.Assets().GetBytes("tray.green.ico")
	}
	return app.Assets().GetBytes("tray.blue.ico")
}

// Items implements tray.Interface.
func (*root) Items() []tray.IMenuItem {
	return []tray.IMenuItem{
		newOpenWindow(),
		newSeparator(),
		newQuit(),
	}
}

// Title implements tray.Interface.
func (*root) Title() string {
	return app.T().AppName
}

// Tooltip implements tray.Interface.
func (*root) Tooltip() string {
	return app.T().AppName
}

func newRoot() tray.Interface {
	return &root{}
}
