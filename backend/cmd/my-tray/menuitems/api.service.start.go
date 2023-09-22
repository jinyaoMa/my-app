package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type apiServiceStart struct {
	*tray.MenuItem
}

// Visible implements IMenuItem.
func (*apiServiceStart) Visible() bool {
	return !app.SERVER().IsRunning()
}

// Key implements tray.IMenuItem.
func (*apiServiceStart) Key() string {
	return "api.service.start"
}

// CanClick implements tray.IMenuItem.
func (*apiServiceStart) CanClick() bool {
	return true
}

// OnClick implements tray.IMenuItem.
func (s *apiServiceStart) OnClick() (quit bool) {
	if app.StartAPI() {
		tray.Update(_root)
	}
	return false
}

// Title implements tray.IMenuItem.
func (*apiServiceStart) Title() string {
	return app.T().APIService.Start
}

// Tooltip implements tray.IMenuItem.
func (*apiServiceStart) Tooltip() string {
	return app.T().APIService.Start
}

func newApiServiceStart(ctx context.Context) tray.IMenuItem {
	return &apiServiceStart{
		MenuItem: tray.NewMenuItem(ctx),
	}
}
