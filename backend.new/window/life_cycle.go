package window

import (
	"context"
	"my-app/backend.new/app"
	"my-app/backend.new/model"
	"my-app/backend/tray"
)

// startup is called at application startup
func (w *window) startup(ctx context.Context) {
	app.App().SetContext(ctx).
		Log().Wails().Print("WAILS START UP")
}

// domReady is called after the front-end dom has been loaded
func (w *window) domReady(ctx context.Context) {
	app.App().UseCfg(func(cfg *app.Config) {
		if cfg.Get(model.OptionWebAutoStart) == "true" {
			tray.Tray().StartWebService()
		}
	}).Log().Wails().Print("WAILS DOM READY")
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
func (w *window) beforeClose(ctx context.Context) (prevent bool) {
	app.App().Log().Wails().Print("WAILS BEFORE CLOSE")
	return false
}

// shutdown is called at application termination
func (w *window) shutdown(ctx context.Context) {
	app.App().Log().Wails().Print("WAILS SHUTDOWN")
}

// suspend is called when Windows enters low power mode
func (w *window) suspend() {
	app.App().Log().Wails().Print("WAILS SUSPEND")
}

// resume is called when Windows resumes from low power mode
func (w *window) resume() {
	app.App().Log().Wails().Print("WAILS RESUME")
}
