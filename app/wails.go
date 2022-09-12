package app

import (
	"context"
	"log"
)

// Wailsapp Life Cycle
type Wails struct {
	ctx context.Context
}

func (w *Wails) Startup(ctx context.Context) {
	w.ctx = ctx

	log.Println("WAILS START UP")
}

func (w *Wails) DomReady(ctx context.Context) {
	log.Println("WAILS DOM READY")
}

func (w *Wails) BeforeClose(ctx context.Context) (prevent bool) {
	log.Println("WAILS BEFORE CLOSE")
	return false
}

func (w *Wails) Shutdown(ctx context.Context) {
	log.Println("WAILS SHUTDOWN")
}

func (w *Wails) Suspend() {
	log.Println("WAILS SUSPEND")
}

func (w *Wails) Resume() {
	log.Println("WAILS RESUME")
}
