package menus

import (
	"github.com/getlantern/systray"
)

type MenuItem struct {
	*systray.MenuItem
}

func NewItem(title string, tooltip string, icon ...[]byte) *MenuItem {
	m := &MenuItem{
		MenuItem: systray.AddMenuItem(title, tooltip),
	}

	if len(icon) > 0 {
		m.SetTemplateIcon(icon[0], icon[0])
	}

	return m
}
