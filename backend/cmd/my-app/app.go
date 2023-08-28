package main

import (
	"context"
	"my-app/backend/cmd/my-tray/menuitem"
	"my-app/backend/internal/app"
	"my-app/backend/internal/implements/crud"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/tray"
)

// App struct
type App struct {
	ctx        context.Context
	crudOption interfaces.ICRUDOption
}

// NewApp creates a new App application struct
func NewApp() *App {
	tray.Register(menuitem.Root())
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	menuitem.Root().SetContext(ctx)
	a.crudOption = crud.NewCRUDOption(app.DB())
}
