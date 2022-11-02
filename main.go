package main

import "my-app/backend.new/window"

/*
	"my-app/backend/app"
	"my-app/backend/pkg/utils"
	"my-app/backend/service"
	"my-app/backend/tray"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
*/

func main() {
	window.Window().Run()
	/*
		w := &wailsapp{}

		err := wails.Run(&options.App{
			Title:             w.title(),
			Width:             1024, // 16:10
			Height:            640,  // 16:10
			DisableResize:     false,
			Fullscreen:        false,
			Frameless:         false,
			MinWidth:          1024, // 16:10
			MinHeight:         640,  // 16:10
			MaxWidth:          -1,
			MaxHeight:         -1,
			StartHidden:       false,
			HideWindowOnClose: true,
			AlwaysOnTop:       false,
			// BackgroundColour:  &options.RGBA{R: 242, G: 242, B: 242, A: 0},
			// RGBA:              &options.RGBA{},
			Assets:             os.DirFS(utils.GetExecutablePath("Assets")),
			AssetsHandler:      nil,
			Menu:               nil,
			Logger:             app.App().Log().Wails(),
			LogLevel:           logger.INFO,
			LogLevelProduction: logger.ERROR,
			OnStartup:          w.startup,
			OnDomReady:         w.domReady,
			OnShutdown:         w.shutdown,
			OnBeforeClose:      w.beforeClose,
			Bind: []interface{}{
				tray.Tray(),
				service.Service(),
			},
			WindowStartState: options.Normal,
			Windows: &windows.Options{
				WebviewIsTransparent:              true,
				WindowIsTranslucent:               false,
				DisableWindowIcon:                 false,
				DisableFramelessWindowDecorations: false,
				WebviewUserDataPath:               utils.GetExecutablePath("UserData"),
				WebviewBrowserPath:                "",
				Theme:                             w.windowTheme(),
				CustomTheme:                       nil, /*&windows.ThemeSettings{
					// Theme to use when window is active
					DarkModeTitleBar:   windows.RGB(255, 0, 0), // Red
					DarkModeTitleText:  windows.RGB(0, 255, 0), // Green
					DarkModeBorder:     windows.RGB(0, 0, 255), // Blue
					LightModeTitleBar:  windows.RGB(200, 200, 200),
					LightModeTitleText: windows.RGB(20, 20, 20),
					LightModeBorder:    windows.RGB(200, 200, 200),
					// Theme to use when window is inactive
					DarkModeTitleBarInactive:   windows.RGB(128, 0, 0),
					DarkModeTitleTextInactive:  windows.RGB(0, 128, 0),
					DarkModeBorderInactive:     windows.RGB(0, 0, 128),
					LightModeTitleBarInactive:  windows.RGB(100, 100, 100),
					LightModeTitleTextInactive: windows.RGB(10, 10, 10),
					LightModeBorderInactive:    windows.RGB(100, 100, 100),
				},*
				Messages:         nil,
				ResizeDebounceMS: 0,
				OnSuspend:        w.suspend,
				OnResume:         w.resume,
			},
			Mac:          &mac.Options{},
			Linux:        &linux.Options{},
			Experimental: &options.Experimental{},
		})

		if err != nil {
			app.App().Log().Wails().Fatal("fail to run wails: " + err.Error())
		}
	*/
}
