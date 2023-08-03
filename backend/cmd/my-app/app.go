package main

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/internal/interfaces"
	"my-app/backend/internal/service"
)

// App struct
type App struct {
	ctx           context.Context
	optionService interfaces.IOptionService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.optionService = service.NewOptionService(app.App().Db())
}
