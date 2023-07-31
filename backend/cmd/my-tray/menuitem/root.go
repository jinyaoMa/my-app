package menuitem

import (
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type Root struct {
	tray.Interface
}

// Icon implements tray.Interface.
func (*Root) Icon() []byte {
	if app.Web().IsStopping() {
		return app.Assets().GetBytes("tray.orange.ico")
	}
	if app.Web().IsRunning() {
		return app.Assets().GetBytes("tray.green.ico")
	}
	return app.Assets().GetBytes("tray.blue.ico")
}

// Items implements tray.Interface.
func (*Root) Items() []tray.Interface {
	panic("unimplemented")
}

// Title implements tray.Interface.
func (*Root) Title() string {
	return ""
}

// Tooltip implements tray.Interface.
func (*Root) Tooltip() string {
	panic("unimplemented")
}

func NewRoot() tray.Interface {
	return &Root{}
}
