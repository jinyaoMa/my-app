package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type apiServiceVitePress struct {
	*tray.MenuItem
}

// Visible implements IMenuItem.
func (*apiServiceVitePress) Visible() bool {
	return app.API().IsRunning()
}

// Key implements tray.IMenuItem.
func (*apiServiceVitePress) Key() string {
	return "api.service.vite.press"
}

// CanClick implements tray.IMenuItem.
func (*apiServiceVitePress) CanClick() bool {
	return true
}

// OnClick implements tray.IMenuItem.
func (s *apiServiceVitePress) OnClick() (quit bool) {
	return false
}

// Title implements tray.IMenuItem.
func (*apiServiceVitePress) Title() string {
	return app.T().APIService.VitePress
}

// Tooltip implements tray.IMenuItem.
func (*apiServiceVitePress) Tooltip() string {
	return app.T().APIService.VitePress
}

func newApiServiceVitePress(ctx context.Context) tray.IMenuItem {
	return &apiServiceVitePress{
		MenuItem: tray.NewMenuItem(ctx),
	}
}
