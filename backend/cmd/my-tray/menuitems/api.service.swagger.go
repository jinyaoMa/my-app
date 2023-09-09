package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/internal/crud"
	"my-app/backend/pkg/tray"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type apiServiceSwagger struct {
	*tray.MenuItem
}

// Visible implements IMenuItem.
func (*apiServiceSwagger) Visible() bool {
	return app.API().IsRunning()
}

// Key implements tray.IMenuItem.
func (*apiServiceSwagger) Key() string {
	return "api.service.swagger"
}

// CanClick implements tray.IMenuItem.
func (*apiServiceSwagger) CanClick() bool {
	return true
}

// OnClick implements tray.IMenuItem.
func (s *apiServiceSwagger) OnClick() (quit bool) {
	url, _, _ := app.OPTION().GetOrCreateByOptionName(crud.OptionNameWebSwagger, "https://localhost:10443/swagger/index.html")
	runtime.BrowserOpenURL(s.Ctx, url)
	return false
}

// Title implements tray.IMenuItem.
func (*apiServiceSwagger) Title() string {
	return app.T().APIService.Swagger
}

// Tooltip implements tray.IMenuItem.
func (*apiServiceSwagger) Tooltip() string {
	return app.T().APIService.Swagger
}

func newApiServiceSwagger(ctx context.Context) tray.IMenuItem {
	return &apiServiceSwagger{
		MenuItem: tray.NewMenuItem(ctx),
	}
}
