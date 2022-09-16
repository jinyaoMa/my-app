package menus

import (
	"my-app/backend/app"
	"my-app/backend/pkg/i18n"

	"github.com/getlantern/systray"
)

type ColorThemeListener struct {
	OnColorThemeChanged func(theme string) (ok bool)
}

type ColorTheme struct {
	isWatched bool
	chanStop  chan struct{}
	title     *systray.MenuItem
	system    *systray.MenuItem
	light     *systray.MenuItem
	dark      *systray.MenuItem
}

func NewColorTheme() *ColorTheme {
	return &ColorTheme{
		chanStop: make(chan struct{}, 1),
		title:    systray.AddMenuItem("", ""),
		system:   systray.AddMenuItem("", ""),
		light:    systray.AddMenuItem("", ""),
		dark:     systray.AddMenuItem("", ""),
	}
}

func (ct *ColorTheme) SetIcon(templateIconBytes []byte, regularIconBytes []byte) *ColorTheme {
	ct.title.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return ct
}

func (ct *ColorTheme) SetLocale(locale i18n.Locale) *ColorTheme {
	ct.title.SetTitle(locale.ColorTheme.Title)
	ct.title.SetTooltip(locale.ColorTheme.Title)
	ct.title.Disable()
	ct.system.SetTitle(locale.ColorTheme.System)
	ct.system.SetTooltip(locale.ColorTheme.System)
	ct.light.SetTitle(locale.ColorTheme.Light)
	ct.light.SetTooltip(locale.ColorTheme.Light)
	ct.dark.SetTitle(locale.ColorTheme.Dark)
	ct.dark.SetTooltip(locale.ColorTheme.Dark)
	return ct
}

func (ct *ColorTheme) Watch(listener ColorThemeListener) *ColorTheme {
	if ct.isWatched {
		return ct
	}

	ct.isWatched = true
	go func() {
		for {
			select {
			case <-ct.system.ClickedCh:
				if listener.OnColorThemeChanged(app.ColorThemeSystem) {
					ct.system.Check()
					ct.light.Uncheck()
					ct.dark.Uncheck()
				}
			case <-ct.light.ClickedCh:
				if listener.OnColorThemeChanged(app.ColorThemeLight) {
					ct.system.Uncheck()
					ct.light.Check()
					ct.dark.Uncheck()
				}
			case <-ct.dark.ClickedCh:
				if listener.OnColorThemeChanged(app.ColorThemeDark) {
					ct.system.Uncheck()
					ct.light.Uncheck()
					ct.dark.Check()
				}
			case <-ct.chanStop:
				return
			}
		}
	}()
	return ct
}

func (ct *ColorTheme) StopWatch() *ColorTheme {
	if ct.isWatched {
		ct.chanStop <- struct{}{}
	}
	return ct
}

func (ct *ColorTheme) ClickSystem() *ColorTheme {
	ct.system.ClickedCh <- struct{}{}
	return ct
}

func (ct *ColorTheme) ClickLight() *ColorTheme {
	ct.light.ClickedCh <- struct{}{}
	return ct
}

func (ct *ColorTheme) ClickDark() *ColorTheme {
	ct.dark.ClickedCh <- struct{}{}
	return ct
}
