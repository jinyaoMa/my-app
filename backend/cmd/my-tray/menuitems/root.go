package menuitems

import (
	"context"
	"fmt"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"

	"github.com/wailsapp/wails/v2/pkg/options/windows"
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
	T := app.T()

	webServiceState := T.APIService.Disabled
	if app.API().IsRunning() {
		webServiceState = T.APIService.Enabled
	}

	displayLanguageText := T.Lang.Text

	colorThemeText := T.ColorTheme.System
	switch app.THEME() {
	case windows.Light:
		colorThemeText = T.ColorTheme.Light
	case windows.Dark:
		colorThemeText = T.ColorTheme.Dark
	}

	return fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		T.AppName,
		T.APIService.Label+webServiceState,
		T.DisplayLanguage.Label+displayLanguageText,
		T.ColorTheme.Label+colorThemeText,
	)
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
