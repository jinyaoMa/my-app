package main

import (
	"embed"
	"log"

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
	wlc := DefaultWailsLifeCycle()

	err := wails.Run(&options.App{
		Title:             "My Application",
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
		// RGBA:               &options.RGBA{},
		Assets:             frontend,
		AssetsHandler:      nil,
		Menu:               nil,
		Logger:             nil,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		OnStartup:          wlc.startup,
		OnDomReady:         wlc.domReady,
		OnShutdown:         wlc.shutdown,
		OnBeforeClose:      wlc.beforeClose,
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
			OnSuspend:                         wlc.suspend,
			OnResume:                          wlc.resume,
		},
		Mac:          &mac.Options{},
		Linux:        &linux.Options{},
		Experimental: &options.Experimental{},
	})

	if err != nil {
		log.Fatalf("fail to run wails: %+v\n", err)
	}
}
