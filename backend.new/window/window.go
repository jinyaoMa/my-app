package window

import (
	"embed"
	"my-app/backend.new/app"
	"my-app/backend.new/app/types"
	"my-app/backend.new/model"
	"my-app/backend.new/services"
	"my-app/backend.new/tray"
	"my-app/backend.new/utils"
	"os"
	"sync"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed assets
var assets embed.FS

var (
	instance *window
	once     sync.Once
)

type window struct{}

func Window() *window {
	once.Do(func() {
		instance = &window{}
		app.App().Log().Wails().Print("WINDOW INSTANCE INITIALIZED")
	})
	return instance
}

func (w *window) Run() {
	// default wails options
	opts := &options.App{
		Title:             services.Services().General().GetAppName(),
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
		Assets:             nil,
		AssetsHandler:      nil,
		Menu:               nil,
		Logger:             app.App().Log().Wails(),
		LogLevel:           logger.INFO,
		LogLevelProduction: logger.ERROR,
		OnStartup:          w.startup,
		OnDomReady:         w.domReady,
		OnShutdown:         w.shutdown,
		OnBeforeClose:      w.beforeClose,
		Bind:               append(services.Services().All(), tray.Tray()),
		WindowStartState:   options.Normal,
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
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
			},*/
			Messages:         nil,
			ResizeDebounceMS: 0,
			OnSuspend:        w.suspend,
			OnResume:         w.resume,
		},
		Mac:          &mac.Options{},
		Linux:        &linux.Options{},
		Experimental: &options.Experimental{},
	}

	// configure wails options
	app.App().UseConfig(func(cfg *app.Config) {
		// get stored Assets directory
		dirAssets := cfg.Get(model.OptionNameDirAssets)
		if utils.Utils().HasDir(dirAssets) {
			opts.Assets = os.DirFS(dirAssets)
			app.App().Log().Wails().Print("WAILS LOAD ASSET FROM dirAssets: " + dirAssets)
		} else {
			opts.Assets = assets
			// extract assets into dirAssets
			assetHelper := utils.NewEmbedFS(assets, "assets")
			if err := assetHelper.Extract(dirAssets); err != nil {
				app.App().Log().Wails().Fatal("failed to extract embed assets into dirAssets (" + dirAssets + "): " + err.Error())
			}
			app.App().Log().Wails().Print("WAILS LOAD ASSET FROM embed: backend/window/assets")
		}
		// get stored UserData directory
		opts.Windows.WebviewUserDataPath = cfg.Get(model.OptionNameDirUserData)
		// get stored color theme
		switch cfg.Get(model.OptionNameColorTheme) {
		default:
			opts.Windows.Theme = windows.SystemDefault
		case types.ColorThemeLight.ToString():
			opts.Windows.Theme = windows.Light
		case types.ColorThemeDark.ToString():
			opts.Windows.Theme = windows.Dark
		}
	})

	if err := wails.Run(opts); err != nil {
		app.App().Log().Wails().Fatal("failed to run wails: " + err.Error())
	}
}
