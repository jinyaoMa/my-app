package tray

import "github.com/getlantern/systray"

type Interface interface {
	// identify menuitems, used when initialized, error if changed after initialized
	Key() string

	// append separator to systray root menu only
	Separator() bool

	Icon() []byte

	Title() string

	Tooltip() string

	Visible(visible ...bool) bool

	Enabled(enable ...bool) bool

	Checked(checked ...bool) bool

	OnClick(self Interface, menuItem *systray.MenuItem) (quit bool)

	// submenu
	Items() []Interface
}
