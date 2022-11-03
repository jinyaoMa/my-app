package tray

import (
	"embed"
	"sync"

	"my-app/backend.new/app"
	"my-app/backend.new/app/types"
	"my-app/backend.new/tray/menus"
	"my-app/backend.new/utils"
	"my-app/backend.new/web"

	"github.com/getlantern/systray"
)

//go:embed icons
var icons embed.FS

var _tray = &tray{}

type tray struct {
	once            sync.Once
	openWindow      *menus.SingleItem
	webService      *menus.SwitchGroup
	displayLanguage *menus.SelectList
	colorTheme      *menus.SelectList
	copyright       *menus.SingleItem
	quit            *menus.SingleItem
}

// systray
func Tray() *tray {
	_tray.once.Do(func() {
		waitUntilReady := make(chan struct{}) // wait until menus built completely
		systray.Register(func() {
			// load icons
			assetHelper := utils.NewEmbedFS(icons, "icons")
			ico, _ := assetHelper.GetFileBytes("icon.ico")
			icoOpenWindow, _ := assetHelper.GetFileBytes("open-window.ico")
			icoOpenVitePress, _ := assetHelper.GetFileBytes("open-vitepress.ico")
			icoOpenSwagger, _ := assetHelper.GetFileBytes("open-swagger.ico")
			icoWebStop, _ := assetHelper.GetFileBytes("web-stop.ico")
			icoWebStart, _ := assetHelper.GetFileBytes("web-start.ico")

			// setup tray icon and menus
			T := app.App().T()
			i18n := app.App().I18n()

			// tray icon
			systray.SetTitle(T.AppName)
			systray.SetTemplateIcon(ico, ico)
			_tray.updateIconTooltip()

			// open window menu
			_tray.openWindow = menus.NewSingleItem(
				MenuIdOpenWindow, T.OpenWindow, icoOpenWindow,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T.OpenWindow)
			})

			systray.AddSeparator()

			// web service menu
			_tray.webService = menus.NewSwitchGroup(
				web.Web().IsRunning(),
				3, 1,
			).AddItems2OnGroup(
				menus.NewSingleItem(
					MenuIdOpenVitePress, T.WebService.VitePress, icoOpenVitePress,
				).SetTextUpdater(func(updateText func(text string)) {
					updateText(T.WebService.VitePress)
				}),
				menus.NewSingleItem(
					MenuIdOpenSwagger, T.WebService.Swagger, icoOpenSwagger,
				).SetTextUpdater(func(updateText func(text string)) {
					updateText(T.WebService.Swagger)
				}),
				menus.NewSingleItem(
					MenuIdStopWeb, T.WebService.Stop, icoWebStop,
				).SetTextUpdater(func(updateText func(text string)) {
					updateText(T.WebService.Stop)
				}),
			).AddItems2OffGroup(
				menus.NewSingleItem(
					MenuIdStartWeb, T.WebService.Start, icoWebStart,
				).SetTextUpdater(func(updateText func(text string)) {
					updateText(T.WebService.Start)
				}),
			)

			systray.AddSeparator()

			// display language menu
			_tray.displayLanguage = menus.NewSelectList(
				menus.NewSingleItem(
					MenuIdDisplayLanguage, T.DisplayLanguage.Title,
				).SetTextUpdater(func(updateText func(text string)) {
					updateText(T.DisplayLanguage.Title)
				}),
				len(i18n.AvailableLanguages()),
			)
			cLang := i18n.ParseLanguage(app.App().Cfg().Get(types.ConfigNameDisplayLanguage))
			for _, lang := range i18n.AvailableLanguages() {
				opt := menus.NewSingleItem(
					lang.ToString(), i18n.Translation(lang).Lang.Text,
				).SetTextUpdater(func(updateText func(text string)) {
					updateText(i18n.Translation(lang).Lang.Text)
				})
				if cLang == lang {
					opt.Check()
				}
				_tray.displayLanguage.AddOption(opt)
			}

			systray.AddSeparator()

			// color theme menu
			_tray.colorTheme = menus.NewSelectList(
				menus.NewSingleItem(
					MenuIdColorTheme, T.ColorTheme.Title,
				).SetTextUpdater(func(updateText func(text string)) {
					updateText(T.ColorTheme.Title)
				}),
				3,
			).AddOptions(
				menus.NewSingleItem(
					types.ColorThemeDefault.ToString(), T.ColorTheme.System,
				).SetTextUpdater(func(updateText func(text string)) {
					updateText(T.ColorTheme.System)
				}),
				menus.NewSingleItem(
					types.ColorThemeLight.ToString(), T.ColorTheme.Light,
				).SetTextUpdater(func(updateText func(text string)) {
					updateText(T.ColorTheme.Light)
				}),
				menus.NewSingleItem(
					types.ColorThemeDark.ToString(), T.ColorTheme.Dark,
				).SetTextUpdater(func(updateText func(text string)) {
					updateText(T.ColorTheme.Dark)
				}),
			).Check(types.ParseColorTheme(app.App().Cfg().Get(types.ConfigNameColorTheme)).ToString())

			systray.AddSeparator()

			// copyright menu
			_tray.copyright = menus.NewSingleItem(
				MenuIdCopyright, utils.Copyright,
			).Disable()

			systray.AddSeparator()

			// quit menu
			_tray.quit = menus.NewSingleItem(
				MenuIdQuit, T.Quit,
			).SetTextUpdater(func(updateText func(text string)) {
				updateText(T.Quit)
			})

			_tray.watch()

			app.App().Log().Tray().Println("TRAY IS READY")
			close(waitUntilReady)
		}, nil)
		<-waitUntilReady
	})
	return _tray
}
