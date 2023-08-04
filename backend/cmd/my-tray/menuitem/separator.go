package menuitem

import (
	"context"
	"my-app/backend/pkg/tray"
)

type separator struct {
	tray.IMenuItem
	ctx context.Context
}

// Separator implements tray.IMenuItem.
func (*separator) Separator() bool {
	return true
}

// SetContext implements tray.Interface.
func (s *separator) SetContext(ctx context.Context) {
	s.ctx = ctx
}

func newSeparator(ctx context.Context) tray.IMenuItem {
	return &separator{
		ctx: ctx,
	}
}
