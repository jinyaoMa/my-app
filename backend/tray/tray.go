package tray

import (
	"embed"
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

//go:embed icons
var icons embed.FS

var (
	once     sync.Once
	instance *tray
)

type tray struct {
	openWindow      *menus.OpenWindow
	webService      *menus.WebService
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

func (t *tray) IsWebServiceRunning() bool {
	return t.webService.IsEnabled()
}

func (t *tray) OpenVitePress() *tray {
	t.webService.ClickOpenVitePress()
	return t
}

func (t *tray) OpenSwagger() *tray {
	t.webService.ClickOpenSwagger()
	return t
}

func (t *tray) StartWebService() *tray {
	t.webService.ClickStart()
	return t
}

func (t *tray) StopWebService() *tray {
	t.webService.ClickStop()
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
	icon, _ := icons.ReadFile("icons/icon.ico")
	iconOpenWindow, _ := icons.ReadFile("icons/open-window.ico")
	iconOpenVitePress, _ := icons.ReadFile("icons/open-vitepress.ico")
	iconOpenSwagger, _ := icons.ReadFile("icons/open-swagger.ico")
	iconWebStart, _ := icons.ReadFile("icons/web-start.ico")
	iconWebStop, _ := icons.ReadFile("icons/web-stop.ico")

	systray.SetTemplateIcon(icon, icon)

	t.openWindow = menus.
		NewOpenWindow().
		SetIcon(iconOpenWindow, iconOpenWindow).
		Watch(t.openWindowListener())

	systray.AddSeparator()

	t.webService = menus.
		NewWebService().
		SetIconVitePress(iconOpenVitePress, iconOpenVitePress).
		SetIconSwagger(iconOpenSwagger, iconOpenSwagger).
		SetIconStart(iconWebStart, iconWebStart).
		SetIconStop(iconWebStop, iconWebStop).
		Watch(t.webServiceListener())

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
	t.webService.StopWatch()
	t.displayLanguage.StopWatch()
	t.colorTheme.StopWatch()
	t.quit.StopWatch()

	web.Web().Stop()
	runtime.Quit(app.App().WailsContext())
}

func (t *tray) refreshTooltip() {
	locale := i18n.I18n().Locale()
	separator := ": "

	WebServiceState := locale.WebService.Disabled
	if t.IsWebServiceRunning() {
		WebServiceState = locale.WebService.Enabled
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
			locale.WebService.Label+separator+WebServiceState,
			locale.DisplayLanguage.Label+separator+locale.Lang.Text,
			locale.ColorTheme.Label+separator+ColorThemeText,
		),
	)
}
