package menuitem

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type quit struct {
	ctx context.Context
}

// SetContext implements tray.IMenuItem.
func (q *quit) SetContext(ctx context.Context) {
	q.ctx = ctx
}

// CanCheck implements tray.IMenuItem.
func (*quit) CanCheck() bool {
	return false
}

// Checked implements tray.IMenuItem.
func (*quit) Checked() bool {
	return false
}

// Enabled implements tray.IMenuItem.
func (*quit) Enabled() bool {
	return true
}

// Icon implements tray.IMenuItem.
func (*quit) Icon() []byte {
	return nil
}

// Items implements tray.IMenuItem.
func (*quit) Items() []tray.IMenuItem {
	return nil
}

// Key implements tray.IMenuItem.
func (*quit) Key() string {
	return "quit"
}

// OnClick implements tray.IMenuItem.
func (q *quit) OnClick() (quit bool) {
	//runtime.Quit(q.ctx)
	return true
}

// Separator implements tray.IMenuItem.
func (*quit) Separator() bool {
	return false
}

// Title implements tray.IMenuItem.
func (*quit) Title() string {
	return app.T().Quit
}

// Tooltip implements tray.IMenuItem.
func (*quit) Tooltip() string {
	return app.T().Quit
}

// Visible implements tray.IMenuItem.
func (*quit) Visible() bool {
	return true
}

func newQuit(ctx context.Context) tray.IMenuItem {
	return &quit{
		ctx: ctx,
	}
}
