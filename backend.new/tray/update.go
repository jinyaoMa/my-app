package tray

import (
	"fmt"
	"my-app/backend.new/app"
	"my-app/backend.new/app/types"
	"my-app/backend.new/tray/menus"
	"my-app/backend.new/web"
	"reflect"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// updateLanguage update display text with current translation
func (t *tray) updateLanguage() *tray {
	T := app.App().T()

	// update taskbar
	runtime.WindowSetTitle(app.App().Ctx(), T.AppName)

	// update tray icon
	systray.SetTitle(T.AppName)
	t.updateIconTooltip()

	// update menus
	ms := reflect.ValueOf(t)
	for i := 0; i < ms.NumField(); i++ {
		if ms.Field(i).CanInterface() {
			if menu, ok := ms.Field(i).Interface().(menus.IRefresh); ok {
				menu.UpdateText()
			}
		}
	}
	return t
}

// updateIconTooltip update the tooltip hint for tray icon on taskbar
func (t *tray) updateIconTooltip() *tray {
	T := app.App().T()

	webServiceState := T.WebService.Disabled
	if web.Web().IsRunning() {
		webServiceState = T.WebService.Enabled
	}

	displayLanguageText := T.Lang.Text

	colorThemeText := T.ColorTheme.System
	switch app.App().Cfg().Get(types.ConfigNameColorTheme) {
	case types.ColorThemeLight.ToString():
		colorThemeText = T.ColorTheme.Light
	case types.ColorThemeDark.ToString():
		colorThemeText = T.ColorTheme.Dark
	}

	separator := ": "
	systray.SetTooltip(
		fmt.Sprintf(
			"%s\n%s\n%s\n%s",
			T.AppName,
			T.WebService.Label+separator+webServiceState,
			T.DisplayLanguage.Label+separator+displayLanguageText,
			T.ColorTheme.Label+separator+colorThemeText,
		),
	)
	return t
}
