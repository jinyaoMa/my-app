package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type apiService struct {
	*tray.MenuItem
}

// Key implements tray.IMenuItem.
func (*apiService) Key() string {
	return "api.service"
}

// Title implements tray.IMenuItem.
func (*apiService) Title() string {
	return app.T().APIService.Title
}

// Tooltip implements tray.IMenuItem.
func (*apiService) Tooltip() string {
	return app.T().APIService.Title
}

func newAPIService(ctx context.Context) tray.IMenuItem {
	return &apiService{
		MenuItem: tray.NewMenuItem(ctx,
			newApiServiceStart(ctx),
			newApiServiceStop(ctx),
			newSeparator(ctx),
			newApiServiceSwagger(ctx),
			newApiServiceVitePress(ctx)),
	}
}
