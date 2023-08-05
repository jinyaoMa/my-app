package tray

import (
	"context"
)

type MenuItemBase struct {
	Ctx       context.Context
	MenuItems []IMenuItem
}

// Icon implements Interface.
func (*MenuItemBase) Icon() []byte {
	return nil
}

// Items implements Interface.
func (b *MenuItemBase) Items() []IMenuItem {
	return b.MenuItems
}

// SetContext implements Interface.
func (b *MenuItemBase) SetContext(ctx context.Context) {
	b.Ctx = ctx
	for _, item := range b.MenuItems {
		item.SetContext(ctx)
	}
}

// Title implements Interface.
func (*MenuItemBase) Title() string {
	return "[Title]"
}

// Tooltip implements Interface.
func (*MenuItemBase) Tooltip() string {
	return "[Tooltip]"
}

func NewMenuItemBase(ctx context.Context, menuitems ...IMenuItem) IMenuItemBase {
	return &MenuItemBase{
		Ctx:       ctx,
		MenuItems: menuitems,
	}
}
