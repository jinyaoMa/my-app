package tray

import (
	"context"
	_ "embed"
	"log"
	"my-app/backend/i18n"
	"my-app/backend/tray/menus"
	"sync"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed icons/icon.ico
var icon []byte

//go:embed icons/open-window.ico
var iconOpenWindow []byte

//go:embed icons/api-start.ico
var iconApiStart []byte

//go:embed icons/api-stop.ico
var iconApiStop []byte

var (
	once     sync.Once
	instance *tray
)

type tray struct {
	wailsCtx        context.Context
	openWindow      *menus.OpenWindow
	apiService      *menus.ApiService
	displayLanguage *menus.DisplayLanguage
	colorTheme      *menus.ColorTheme
	quit            *menus.Quit
}

func Tray() *tray {
	once.Do(func() {
		instance = &tray{}
		systray.Register(instance.onReady, instance.onQuit)
	})
	return instance
}

func (t *tray) SetWailsContext(ctx context.Context) *tray {
	t.wailsCtx = ctx
	return t
}

func (t *tray) ChangeLanguage(lang string) *tray {
	switch lang {
	case i18n.Zh:
		t.displayLanguage.ClickChinese()
	case i18n.En:
		t.displayLanguage.ClickEnglish()
	}
	return t
}

func (t *tray) ChangeTheme(theme string) *tray {
	switch theme {
	case menus.ColorThemeSystem:
		t.colorTheme.ClickSystem()
	case menus.ColorThemeLight:
		t.colorTheme.ClickLight()
	case menus.ColorThemeDark:
		t.colorTheme.ClickDark()
	}
	return t
}

func (t *tray) updateLocales() {
	locale := i18n.I18n().Locale()
	systray.SetTitle(locale.AppName)
	systray.SetTooltip(locale.AppName)
	t.openWindow.SetLocale(locale)
	t.apiService.SetLocale(locale)
	t.displayLanguage.SetLocale(locale)
	t.colorTheme.SetLocale(locale)
	t.quit.SetLocale(locale)
}

func (t *tray) onReady() {
	systray.SetTemplateIcon(icon, icon)

	t.openWindow = menus.
		NewOpenWindow().
		SetIcon(iconOpenWindow, iconOpenWindow).
		Watch(menus.OpenWindowListener{
			OnOpenWindow: func() {
				runtime.Show(t.wailsCtx)
			},
		})

	systray.AddSeparator()

	t.apiService = menus.
		NewApiService().
		SetIconStart(iconApiStart, iconApiStart).
		SetIconStop(iconApiStop, iconApiStop).
		Watch(menus.ApiServiceListener{
			OnStart: func() bool {
				return true
			},
			OnStop: func() bool {
				return true
			},
			OnOpenSwagger: func() {
				runtime.BrowserOpenURL(
					t.wailsCtx,
					"https://www.baidu.com",
				)
			},
		})

	systray.AddSeparator()

	t.displayLanguage = menus.
		NewDisplayLanguage().
		Watch(menus.DisplayLanguageListener{
			OnDisplayLanguageChanged: func(lang string) bool {
				locale := i18n.I18n().Change(lang).Locale()
				runtime.WindowSetTitle(t.wailsCtx, locale.AppName)
				runtime.EventsEmit(t.wailsCtx, "onLanguageChanged", lang)
				t.updateLocales()
				return true
			},
		})

	systray.AddSeparator()

	t.colorTheme = menus.
		NewColorTheme().
		Watch(menus.ColorThemeListener{
			OnColorThemeChanged: func(theme string) bool {
				switch theme {
				case menus.ColorThemeSystem:
					runtime.WindowSetSystemDefaultTheme(t.wailsCtx)
				case menus.ColorThemeLight:
					runtime.WindowSetLightTheme(t.wailsCtx)
				case menus.ColorThemeDark:
					runtime.WindowSetDarkTheme(t.wailsCtx)
				}
				return true
			},
		})

	systray.AddSeparator()

	t.quit = menus.
		NewQuit().
		Watch(menus.QuitListener{
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
					Icon:          icon,
				})
				if err != nil {
					log.Fatalf("fail to open quit dialog: %+v\n", err)
				}
				if dialog == "Yes" || dialog == locale.QuitDialog.DefaultButton {
					// when "Yes" or default button is clicked
					systray.Quit()
				}
			},
		})
}

func (t *tray) onQuit() {
	// end menus properly
	t.openWindow.StopWatch()
	t.apiService.StopWatch()
	t.displayLanguage.StopWatch()
	t.colorTheme.StopWatch()
	t.quit.StopWatch()

	runtime.Quit(t.wailsCtx)
}
