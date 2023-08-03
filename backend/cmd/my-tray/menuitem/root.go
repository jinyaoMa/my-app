package menuitem

import (
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type root struct{}

// Icon implements tray.Interface.
func (*root) Icon() []byte {
	if app.App().Web().IsStopping() {
		return app.App().Assets().GetBytes("tray.orange.ico")
	}
	if app.App().Web().IsRunning() {
		return app.App().Assets().GetBytes("tray.green.ico")
	}
	return app.App().Assets().GetBytes("tray.blue.ico")
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
	return app.App().T().AppName
}

// Tooltip implements tray.Interface.
func (*root) Tooltip() string {
	return app.App().T().AppName
}

func newRoot() tray.Interface {
	return &root{}
}
