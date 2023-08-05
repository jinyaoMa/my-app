package tray

import (
	"context"
)

type MenuItem struct {
	*MenuItemBase
}

// CanCheck implements IMenuItem.
func (*MenuItem) CanCheck() bool {
	return false
}

// Checked implements IMenuItem.
func (*MenuItem) Checked() bool {
	return false
}

// Enabled implements IMenuItem.
func (*MenuItem) Enabled() bool {
	return true
}

// Key implements IMenuItem.
func (*MenuItem) Key() string {
	return ""
}

// OnClick implements IMenuItem.
func (*MenuItem) OnClick() (quit bool) {
	return false
}

// Separator implements IMenuItem.
func (*MenuItem) Separator() bool {
	return false
}

// Visible implements IMenuItem.
func (*MenuItem) Visible() bool {
	return true
}

func NewMenuItem(ctx context.Context, menuitems ...IMenuItem) IMenuItem {
	return &MenuItem{
		MenuItemBase: &MenuItemBase{
			Ctx:       ctx,
			MenuItems: menuitems,
		},
	}
}
