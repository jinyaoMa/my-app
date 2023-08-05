package menuitem

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type root struct {
	*tray.MenuItemBase
}

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

// Title implements tray.Interface.
func (*root) Title() string {
	return app.T().AppName
}

// Tooltip implements tray.Interface.
func (*root) Tooltip() string {
	return app.T().AppName
}

func newRoot(ctx context.Context) tray.IMenuItemBase {
	return &root{
		MenuItemBase: &tray.MenuItemBase{
			Ctx: ctx,
			MenuItems: []tray.IMenuItem{
				newOpenWindow(ctx),
				newSeparator(ctx),
				newColorTheme(ctx),
				newQuit(ctx),
			},
		},
	}
}
