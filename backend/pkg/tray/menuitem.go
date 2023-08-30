package tray

import (
	"context"
)

type MenuItem struct {
	*MenuItemBase
}

// CanClick implements IMenuItem.
func (*MenuItem) CanClick() bool {
	return false
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

func NewMenuItem(ctx context.Context, menuitems ...IMenuItem) *MenuItem {
	return &MenuItem{
		MenuItemBase: &MenuItemBase{
			Ctx:       ctx,
			MenuItems: menuitems,
		},
	}
}

func NewIMenuItem(ctx context.Context, menuitems ...IMenuItem) IMenuItem {
	return NewMenuItem(ctx, menuitems...)
}

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

func NewMenuItemBase(ctx context.Context, menuitems ...IMenuItem) *MenuItemBase {
	return &MenuItemBase{
		Ctx:       ctx,
		MenuItems: menuitems,
	}
}

func NewIMenuItemBase(ctx context.Context, menuitems ...IMenuItem) IMenuItemBase {
	return NewMenuItemBase(ctx, menuitems...)
}
