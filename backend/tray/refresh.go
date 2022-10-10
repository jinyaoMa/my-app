package tray

import (
	"fmt"
	"my-app/backend/app"
	"my-app/backend/app/config"
	"my-app/backend/web"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (t *tray) refreshLanguage() {
	ct := app.App().CurrentTranslation()

	runtime.WindowSetTitle(t.ctx, ct.AppName)
	systray.SetTitle(ct.AppName)

	t.openWindow.SetTitle(ct.OpenWindow)
	t.openWindow.SetTooltip(ct.OpenWindow)

	t.webService.FalseOptions[MniWebServiceStart].SetTitle(ct.WebService.Start)
	t.webService.FalseOptions[MniWebServiceStart].SetTooltip(ct.WebService.Start)
	t.webService.TrueOptions[MniWebServiceVitePress].SetTitle(ct.WebService.VitePress)
	t.webService.TrueOptions[MniWebServiceVitePress].SetTooltip(ct.WebService.VitePress)
	t.webService.TrueOptions[MniWebServiceSwagger].SetTitle(ct.WebService.Swagger)
	t.webService.TrueOptions[MniWebServiceSwagger].SetTooltip(ct.WebService.Swagger)
	t.webService.TrueOptions[MniWebServiceStop].SetTitle(ct.WebService.Stop)
	t.webService.TrueOptions[MniWebServiceStop].SetTooltip(ct.WebService.Stop)
	t.webService.Refresh()

	t.displayLanguage.Title.SetTitle(ct.DisplayLanguage.Title)
	t.displayLanguage.Title.SetTitle(ct.DisplayLanguage.Title)
	t.displayLanguage.Options[config.DisplayLanguageEn].SetTitle(ct.DisplayLanguage.En)
	t.displayLanguage.Options[config.DisplayLanguageEn].SetTooltip(ct.DisplayLanguage.En)
	t.displayLanguage.Options[config.DisplayLanguageZh].SetTitle(ct.DisplayLanguage.Zh)
	t.displayLanguage.Options[config.DisplayLanguageZh].SetTooltip(ct.DisplayLanguage.Zh)

	t.colorTheme.Title.SetTitle(ct.ColorTheme.Title)
	t.colorTheme.Title.SetTitle(ct.ColorTheme.Title)
	t.colorTheme.Options[config.ColorThemeSystem].SetTitle(ct.ColorTheme.System)
	t.colorTheme.Options[config.ColorThemeSystem].SetTooltip(ct.ColorTheme.System)
	t.colorTheme.Options[config.ColorThemeLight].SetTitle(ct.ColorTheme.Light)
	t.colorTheme.Options[config.ColorThemeLight].SetTooltip(ct.ColorTheme.Light)
	t.colorTheme.Options[config.ColorThemeDark].SetTitle(ct.ColorTheme.Dark)
	t.colorTheme.Options[config.ColorThemeDark].SetTooltip(ct.ColorTheme.Dark)

	t.quit.SetTitle(ct.Quit)
	t.quit.SetTooltip(ct.Quit)
}

func (t *tray) refreshTooltip() {
	ct := app.App().CurrentTranslation()
	separator := ": "

	var sebServiceState string
	if web.Web().IsRunning() {
		sebServiceState = ct.WebService.Enabled
	} else {
		sebServiceState = ct.WebService.Disabled
	}

	var displayLanguageText string
	switch app.App().Config().DisplayLanguage {
	default:
		displayLanguageText = ct.DisplayLanguage.En
	case config.DisplayLanguageZh:
		displayLanguageText = ct.DisplayLanguage.Zh
	}

	var colorThemeText string
	switch app.App().Config().ColorTheme {
	default:
		colorThemeText = ct.ColorTheme.System
	case config.ColorThemeLight:
		colorThemeText = ct.ColorTheme.Light
	case config.ColorThemeDark:
		colorThemeText = ct.ColorTheme.Dark
	}

	systray.SetTooltip(
		fmt.Sprintf(
			"%s\n%s\n%s\n%s",
			ct.AppName,
			ct.WebService.Label+separator+sebServiceState,
			ct.DisplayLanguage.Label+separator+displayLanguageText,
			ct.ColorTheme.Label+separator+colorThemeText,
		),
	)
}
