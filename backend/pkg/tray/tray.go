package tray

import (
	"errors"
	"reflect"
	"sync"

	"github.com/getlantern/systray"
)

var (
	once          sync.Once
	menus         []Interface
	cases         []reflect.SelectCase
	menuItemCache map[string]*systray.MenuItem
)

func Register(tray Interface) {
	once.Do(func() {
		systray.Register(onReady(tray), nil)
	})
}

func Run(tray Interface) {
	systray.Run(onReady(tray), nil)
}

// update systray ui based on state loaded from tray interface
func Update(tray Interface, initialized bool, menuItems ...*systray.MenuItem) error {
	key := tray.Key()
	icon := tray.Icon()
	title := tray.Title()
	tooltip := tray.Tooltip()

	if key != "" {
		if len(menuItems) > 0 && initialized {
			menuItemCache[key] = menuItems[0]
		}

		if menuItem, ok := menuItemCache[key]; ok {
			if initialized {
				cases = append(cases, reflect.SelectCase{
					Dir:  reflect.SelectRecv,
					Chan: reflect.ValueOf(menuItem.ClickedCh),
				})
				menus = append(menus, tray)
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
			if tray.Visible() {
				menuItem.Show()
			} else {
				menuItem.Hide()
			}
			if tray.Enabled() {
				menuItem.Enable()
			} else {
				menuItem.Disable()
			}
			if tray.Checked() {
				menuItem.Check()
			} else {
				menuItem.Uncheck()
			}
			for _, item := range tray.Items() {
				if initialized {
					var mi *systray.MenuItem
					if tray.Checked() {
						mi = menuItem.AddSubMenuItemCheckbox("", "", tray.Checked())
					} else {
						mi = menuItem.AddSubMenuItem("", "")
					}
					Update(item, initialized, mi)
				} else {
					Update(item, initialized)
				}
			}
		} else {
			return errors.New("tray menu key changed after initialized")
		}
	} else {
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
			if initialized {
				if item.Separator() {
					systray.AddSeparator()
				} else {
					var mi *systray.MenuItem
					if tray.Checked() {
						mi = systray.AddMenuItemCheckbox("", "", tray.Checked())
					} else {
						mi = systray.AddMenuItem("", "")
					}
					Update(item, initialized, mi)
				}
			} else {
				Update(item, initialized)
			}
		}
	}
	return nil
}

func Quit() {
	systray.Quit()
}

func onReady(tray Interface) func() {
	return func() {
		Update(tray, true)
		go func() {
			for {
				chosen, _, _ := reflect.Select(cases)
				if menus[chosen].OnClick() {
					break
				}
			}
		}()
	}
}
