package menus

import (
	"github.com/getlantern/systray"
)

type MenuSwitch struct {
	flag         bool
	FalseOptions map[string]*systray.MenuItem
	TrueOptions  map[string]*systray.MenuItem
}

type MenuSwitchOption struct {
	Title   string
	Tooltip string
	Icon    []byte
	Flag    bool // to indicate show/hide when switch flag
}

// arguement 'keys' to indicate menu items and order them in sequence, and
// arguement 'flag' to indicate default state of the switch menu
func NewMenuSwitch(keys []string, opts map[string]MenuSwitchOption, flag bool) *MenuSwitch {
	m := &MenuSwitch{
		FalseOptions: make(map[string]*systray.MenuItem),
		TrueOptions:  make(map[string]*systray.MenuItem),
	}
	for _, key := range keys {
		opt, ok := opts[key]
		if ok {
			if opt.Flag {
				m.TrueOptions[key] = systray.AddMenuItem(opt.Title, opt.Tooltip)
				m.TrueOptions[key].SetTemplateIcon(opt.Icon, opt.Icon)
			} else {
				m.FalseOptions[key] = systray.AddMenuItem(opt.Title, opt.Tooltip)
				m.FalseOptions[key].SetTemplateIcon(opt.Icon, opt.Icon)
			}
		}
	}

	m.SetFlag(flag)

	return m
}

func (m *MenuSwitch) SetFlag(flag bool) {
	m.flag = flag
	if flag {
		for _, option := range m.FalseOptions {
			option.Hide()
		}
		for _, option := range m.TrueOptions {
			option.Show()
		}
	} else {
		for _, option := range m.FalseOptions {
			option.Show()
		}
		for _, option := range m.TrueOptions {
			option.Hide()
		}
	}
}

func (m *MenuSwitch) Refresh() {
	m.SetFlag(m.flag)
}
