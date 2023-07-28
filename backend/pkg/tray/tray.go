package tray

import (
	"sync"

	"github.com/getlantern/systray"
)

var (
	once         sync.Once
	menuCacheMap map[string]*MenuCache
)

type MenuCache struct {
	menuItem *systray.MenuItem
	onClick  func()
}

func Run(tray Interface) {
	systray.Run(onReady(tray), onExit(tray))
}

func Update(tray Interface, isInit bool, menuItems ...*systray.MenuItem) {
	var menuCache *MenuCache
	key := tray.Key()
	icon := tray.Icon()
	title := tray.Title()
	tooltip := tray.Tooltip()

	if key != "" {
		if len(menuItems) > 0 {
			menuCache.menuItem = menuItems[0]
			menuCache.onClick = tray.OnClick
		} else {
			var ok bool
			if menuCache, ok = menuCacheMap[key]; !ok {
				return
			}
		}

		if len(icon) > 0 {
			menuCache.menuItem.SetTemplateIcon(icon, icon)
		}
		if title != "" {
			menuCache.menuItem.SetTitle(title)
		}
		if tooltip != "" {
			menuCache.menuItem.SetTooltip(tooltip)
		}
		if tray.Visible() {
			menuCache.menuItem.Show()
		} else {
			menuCache.menuItem.Hide()
		}
		if tray.Enabled() {
			menuCache.menuItem.Enable()
		} else {
			menuCache.menuItem.Disable()
		}
		if tray.Checked() {
			menuCache.menuItem.Check()
		} else {
			menuCache.menuItem.Uncheck()
		}
		for _, item := range tray.Items() {
			if isInit {
				Update(item, isInit, menuCache.menuItem.AddSubMenuItem("", ""))
			} else {
				Update(item, isInit)
			}
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
			if isInit {
				Update(item, isInit, systray.AddMenuItem("", ""))
			} else {
				Update(item, isInit)
			}
		}
	}
}

func onReady(tray Interface) func() {
	return func() {
		Update(tray, true)
	}
}

func onExit(tray Interface) func() {
	return func() {

	}
}
