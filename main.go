package main

import (
	"embed"
	"log"
	"my-app/backend/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	wlc := app.App().WailsLifeCycle().Initialize()

	err := wails.Run(&options.App{
		Title:             "My Application",
		Width:             1024, // 16:9
		Height:            576,  // 16:9
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		MinWidth:          1024, // 16:9
		MinHeight:         576,  // 16:9
		MaxWidth:          -1,
		MaxHeight:         -1,
		StartHidden:       false,
		HideWindowOnClose: true,
		AlwaysOnTop:       false,
		// BackgroundColour:  &options.RGBA{R: 242, G: 242, B: 242, A: 0},
		// RGBA:               &options.RGBA{},
		Assets:             frontend,
		AssetsHandler:      nil,
		Menu:               nil,
		Logger:             nil,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		OnStartup:          wlc.Startup,
		OnDomReady:         wlc.DomReady,
		OnShutdown:         wlc.Shutdown,
		OnBeforeClose:      wlc.BeforeClose,
		Bind:               []interface{}{},
		WindowStartState:   options.Normal,
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
			CustomTheme:                       nil,
			TranslucencyType:                  windows.Auto,
			Messages:                          nil,
			ResizeDebounceMS:                  0,
			OnSuspend:                         wlc.Suspend,
			OnResume:                          wlc.Resume,
		},
		Mac:          &mac.Options{},
		Linux:        &linux.Options{},
		Experimental: &options.Experimental{},
	})

	if err != nil {
		log.Fatalf("fail to run wails: %+v\n", err)
	}
}
