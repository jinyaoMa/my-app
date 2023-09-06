package main

import (
	"context"
	"my-app/backend/cmd/my-tray/menuitems"
	"my-app/backend/pkg/tray"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	tray.Register(menuitems.Root())
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	menuitems.Root().SetContext(ctx)
}
