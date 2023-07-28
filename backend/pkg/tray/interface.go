package tray

type Interface interface {
	// identify menuitems, used when initialized, error if changed after initialized
	Key() string

	// append separator to systray root menu only
	Separator() bool

	Icon() []byte

	Title() string

	Tooltip() string

	Visible() bool

	Enabled() bool

	Checked() bool

	OnClick() (quit bool)

	// submenu
	Items() []Interface
}
