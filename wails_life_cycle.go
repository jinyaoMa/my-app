package main

import (
	"context"
	"log"
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
	tray.Tray().SetWailsContext(ctx)
	log.Println("WAILS START UP")
}

// domReady is called after the front-end dom has been loaded
func (wlc *WailsLifeCycle) domReady(ctx context.Context) {
	cfg := app.App().Config()
	tray.Tray().
		ChangeTheme(cfg.Theme).
		ChangeLanguage(cfg.Language)
	log.Println("WAILS DOM READY")
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
func (wlc *WailsLifeCycle) beforeClose(ctx context.Context) (prevent bool) {
	log.Println("WAILS BEFORE CLOSE")
	return false
}

// shutdown is called at application termination
func (wlc *WailsLifeCycle) shutdown(ctx context.Context) {
	log.Println("WAILS SHUTDOWN")
}

// suspend is called when Windows enters low power mode
func (wlc *WailsLifeCycle) suspend() {
	log.Println("WAILS SUSPEND")
}

// resume is called when Windows resumes from low power mode
func (wlc *WailsLifeCycle) resume() {
	log.Println("WAILS RESUME")
}
