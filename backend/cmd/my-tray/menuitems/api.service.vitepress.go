package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/internal/crud"
	"my-app/backend/pkg/tray"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type apiServiceVitePress struct {
	*tray.MenuItem
}

// Visible implements IMenuItem.
func (*apiServiceVitePress) Visible() bool {
	return app.SERVER().IsRunning()
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
	url, _, _ := app.OPTION().GetOrSaveByOptionName(crud.OptionNameWebVitePress, "https://localhost:10443/docs/index.html")
	runtime.BrowserOpenURL(s.Ctx, url)
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
