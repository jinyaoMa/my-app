package tray

import (
	"embed"
	"sync"

	"my-app/backend/app"
	"my-app/backend/app/types"
	"my-app/backend/tray/menus"
	"my-app/backend/utils"
	"my-app/backend/web"

	"github.com/getlantern/systray"
)

// menu ids
const (
	MenuIdOpenWindow      = "OpenWindow"
	MenuIdOpenVitePress   = "OpenVitePress"
	MenuIdOpenSwagger     = "OpenSwagger"
	MenuIdStopWeb         = "StopWeb"
	MenuIdStartWeb        = "StartWeb"
	MenuIdDisplayLanguage = "DisplayLanguage"
	MenuIdColorTheme      = "ColorTheme"
	MenuIdCopyright       = "Copyright"
	MenuIdQuit            = "Quit"
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

			/* setup tray icon and menus */

			// tray icon
			systray.SetTitle(app.App().T().AppName)
			systray.SetTemplateIcon(ico, ico)
			_tray.updateIconTooltip()

			// open window menu
			_tray.openWindow = menus.NewSingleItem(
				MenuIdOpenWindow, app.App().T().OpenWindow, icoOpenWindow,
			).SetTextUpdater(func(id string) (updateText string) {
				return app.App().T().OpenWindow
			})

			systray.AddSeparator()

			// web service menu
			_tray.webService = menus.NewSwitchGroup(
				web.Web().IsRunning(),
				3, 1,
			).AddItems2OnGroup(
				menus.NewSingleItem(
					MenuIdOpenVitePress, app.App().T().WebService.VitePress, icoOpenVitePress,
				).SetTextUpdater(func(id string) (updateText string) {
					return app.App().T().WebService.VitePress
				}),
				menus.NewSingleItem(
					MenuIdOpenSwagger, app.App().T().WebService.Swagger, icoOpenSwagger,
				).SetTextUpdater(func(id string) (updateText string) {
					return app.App().T().WebService.Swagger
				}),
				menus.NewSingleItem(
					MenuIdStopWeb, app.App().T().WebService.Stop, icoWebStop,
				).SetTextUpdater(func(id string) (updateText string) {
					return app.App().T().WebService.Stop
				}),
			).AddItems2OffGroup(
				menus.NewSingleItem(
					MenuIdStartWeb, app.App().T().WebService.Start, icoWebStart,
				).SetTextUpdater(func(id string) (updateText string) {
					return app.App().T().WebService.Start
				}),
			)

			systray.AddSeparator()

			// display language menu
			_tray.displayLanguage = menus.NewSelectList(
				menus.NewSingleItem(
					MenuIdDisplayLanguage, app.App().T().DisplayLanguage.Title,
				).SetTextUpdater(func(id string) (updateText string) {
					return app.App().T().DisplayLanguage.Title
				}),
				len(app.App().I18n().AvailableLanguages()),
			)
			for _, lang := range app.App().I18n().AvailableLanguages() {
				opt := menus.NewSingleItem(
					lang, app.App().I18n().Translation(lang).Lang.Text,
				).SetTextUpdater(func(id string) (updateText string) {
					return app.App().I18n().Translation(id).Lang.Text
				})
				if app.App().Lang() == lang {
					opt.Check()
				}
				_tray.displayLanguage.AddOption(opt)
			}

			systray.AddSeparator()

			// color theme menu
			_tray.colorTheme = menus.NewSelectList(
				menus.NewSingleItem(
					MenuIdColorTheme, app.App().T().ColorTheme.Title,
				).SetTextUpdater(func(id string) (updateText string) {
					return app.App().T().ColorTheme.Title
				}),
				3,
			).AddOptions(
				menus.NewSingleItem(
					types.ColorThemeSystem.ToString(), app.App().T().ColorTheme.System,
				).SetTextUpdater(func(id string) (updateText string) {
					return app.App().T().ColorTheme.System
				}),
				menus.NewSingleItem(
					types.ColorThemeLight.ToString(), app.App().T().ColorTheme.Light,
				).SetTextUpdater(func(id string) (updateText string) {
					return app.App().T().ColorTheme.Light
				}),
				menus.NewSingleItem(
					types.ColorThemeDark.ToString(), app.App().T().ColorTheme.Dark,
				).SetTextUpdater(func(id string) (updateText string) {
					return app.App().T().ColorTheme.Dark
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
				MenuIdQuit, app.App().T().Quit,
			).SetTextUpdater(func(id string) (updateText string) {
				return app.App().T().Quit
			})

			_tray.watch()

			app.App().Log().Tray().Println("TRAY IS READY")
			close(waitUntilReady)
		}, nil)
		<-waitUntilReady
	})
	return _tray
}
