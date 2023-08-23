package tray

import "context"

type IMenuItemBase interface {
	Icon() []byte
	Title() string
	Tooltip() string

	// submenu
	Items() []IMenuItem

	SetContext(ctx context.Context)
}

type IMenuItem interface {
	IMenuItemBase

	// identify menuitems, used when initialized, error if changed after initialized
	// if key is empty, do not listen on clicks
	Key() string

	// append separator to systray root menu only
	Separator() bool

	Visible() bool

	Enabled() bool

	CanCheck() bool // for linux to use checkbox menuitem
	Checked() bool

	OnClick() (quit bool)
}
