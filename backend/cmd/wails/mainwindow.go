package main

import "github.com/wailsapp/wails/v3/pkg/application"

const (
	MainWindowName = "Main Window"
)

var (
	MainWindow *application.WebviewWindow
)

func newMainWindow(a *application.App) *application.WebviewWindow {
	MainWindow = a.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:                     MainWindowName,
		Title:                    AppName,
		Width:                    1024,
		Height:                   640,
		BackgroundColour:         application.NewRGB(27, 38, 54),
		ContentProtectionEnabled: true,
		URL:                      "/app/",
	})
	return MainWindow
}
