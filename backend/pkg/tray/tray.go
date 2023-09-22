package tray

import (
	"errors"
	"reflect"
	"sync"

	"github.com/getlantern/systray"
)

var (
	once          sync.Once
	menus         []IMenuItem
	cases         []reflect.SelectCase
	menuItemCache map[string]*systray.MenuItem
)

func init() {
	menuItemCache = make(map[string]*systray.MenuItem)
}

func Register(tray IMenuItemBase) {
	once.Do(func() {
		systray.Register(onReady(tray), nil)
	})
}

func Run(tray IMenuItemBase) {
	systray.Run(onReady(tray), nil)
}

func UpdateTrayIcon(templateIconBytes []byte, regularIconBytes []byte) {
	systray.SetTemplateIcon(templateIconBytes, regularIconBytes)
}

// update systray ui based on state loaded from tray interface
func Update(tray IMenuItemBase, initialized ...bool) error {
	if len(initialized) == 0 {
		initialized = append(initialized, false)
	}

	icon := tray.Icon()
	title := tray.Title()
	tooltip := tray.Tooltip()

	if len(icon) > 0 {
		systray.SetTemplateIcon(icon, icon)
	}
	if title != "" {
		systray.SetTitle(title)
	}
	if tooltip != "" {
		systray.SetTooltip(tooltip)
	}
	for _, item := range tray.Items() {
		if initialized[0] {
			if item.Separator() {
				systray.AddSeparator()
			} else {
				var mi *systray.MenuItem
				if item.CanCheck() {
					mi = systray.AddMenuItemCheckbox("", "", item.Checked())
				} else {
					mi = systray.AddMenuItem("", "")
				}
				if err := update(item, initialized[0], mi); err != nil {
					return err
				}
			}
		} else if !item.Separator() {
			if err := update(item, initialized[0]); err != nil {
				return err
			}
		}
	}

	return nil
}

func update(item IMenuItem, initialized bool, menuItems ...*systray.MenuItem) error {
	key := item.Key()
	icon := item.Icon()
	title := item.Title()
	tooltip := item.Tooltip()

	if len(menuItems) > 0 && initialized {
		menuItemCache[key] = menuItems[0]
	}

	if menuItem, ok := menuItemCache[key]; ok {
		if initialized && item.CanClick() {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(menuItem.ClickedCh),
			})
			menus = append(menus, item)
		}

		if len(icon) > 0 {
			menuItem.SetTemplateIcon(icon, icon)
		}
		if title != "" {
			menuItem.SetTitle(title)
		}
		if tooltip != "" {
			menuItem.SetTooltip(tooltip)
		}
		if item.Visible() {
			menuItem.Show()
			if item.Enabled() {
				menuItem.Enable()
			} else {
				menuItem.Disable()
			}
			if item.Checked() {
				menuItem.Check()
			} else {
				menuItem.Uncheck()
			}
		} else {
			menuItem.Hide()
		}
		for _, item := range item.Items() {
			if initialized {
				var mi *systray.MenuItem
				if item.CanCheck() {
					mi = menuItem.AddSubMenuItemCheckbox("", "", item.Checked())
				} else {
					mi = menuItem.AddSubMenuItem("", "")
				}
				if err := update(item, initialized, mi); err != nil {
					return err
				}
			} else {
				if err := update(item, initialized); err != nil {
					return err
				}
			}
		}
	} else {
		return errors.New("tray menu key changed after initialized")
	}
	return nil
}

func onReady(tray IMenuItemBase) func() {
	return func() {
		Update(tray, true)
		go func() {
			for {
				chosen, _, _ := reflect.Select(cases)
				if menus[chosen].OnClick() {
					break
				}
			}
			systray.Quit()
		}()
	}
}
