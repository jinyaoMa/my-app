package tray

import (
	"my-app/backend/app"
	"my-app/backend/app/config"
	"my-app/backend/pkg"
	"my-app/backend/tray/menus"

	"github.com/getlantern/systray"
)

func (t *tray) onReady() {
	cfg := app.App().Config()
	ct := app.App().CurrentTranslation()

	ico, _ := t.icons.GetFileBytes("icon.ico")
	icoOpenWindow, _ := t.icons.GetFileBytes("open-window.ico")
	icoOpenVitePress, _ := t.icons.GetFileBytes("open-vitepress.ico")
	icoOpenSwagger, _ := t.icons.GetFileBytes("open-swagger.ico")
	icoWebStart, _ := t.icons.GetFileBytes("web-start.ico")
	icoWebStop, _ := t.icons.GetFileBytes("web-stop.ico")

	systray.SetTemplateIcon(ico, ico)

	// open window menu
	t.openWindow = menus.NewItem(ct.OpenWindow, ct.OpenWindow, icoOpenWindow)

	systray.AddSeparator()

	// web service menu
	t.webService = menus.NewMenuSwitch([]string{
		MniWebServiceStart,
		MniWebServiceVitePress,
		MniWebServiceSwagger,
		MniWebServiceStop,
	}, map[string]menus.MenuSwitchOption{
		MniWebServiceStart: {
			Title:   ct.WebService.Start,
			Tooltip: ct.WebService.Start,
			Icon:    icoWebStart,
			Flag:    false,
		},
		MniWebServiceVitePress: {
			Title:   ct.WebService.VitePress,
			Tooltip: ct.WebService.VitePress,
			Icon:    icoOpenVitePress,
			Flag:    true,
		},
		MniWebServiceSwagger: {
			Title:   ct.WebService.Swagger,
			Tooltip: ct.WebService.Swagger,
			Icon:    icoOpenSwagger,
			Flag:    true,
		},
		MniWebServiceStop: {
			Title:   ct.WebService.Stop,
			Tooltip: ct.WebService.Stop,
			Icon:    icoWebStop,
			Flag:    true,
		},
	}, cfg.Web().IsAutoStart())

	systray.AddSeparator()

	// display language menu
	t.displayLanguage = menus.NewMenuSelect(ct.DisplayLanguage.Title, []string{
		config.DisplayLanguageEn,
		config.DisplayLanguageZh,
	}, map[string]menus.MenuSelectOption{
		config.DisplayLanguageEn: {
			Title:   ct.DisplayLanguage.En,
			Tooltip: ct.DisplayLanguage.En,
		},
		config.DisplayLanguageZh: {
			Title:   ct.DisplayLanguage.Zh,
			Tooltip: ct.DisplayLanguage.Zh,
		},
	}, cfg.DisplayLanguage)

	systray.AddSeparator()

	// color theme menu
	t.colorTheme = menus.NewMenuSelect(ct.ColorTheme.Title, []string{
		config.ColorThemeSystem,
		config.ColorThemeLight,
		config.ColorThemeDark,
	}, map[string]menus.MenuSelectOption{
		config.ColorThemeSystem: {
			Title:   ct.ColorTheme.System,
			Tooltip: ct.ColorTheme.System,
		},
		config.ColorThemeLight: {
			Title:   ct.ColorTheme.Light,
			Tooltip: ct.ColorTheme.Label,
		},
		config.ColorThemeDark: {
			Title:   ct.ColorTheme.Dark,
			Tooltip: ct.ColorTheme.Dark,
		},
	}, cfg.ColorTheme)

	systray.AddSeparator()

	systray.AddMenuItem(pkg.Copyright, pkg.Copyright).Disable()

	systray.AddSeparator()

	// quit menu
	t.quit = menus.NewItem(ct.Quit, ct.Quit)

	t.refreshTooltip()
	go t.listen()
}
