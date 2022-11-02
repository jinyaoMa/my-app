package tray

import (
	"context"
	"fmt"
	"my-app/backend.new/app"
	"my-app/backend.new/app/i18n"
	"my-app/backend.new/app/types"
	"my-app/backend.new/model"
	"my-app/backend.new/tray/menus"
	"my-app/backend.new/web"
	"reflect"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (t *tray) updateLanguage() *tray {
	app.App().UseContextAndI18n(func(ctx context.Context, T func() *i18n.Translation, i18n *i18n.I18n) {
		// update taskbar
		runtime.WindowSetTitle(ctx, T().AppName)

		// update tray icon
		systray.SetTitle(T().AppName)
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
	})
	return t
}

func (t *tray) updateIconTooltip() *tray {
	app.App().UseConfigAndI18n(func(cfg *app.Config, T func() *i18n.Translation, i18n *i18n.I18n) {
		separator := ": "

		webServiceState := T().WebService.Disabled
		if web.Web().IsRunning() {
			webServiceState = T().WebService.Enabled
		}

		displayLanguageText := T().Lang.Text

		colorThemeText := T().ColorTheme.System
		switch cfg.Get(model.OptionNameColorTheme) {
		case types.ColorThemeLight.ToString():
			colorThemeText = T().ColorTheme.Light
		case types.ColorThemeDark.ToString():
			colorThemeText = T().ColorTheme.Dark
		}

		systray.SetTooltip(
			fmt.Sprintf(
				"%s\n%s\n%s\n%s",
				T().AppName,
				T().WebService.Label+separator+webServiceState,
				T().DisplayLanguage.Label+separator+displayLanguageText,
				T().ColorTheme.Label+separator+colorThemeText,
			),
		)
	})
	return t
}
