package tray

import (
	"embed"
	"sync"

	"my-app/backend.new/app"
	"my-app/backend.new/app/i18n"
	"my-app/backend.new/model"
	"my-app/backend.new/tray/menus"
	"my-app/backend.new/utils"

	"github.com/getlantern/systray"
)

//go:embed icons
var icons embed.FS

var (
	instance *tray
	once     sync.Once
)

type tray struct {
	openWindow      *menus.SingleItem
	webService      *menus.SwitchGroup
	displayLanguage *menus.SelectList
	colorTheme      *menus.SelectList
	copyright       *menus.SingleItem
	quit            *menus.SingleItem
}

func Tray() *tray {
	once.Do(func() {
		instance = &tray{}
		systray.Register(instance.onReady, nil)
	})
	return instance
}

func (t *tray) onReady() {
	// load icons
	assetHelper := utils.NewEmbedFS(icons, "icons")
	ico, _ := assetHelper.GetFileBytes("icon.ico")
	icoOpenWindow, _ := assetHelper.GetFileBytes("open-window.ico")
	icoOpenVitePress, _ := assetHelper.GetFileBytes("open-vitepress.ico")
	icoOpenSwagger, _ := assetHelper.GetFileBytes("open-swagger.ico")
	icoWebStop, _ := assetHelper.GetFileBytes("web-stop.ico")
	icoWebStart, _ := assetHelper.GetFileBytes("web-start.ico")

	// setup menus
	systray.SetTemplateIcon(ico, ico)
	app.App().UseConfigAndI18n(func(cfg *app.Config, T func() *i18n.Translation, i18n *i18n.I18n) {
		// open window menu
		t.openWindow = menus.NewSingleItem(
			"OpenWindow", T().OpenWindow, icoOpenWindow,
		).SetTextUpdater(func(updateText func(text string)) {
			updateText(T().OpenWindow)
		})

		systray.AddSeparator()

		// web service menu
		t.webService = menus.NewSwitchGroup(
			cfg.Get(model.OptionNameWebAutoStart) == string(app.ConfigOptionTrue),
			3, 1,
		).AddItems2OnGroup(
			menus.NewSingleItem(
				"OpenVitePress", T().WebService.VitePress, icoOpenVitePress,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T().WebService.VitePress)
			}),
			menus.NewSingleItem(
				"OpenSwagger", T().WebService.Swagger, icoOpenSwagger,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T().WebService.Swagger)
			}),
			menus.NewSingleItem(
				"WebStop", T().WebService.Stop, icoWebStop,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T().WebService.Stop)
			}),
		).AddItems2OffGroup(
			menus.NewSingleItem(
				"WebStart", T().WebService.Start, icoWebStart,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T().WebService.Start)
			}),
		)

		systray.AddSeparator()

		// display language menu
		t.displayLanguage = menus.NewSelectList(
			menus.NewSingleItem(
				"DisplayLanguage", T().DisplayLanguage.Title,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T().DisplayLanguage.Title)
			}),
			len(i18n.AvailableLanguages()),
		)
		cLang := cfg.Get(model.OptionNameDisplayLanguage)
		for _, lang := range i18n.AvailableLanguages() {
			opt := menus.NewSingleItem(
				lang, i18n.Translation(lang).Lang.Text,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(i18n.Translation(lang).Lang.Text)
			})
			if cLang == lang {
				opt.Check()
			}
			t.displayLanguage.AddOption(opt)
		}

		systray.AddSeparator()

		// color theme menu
		t.colorTheme = menus.NewSelectList(
			menus.NewSingleItem(
				"ColorTheme", T().ColorTheme.Title,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T().ColorTheme.Title)
			}),
			3,
		).AddOptions(
			menus.NewSingleItem(
				string(app.ConfigOptionColorThemeSystem), T().ColorTheme.System,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T().ColorTheme.System)
			}),
			menus.NewSingleItem(
				string(app.ConfigOptionColorThemeLight), T().ColorTheme.Light,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T().ColorTheme.Light)
			}),
			menus.NewSingleItem(
				string(app.ConfigOptionColorThemeDark), T().ColorTheme.Dark,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T().ColorTheme.Dark)
			}),
		).Check(cfg.Get(model.OptionNameColorTheme))

		systray.AddSeparator()

		// copyright menu
		t.copyright = menus.NewSingleItem(
			"Copyright", utils.Copyright,
		).Disable()

		systray.AddSeparator()

		// quit menu
		t.quit = menus.NewSingleItem(
			"Quit", T().Quit,
		).SetTextUpdater(func(updateText func(text string)) {
			updateText(T().Quit)
		})
	})

	go t.watch()
}
