package menuitem

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/tray"
)

type root struct {
	ctx context.Context
}

// SetContext implements tray.Interface.
func (r *root) SetContext(ctx context.Context) {
	r.ctx = ctx
}

// Icon implements tray.Interface.
func (*root) Icon() []byte {
	if app.App().Web().IsStopping() {
		return app.App().Assets().GetBytes("tray.orange.ico")
	}
	if app.App().Web().IsRunning() {
		return app.App().Assets().GetBytes("tray.green.ico")
	}
	return app.App().Assets().GetBytes("tray.blue.ico")
}

// Items implements tray.Interface.
func (r *root) Items() []tray.IMenuItem {
	return []tray.IMenuItem{
		newOpenWindow(r.ctx),
		newSeparator(r.ctx),
		newQuit(r.ctx),
	}
}

// Title implements tray.Interface.
func (*root) Title() string {
	return app.App().T().AppName
}

// Tooltip implements tray.Interface.
func (*root) Tooltip() string {
	return app.App().T().AppName
}

func newRoot(ctx context.Context) tray.Interface {
	return &root{
		ctx: ctx,
	}
}
