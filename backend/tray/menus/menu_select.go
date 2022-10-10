package menus

import (
	"github.com/getlantern/systray"
)

type MenuSelect struct {
	selected string
	Title    *systray.MenuItem
	Options  map[string]*systray.MenuItem
}

type MenuSelectOption struct {
	Title   string
	Tooltip string
}

// arguement 'keys' to indicate menu items and order them in sequence
func NewMenuSelect(title string, keys []string, opts map[string]MenuSelectOption, selected string) *MenuSelect {
	t := systray.AddMenuItem(title, title)
	t.Disable()

	m := &MenuSelect{
		Title:   t,
		Options: make(map[string]*systray.MenuItem),
	}
	for _, key := range keys {
		opt, ok := opts[key]
		if ok {
			m.Options[key] = systray.AddMenuItem(opt.Title, opt.Tooltip)
		}
	}

	m.Select(selected)

	return m
}

func (m *MenuSelect) Select(selected string) (ok bool) {
	for name, option := range m.Options {
		if name == selected {
			option.Check()
			m.selected = selected
			ok = true
		} else {
			option.Uncheck()
		}
	}
	return
}
