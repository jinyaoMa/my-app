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
	if app.Web().IsStopping() {
		return app.Assets().GetBytes("tray.orange.ico")
	}
	if app.Web().IsRunning() {
		return app.Assets().GetBytes("tray.green.ico")
	}
	return app.Assets().GetBytes("tray.blue.ico")
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
	return app.T().AppName
}

// Tooltip implements tray.Interface.
func (*root) Tooltip() string {
	return app.T().AppName
}

func newRoot(ctx context.Context) tray.Interface {
	return &root{
		ctx: ctx,
	}
}
