package stray

import (
	"reflect"

	"github.com/getlantern/systray"
)

func (tray *systemTray[TTranslation]) onReady() {
	tray.update()
	go tray.routine()
}

func (tray *systemTray[TTranslation]) routine() {
	menuitems := tray.flat(tray.options.Menu...)
	count := len(menuitems)
	cases := make([]reflect.SelectCase, count)
	for i := 0; i < count; i++ {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(menuitems[i].bind.ClickedCh),
		}
	}
	for {
		chosen, _, _ := reflect.Select(cases)
		if menuitems[chosen].OnClick != nil && menuitems[chosen].OnClick(tray.ctx) {
			break
		}
	}
	systray.Quit()
}

func (tray *systemTray[TTranslation]) flat(items ...*MenuItem[TTranslation]) (menuitems []*MenuItem[TTranslation]) {
	count := len(items)
	for i := 0; i < count; i++ {
		if !items[i].Seperator {
			menuitems = append(menuitems, items[i])
			if len(items[i].SubMenu) > 0 {
				menuitems = append(menuitems, tray.flat(items[i].SubMenu...)...)
			}
		}
	}
	return
}

func (tray *systemTray[TTranslation]) update(items ...*MenuItem[TTranslation]) {
	var templateIconBytes, regularIconBytes []byte
	var title, tooltip string
	var menu []*MenuItem[TTranslation]

	if len(items) > 0 {
		templateIconBytes, regularIconBytes = items[0].TemplateIcon(tray.translation)
		title = items[0].Title(tray.translation)
		tooltip = items[0].Tooltip(tray.translation)
		menu = items[0].SubMenu

		menuitem := items[0].bind
		menuitem.SetTemplateIcon(templateIconBytes, regularIconBytes)
		menuitem.SetTitle(title)
		menuitem.SetTooltip(tooltip)

		if items[0].Visible {
			menuitem.Show()
			if items[0].Enabled {
				menuitem.Enable()
			} else {
				menuitem.Disable()
			}
			if items[0].Checked {
				menuitem.Check()
			} else {
				menuitem.Uncheck()
			}
		} else {
			menuitem.Hide()
		}
	} else {
		templateIconBytes, regularIconBytes = tray.options.TemplateIcon(tray.translation)
		title = tray.options.Title(tray.translation)
		tooltip = tray.options.Tooltip(tray.translation)
		menu = tray.options.Menu

		systray.SetTemplateIcon(templateIconBytes, regularIconBytes)
		systray.SetTitle(title)
		systray.SetTooltip(tooltip)
	}

	for _, item := range menu {
		if item.bind == nil {
			if item.Seperator {
				if len(items) == 0 {
					systray.AddSeparator()
				}
			} else {
				if item.CanCheck {
					item.bind = systray.AddMenuItemCheckbox("", "", item.Checked)
				} else {
					item.bind = systray.AddMenuItem("", "")
				}
				tray.update(item)
			}
		} else if !item.Seperator {
			tray.update(item)
		}
	}
	return
}
