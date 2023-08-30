package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type quit struct {
	*tray.MenuItem
}

// Key implements tray.IMenuItem.
func (*quit) Key() string {
	return "quit"
}

// CanClick implements tray.IMenuItem.
func (*quit) CanClick() bool {
	return true
}

// OnClick implements tray.IMenuItem.
func (q *quit) OnClick() (quit bool) {
	//runtime.Quit(q.ctx)
	return true
}

// Title implements tray.IMenuItem.
func (*quit) Title() string {
	return app.T().Quit
}

// Tooltip implements tray.IMenuItem.
func (*quit) Tooltip() string {
	return app.T().Quit
}

func newQuit(ctx context.Context) tray.IMenuItem {
	return &quit{
		MenuItem: tray.NewMenuItem(ctx),
	}
}
