package tray

import "context"

type Interface interface {
	Icon() []byte
	Title() string
	Tooltip() string

	// submenu
	Items() []IMenuItem

	SetContext(ctx context.Context)
}

type IMenuItem interface {
	// identify menuitems, used when initialized, error if changed after initialized
	Key() string

	// append separator to systray root menu only
	Separator() bool

	Icon() []byte

	Title() string

	Tooltip() string

	Visible() bool

	Enabled() bool

	CanCheck() bool // for linux to use checkbox menuitem
	Checked() bool

	OnClick() (quit bool)

	// submenu
	Items() []IMenuItem

	SetContext(ctx context.Context)
}
