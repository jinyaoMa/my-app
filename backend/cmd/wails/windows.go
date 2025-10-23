package main

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

const (
	WindowMainName = "window-main"
)

var (
	WindowMain *application.WebviewWindow
)

func useWindows(a *application.App) {
	WindowMain = a.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:                     WindowMainName,
		Title:                    AppName,
		MinWidth:                 1024,
		MinHeight:                640,
		Width:                    1024,
		Height:                   640,
		BackgroundColour:         application.NewRGB(27, 38, 54),
		ContentProtectionEnabled: true,
		URL:                      "/app/",
	})
	WindowMain.RegisterHook(events.Common.WindowClosing, func(event *application.WindowEvent) {
		event.Cancel()
		WindowMain.Hide()
	})
}
