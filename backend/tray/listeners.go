package tray

import (
	"fmt"
	"my-app/backend/app"
	"my-app/backend/model"
	"my-app/backend/pkg/i18n"
	"my-app/backend/tray/menus"
	"my-app/backend/web"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (t *tray) openWindowListener() menus.OpenWindowListener {
	return menus.OpenWindowListener{
		OnOpenWindow: func() {
			runtime.Show(t.wailsCtx)
		},
	}
}

func (t *tray) webServiceListener() menus.WebServiceListener {
	return menus.WebServiceListener{
		OnOpenVitePress: func() {
			runtime.BrowserOpenURL(
				t.wailsCtx,
				fmt.Sprintf("https://localhost%s/docs/", app.App().WebConfig().PortHttps),
			)
		},
		OnOpenSwagger: func() {
			runtime.BrowserOpenURL(
				t.wailsCtx,
				fmt.Sprintf("https://localhost%s/swagger/index.html", app.App().WebConfig().PortHttps),
			)
		},
		OnStart: func() (ok bool, complete func()) {
			return web.Web().Start(), func() {
				t.refreshTooltip()
			}
		},
		OnStop: func() (ok bool, complete func()) {
			return web.Web().Stop(), func() {
				t.refreshTooltip()
			}
		},
	}
}

func (t *tray) displayLanguageListener() menus.DisplayLanguageListener {
	return menus.DisplayLanguageListener{
		OnDisplayLanguageChanged: func(lang string) (ok bool, complete func()) {
			locale := i18n.I18n().Change(lang).Locale()

			runtime.WindowSetTitle(t.wailsCtx, locale.AppName)
			runtime.EventsEmit(t.wailsCtx, "onDisplayLanguageChanged", lang)

			systray.SetTitle(locale.AppName)

			t.openWindow.SetLocale()
			t.webService.SetLocale()
			t.displayLanguage.SetLocale()
			t.colorTheme.SetLocale()
			t.quit.SetLocale()

			option := model.MyOption{
				Name: app.CfgDisplayLanguage,
			}
			result := option.Update(locale.Lang.Code)
			if result.Error != nil {
				app.App().TrayLog().Fatalf("failed to update language option: %+v\n", result.Error)
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
			switch theme {
			case app.ColorThemeLight:
				runtime.WindowSetLightTheme(t.wailsCtx)
			case app.ColorThemeDark:
				runtime.WindowSetDarkTheme(t.wailsCtx)
			default:
				runtime.WindowSetSystemDefaultTheme(t.wailsCtx)
			}
			runtime.EventsEmit(t.wailsCtx, "onColorThemeChanged", theme)

			option := model.MyOption{
				Name: app.CfgColorTheme,
			}
			result := option.Update(theme)
			if result.Error != nil {
				app.App().TrayLog().Fatalf("failed to update theme option: %+v\n", result.Error)
			}

			return true, func() {
				t.refreshTooltip()
				runtime.Show(t.wailsCtx)
			}
		},
	}
}

func (t *tray) quitListener() menus.QuitListener {
	return menus.QuitListener{
		OnQuit: func() {
			locale := i18n.I18n().Locale()
			dialog, err := runtime.MessageDialog(t.wailsCtx, runtime.MessageDialogOptions{
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
