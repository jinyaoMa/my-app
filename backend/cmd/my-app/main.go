package main

import (
	"my-app/backend/internal/app"
	"my-app/backend/internal/service"
	"my-app/frontend"

	"github.com/devfeel/mapper"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func main() {
	// Create an instance of the app structure
	wailsapp := NewApp()

	optionService := service.NewOption(app.DB(), mapper.NewMapper())

	// Create application with options

	err := wails.Run(&options.App{
		Title:         app.T().AppName,
		Width:         1024,
		Height:        720,
		DisableResize: false,
		// Fullscreen:         false,
		WindowStartState:  options.Normal,
		Frameless:         true,
		MinWidth:          1024,
		MinHeight:         600,
		MaxWidth:          1200,
		MaxHeight:         720,
		StartHidden:       false,
		HideWindowOnClose: true,
		BackgroundColour:  &options.RGBA{R: 0, G: 0, B: 0, A: 255},
		AlwaysOnTop:       false,
		AssetServer: &assetserver.Options{
			Assets:     frontend.Assets,
			Handler:    nil,
			Middleware: nil,
		},
		// Menu:               app.applicationMenu(),
		Logger:             nil,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		OnStartup:          wailsapp.startup,
		// OnDomReady:         app.domready,
		// OnShutdown:         app.shutdown,
		// OnBeforeClose:      app.beforeClose,
		CSSDragProperty:                  "--wails-draggable",
		CSSDragValue:                     "drag",
		EnableDefaultContextMenu:         false,
		EnableFraudulentWebsiteDetection: false,
		// ZoomFactor:           1.0,
		// IsZoomControlEnabled: false,
		Bind: []interface{}{
			wailsapp,
			optionService,
		},
		ErrorFormatter: func(err error) any { return err.Error() },
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			BackdropType:                      windows.Mica,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			// WebviewUserDataPath:               "",
			// WebviewBrowserPath:                "",
			Theme: app.THEME(),
			// CustomTheme: &windows.ThemeSettings{
			//     DarkModeTitleBar:   windows.RGB(20, 20, 20),
			//     DarkModeTitleText:  windows.RGB(200, 200, 200),
			//     DarkModeBorder:     windows.RGB(20, 0, 20),
			//     LightModeTitleBar:  windows.RGB(200, 200, 200),
			//     LightModeTitleText: windows.RGB(20, 20, 20),
			//     LightModeBorder:    windows.RGB(200, 200, 200),
			// },
			// // User messages that can be customised
			// Messages *windows.Messages
			// // OnSuspend is called when Windows enters low power mode
			// OnSuspend func()
			// // OnResume is called when Windows resumes from low power mode
			// OnResume func(),
			// WebviewGpuDisabled: false,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "My Application",
				Message: "© 2021 Me",
				// Icon:    icon,
			},
		},
		Linux: &linux.Options{
			// Icon: icon,
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         "wails",
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
