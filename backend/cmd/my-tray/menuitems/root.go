package menuitems

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
	if app.API().IsStopping() {
		return app.ASSETS().GetBytes("tray.orange.ico")
	}
	if app.API().IsRunning() {
		return app.ASSETS().GetBytes("tray.green.ico")
	}
	return app.ASSETS().GetBytes("tray.blue.ico")
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
		MenuItemBase: tray.NewMenuItemBase(ctx,
			newOpenWindow(ctx),
			newSeparator(ctx),
			newAPIService(ctx),
			newSeparator(ctx),
			newdisplayLanguage(ctx),
			newColorTheme(ctx),
			newQuit(ctx)),
	}
}
