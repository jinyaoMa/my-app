package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type apiServiceStop struct {
	*tray.MenuItem
}

// Visible implements IMenuItem.
func (*apiServiceStop) Visible() bool {
	return app.API().IsRunning()
}

// Key implements tray.IMenuItem.
func (*apiServiceStop) Key() string {
	return "api.service.stop"
}

// CanClick implements tray.IMenuItem.
func (*apiServiceStop) CanClick() bool {
	return true
}

// OnClick implements tray.IMenuItem.
func (s *apiServiceStop) OnClick() (quit bool) {
	if app.API().Stop(func() {
		println("try to stop...")
	}) {
		tray.Update(_root)
	}
	return false
}

// Title implements tray.IMenuItem.
func (*apiServiceStop) Title() string {
	return app.T().APIService.Stop
}

// Tooltip implements tray.IMenuItem.
func (*apiServiceStop) Tooltip() string {
	return app.T().APIService.Stop
}

func newApiServiceStop(ctx context.Context) tray.IMenuItem {
	return &apiServiceStop{
		MenuItem: tray.NewMenuItem(ctx),
	}
}
