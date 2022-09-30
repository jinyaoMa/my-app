package main

import (
	"context"
	"my-app/backend/app"
	"my-app/backend/tray"
)

// Wails Life Cycle
type WailsLifeCycle struct {
	ctx context.Context
}

func DefaultWailsLifeCycle() *WailsLifeCycle {
	tray.Tray() // systray must be set up before starting wails life cycle...
	return &WailsLifeCycle{}
}

// startup is called at application startup
func (wlc *WailsLifeCycle) startup(ctx context.Context) {
	wlc.ctx = ctx
	app.App().SetWailsContext(ctx).WailsLog().Print("WAILS START UP")
}

// domReady is called after the front-end dom has been loaded
func (wlc *WailsLifeCycle) domReady(ctx context.Context) {
	cfg := app.App().Config()
	t := tray.Tray().
		ChangeColorTheme(cfg.ColorTheme).
		ChangeLanguage(cfg.DisplayLanguage)
	if cfg.Web.AutoStart == "true" {
		t.StartWebService()
	}

	app.App().WailsLog().Print("WAILS DOM READY")
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
func (wlc *WailsLifeCycle) beforeClose(ctx context.Context) (prevent bool) {
	app.App().WailsLog().Print("WAILS BEFORE CLOSE")
	return false
}

// shutdown is called at application termination
func (wlc *WailsLifeCycle) shutdown(ctx context.Context) {
	app.App().WailsLog().Print("WAILS SHUTDOWN")
}

// suspend is called when Windows enters low power mode
func (wlc *WailsLifeCycle) suspend() {
	app.App().WailsLog().Print("WAILS SUSPEND")
}

// resume is called when Windows resumes from low power mode
func (wlc *WailsLifeCycle) resume() {
	app.App().WailsLog().Print("WAILS RESUME")
}
