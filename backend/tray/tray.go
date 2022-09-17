package tray

import (
	"context"
	_ "embed"
	"fmt"
	"my-app/backend/app"
	"my-app/backend/pkg/i18n"
	"my-app/backend/pkg/utils"
	"my-app/backend/tray/menus"
	"my-app/backend/web"
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
		systray.Register(instance.onReady, instance.onExit)
	})
	return instance
}

func (t *tray) SetWailsContext(ctx context.Context) *tray {
	t.wailsCtx = ctx
	return t
}

func (t *tray) StartApiService() *tray {
	t.apiService.ClickStart()
	return t
}

func (t *tray) StopApiService() *tray {
	t.apiService.ClickStop()
	return t
}

func (t *tray) ChangeLanguage(lang string) *tray {
	switch lang {
	case i18n.Zh:
		t.displayLanguage.ClickChinese()
	default:
		t.displayLanguage.ClickEnglish()
	}
	return t
}

func (t *tray) ChangeColorTheme(theme string) *tray {
	switch theme {
	case app.ColorThemeLight:
		t.colorTheme.ClickLight()
	case app.ColorThemeDark:
		t.colorTheme.ClickDark()
	default:
		t.colorTheme.ClickSystem()
	}
	return t
}

func (t *tray) onReady() {
	systray.SetTemplateIcon(icon, icon)

	t.openWindow = menus.
		NewOpenWindow().
		SetIcon(iconOpenWindow, iconOpenWindow).
		Watch(t.openWindowListener())

	systray.AddSeparator()

	t.apiService = menus.
		NewApiService().
		SetIconStart(iconApiStart, iconApiStart).
		SetIconStop(iconApiStop, iconApiStop).
		Watch(t.apiServiceListener())

	systray.AddSeparator()

	t.displayLanguage = menus.
		NewDisplayLanguage().
		Watch(t.displayLanguageListener())

	systray.AddSeparator()

	t.colorTheme = menus.
		NewColorTheme().
		Watch(t.colorThemeListener())

	systray.AddSeparator()

	systray.AddMenuItem(utils.Copyright, utils.Copyright).Disable()

	systray.AddSeparator()

	t.quit = menus.
		NewQuit().
		Watch(t.quitListener())
}

func (t *tray) onExit() {
	// end menus properly
	t.openWindow.StopWatch()
	t.apiService.StopWatch()
	t.displayLanguage.StopWatch()
	t.colorTheme.StopWatch()
	t.quit.StopWatch()

	web.Web().Stop()
	runtime.Quit(t.wailsCtx)
}

func (t *tray) refreshTooltip() {
	locale := i18n.I18n().Locale()
	separator := ": "

	ApiServiceState := locale.ApiService.Disabled
	if t.apiService.IsEnabled() {
		ApiServiceState = locale.ApiService.Enabled
	}

	ColorThemeText := locale.ColorTheme.System
	switch t.colorTheme.CurrentTheme() {
	case app.ColorThemeLight:
		ColorThemeText = locale.ColorTheme.Light
	case app.ColorThemeDark:
		ColorThemeText = locale.ColorTheme.Dark
	}

	systray.SetTooltip(
		fmt.Sprintf(
			"%s\n%s\n%s\n%s",
			locale.AppName,
			locale.ApiService.Label+separator+ApiServiceState,
			locale.DisplayLanguage.Label+separator+locale.Lang.Text,
			locale.ColorTheme.Label+separator+ColorThemeText,
		),
	)
}
