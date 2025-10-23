package main

import "github.com/wailsapp/wails/v3/pkg/application"

const (
	EventOpenMainWindow = "open-main-window"
)

func useEvents(a *application.App) {
	a.Event.On(EventOpenMainWindow, func(event *application.CustomEvent) {
		if WindowMain != nil {
			if !WindowMain.IsVisible() {
				WindowMain.Show()
			}
			if !WindowMain.IsFocused() {
				WindowMain.Focus()
			}
		}
	})
}
