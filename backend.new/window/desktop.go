package window

import (
	"my-app/backend.new/app"
	"my-app/backend.new/model"
	"my-app/backend.new/services"
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

var (
	instance *window
	once     sync.Once
)

type window struct {
	bind *Bind
}

func Window() *window {
	once.Do(func() {
		instance = &window{
			bind: NewBind(),
		}
	})
	return instance
}

func (w *window) Run() {
	// default wails options
	opts := &options.App{
		Title:             w.bind.GetAppName(),
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
		Bind: append([]interface{}{
			w.bind,
		}, services.Services().All()...),
		WindowStartState: options.Normal,
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
	app.App().UseCfg(func(cfg *app.Config) {
		// get stored Assets directory
		opts.Assets = os.DirFS(cfg.Get(model.OptionDirAssets))
		// get stored UserData directory
		opts.Windows.WebviewUserDataPath = cfg.Get(model.OptionDirUserData)
		// get stored color theme
		switch cfg.Get(model.OptionColorTheme) {
		case utils.ColorThemeSystem.ToString():
			opts.Windows.Theme = windows.SystemDefault
		case utils.ColorThemeLight.ToString():
			opts.Windows.Theme = windows.Light
		case utils.ColorThemeDark.ToString():
			opts.Windows.Theme = windows.Dark
		}
	})

	if err := wails.Run(opts); err != nil {
		app.App().Log().Wails().Fatal("fail to run wails: " + err.Error())
	}
}
