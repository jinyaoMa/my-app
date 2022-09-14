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

func (wlc *WailsLifeCycle) Startup(ctx context.Context) {
	wlc.ctx = ctx
	tray.Tray().
		SetWailsContext(ctx).
		ChangeTheme(menus.ColorThemeSystem).
		ChangeLanguage(i18n.I18n().GetCurrentLanguage())
	log.Println("WAILS START UP")
}

func (wlc *WailsLifeCycle) DomReady(ctx context.Context) {
	log.Println("WAILS DOM READY")
}

func (wlc *WailsLifeCycle) BeforeClose(ctx context.Context) (prevent bool) {
	log.Println("WAILS BEFORE CLOSE")
	return false
}

func (wlc *WailsLifeCycle) Shutdown(ctx context.Context) {
	log.Println("WAILS SHUTDOWN")
}

func (wlc *WailsLifeCycle) Suspend() {
	log.Println("WAILS SUSPEND")
}

func (wlc *WailsLifeCycle) Resume() {
	log.Println("WAILS RESUME")
}
