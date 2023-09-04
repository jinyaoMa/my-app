package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/api"
	"my-app/backend/pkg/tray"

	"github.com/gofiber/fiber/v2"
)

type apiServiceStart struct {
	*tray.MenuItem
}

// Visible implements IMenuItem.
func (*apiServiceStart) Visible() bool {
	return !app.API().IsRunning()
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
	if app.API().Start(&api.Config{
		IsDev: true,
		Log:   app.LOG(),
		Http: api.ConfigHttp{
			Port: 10080,
		},
		Https: api.ConfigHttps{
			Port: 10443,
		},
		Setup: func(app *fiber.App) *fiber.App {
			return app
		},
	}) {
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
