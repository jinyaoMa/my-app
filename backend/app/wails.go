package app

import (
	"context"
	"log"
	"my-app/backend/i18n"
	"my-app/backend/tray"
	"my-app/backend/tray/menus"
)

// Wailsapp Life Cycle
type WailsLifeCycle struct {
	ctx context.Context
}

// Initialize is called before application startup
// **Need to manually run this before running Wails**
func (wlc *WailsLifeCycle) Initialize() *WailsLifeCycle {
	tray.Tray() // systray must be set up before starting wails life cycle...
	log.Println("WAILS INITIALIZE")
	return wlc
}

// Startup is called at application startup
// **Need to manually assign to Wails Options**
func (wlc *WailsLifeCycle) Startup(ctx context.Context) {
	wlc.ctx = ctx
	tray.Tray().SetWailsContext(ctx)
	log.Println("WAILS START UP")
}

// DomReady is called after the front-end dom has been loaded
// **Need to manually assign to Wails Options**
func (wlc *WailsLifeCycle) DomReady(ctx context.Context) {
	tray.Tray().
		ChangeTheme(menus.ColorThemeSystem).
		ChangeLanguage(i18n.I18n().GetCurrentLanguage())
	log.Println("WAILS DOM READY")
}

// BeforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// **Need to manually assign to Wails Options**
func (wlc *WailsLifeCycle) BeforeClose(ctx context.Context) (prevent bool) {
	log.Println("WAILS BEFORE CLOSE")
	return false
}

// Shutdown is called at application termination
// **Need to manually assign to Wails Options**
func (wlc *WailsLifeCycle) Shutdown(ctx context.Context) {
	log.Println("WAILS SHUTDOWN")
}

// Suspend is called when Windows enters low power mode
// **Need to manually assign to Wails Options -> Windows Specific Options**
func (wlc *WailsLifeCycle) Suspend() {
	log.Println("WAILS SUSPEND")
}

// Resume is called when Windows resumes from low power mode
// **Need to manually assign to Wails Options -> Windows Specific Options**
func (wlc *WailsLifeCycle) Resume() {
	log.Println("WAILS RESUME")
}
