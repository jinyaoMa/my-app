package tray

import (
	"fmt"
	"my-app/backend/app"
	"my-app/backend/pkg/i18n"
	"my-app/backend/service"
	"my-app/backend/tray/menus"
	"my-app/backend/web"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (t *tray) openWindowListener() menus.OpenWindowListener {
	return menus.OpenWindowListener{
		OnOpenWindow: func() {
			runtime.Show(app.App().WailsContext())
		},
	}
}

func (t *tray) webServiceListener() menus.WebServiceListener {
	return menus.WebServiceListener{
		OnOpenVitePress: func() {
			runtime.BrowserOpenURL(
				app.App().WailsContext(),
				fmt.Sprintf("https://localhost%s/docs/", app.App().WebConfig().PortHttps),
			)
		},
		OnOpenSwagger: func() {
			runtime.BrowserOpenURL(
				app.App().WailsContext(),
				fmt.Sprintf("https://localhost%s/swagger/index.html", app.App().WebConfig().PortHttps),
			)
		},
		OnStart: func() (ok bool, complete func()) {
			return web.Web().Start(), func() {
				t.refreshTooltip()
				runtime.EventsEmit(app.App().WailsContext(), "onWebServiceChanged", true)
			}
		},
		OnStop: func() (ok bool, complete func()) {
			return web.Web().Stop(), func() {
				t.refreshTooltip()
				runtime.EventsEmit(app.App().WailsContext(), "onWebServiceChanged", false)
			}
		},
	}
}

func (t *tray) displayLanguageListener() menus.DisplayLanguageListener {
	return menus.DisplayLanguageListener{
		OnDisplayLanguageChanged: func(lang string) (ok bool, complete func()) {
			locale := i18n.I18n().Change(lang).Locale()
			ctx := app.App().WailsContext()

			runtime.WindowSetTitle(ctx, locale.AppName)
			runtime.EventsEmit(ctx, "onDisplayLanguageChanged", lang)

			systray.SetTitle(locale.AppName)

			t.openWindow.SetLocale()
			t.webService.SetLocale()
			t.displayLanguage.SetLocale()
			t.colorTheme.SetLocale()
			t.quit.SetLocale()

			if err := service.Settings().SaveOption(app.CfgDisplayLanguage, locale.Lang.Code); err != nil {
				app.App().TrayLog().Fatalf("failed to update language option: %+v\n", err)
			}

			return true, func() {
				t.refreshTooltip()
			}
		},
	}
}

func (t *tray) colorThemeListener() menus.ColorThemeListener {
	return menus.ColorThemeListener{
		OnColorThemeChanged: func(theme string) (ok bool, complete func()) {
			ctx := app.App().WailsContext()

			switch theme {
			case app.ColorThemeLight:
				runtime.WindowSetLightTheme(ctx)
			case app.ColorThemeDark:
				runtime.WindowSetDarkTheme(ctx)
			default:
				runtime.WindowSetSystemDefaultTheme(ctx)
			}
			runtime.EventsEmit(ctx, "onColorThemeChanged", theme)

			if err := service.Settings().SaveOption(app.CfgColorTheme, theme); err != nil {
				app.App().TrayLog().Fatalf("failed to update theme option: %+v\n", err)
			}

			return true, func() {
				t.refreshTooltip()
			}
		},
	}
}

func (t *tray) quitListener() menus.QuitListener {
	return menus.QuitListener{
		OnQuit: func() {
			locale := i18n.I18n().Locale()
			dialog, err := runtime.MessageDialog(app.App().WailsContext(), runtime.MessageDialogOptions{
				Type:    runtime.QuestionDialog,
				Title:   locale.AppName,
				Message: locale.QuitDialog.Message,
				Buttons: []string{
					locale.QuitDialog.DefaultButton,
					locale.QuitDialog.CancelButton,
				},
				DefaultButton: locale.QuitDialog.DefaultButton,
				CancelButton:  locale.QuitDialog.CancelButton,
				// Icon:          nil,
			})
			if err != nil {
				app.App().TrayLog().Fatalf("fail to open quit dialog: %+v\n", err)
			}
			if dialog == "Yes" || dialog == locale.QuitDialog.DefaultButton {
				// when "Yes" or default button is clicked
				systray.Quit()
			}
		},
	}
}
