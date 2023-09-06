package main

import (
	"my-app/backend/internal/app"
	"my-app/frontend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func main() {
	// Create an instance of the app structure
	wailsapp := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "customlayout",
		Width:             1024,
		Height:            600,
		Assets:            frontend.Assets,
		HideWindowOnClose: true,
		BackgroundColour:  &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:         wailsapp.startup,
		Bind: []interface{}{
			wailsapp,
		},
		Windows: &windows.Options{
			Theme: app.THEME(),
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
