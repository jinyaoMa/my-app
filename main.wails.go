package main

/*
import (
	"context"
	"my-app/backend/app"
	"my-app/backend/app/config"
	"my-app/backend/service"
	"my-app/backend/tray"

	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// Wails Life Cycle
type wailsapp struct {
	ctx context.Context
}

// get title from i18n translation
func (w *wailsapp) title() string {
	return app.App().CurrentTranslation().AppName
}

// get current color theme from app config
func (w *wailsapp) windowTheme() (t windows.Theme) {
	switch app.App().Config().ColorTheme {
	default:
		t = windows.SystemDefault
	case config.ColorThemeLight:
		t = windows.Light
	case config.ColorThemeDark:
		t = windows.Dark
	}
	return
}

// startup is called at application startup
func (w *wailsapp) startup(ctx context.Context) {
	w.ctx = ctx
	s := service.Service(ctx)
	{
		s.InitializeSuperUser()
	}
	tray.Tray(ctx)
	app.App().Log().Wails().Print("WAILS START UP")
}

// domReady is called after the front-end dom has been loaded
func (w *wailsapp) domReady(ctx context.Context) {
	if app.App().Config().Web().IsAutoStart() {
		tray.Tray().StartWebService()
	}
	app.App().Log().Wails().Print("WAILS DOM READY")
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
func (w *wailsapp) beforeClose(ctx context.Context) (prevent bool) {
	app.App().Log().Wails().Print("WAILS BEFORE CLOSE")
	return false
}

// shutdown is called at application termination
func (w *wailsapp) shutdown(ctx context.Context) {
	app.App().Log().Wails().Print("WAILS SHUTDOWN")
}

// suspend is called when Windows enters low power mode
func (w *wailsapp) suspend() {
	app.App().Log().Wails().Print("WAILS SUSPEND")
}

// resume is called when Windows resumes from low power mode
func (w *wailsapp) resume() {
	app.App().Log().Wails().Print("WAILS RESUME")
}
*/
