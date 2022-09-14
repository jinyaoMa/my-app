package menus

import (
	"my-app/backend/i18n"

	"github.com/getlantern/systray"
)

type OpenWindowListener struct {
	OnOpenWindow func()
}

type OpenWindow struct {
	isWatched bool
	chanStop  chan struct{}
	item      *systray.MenuItem
}

func NewOpenWindow() *OpenWindow {
	return &OpenWindow{
		chanStop: make(chan struct{}, 1),
		item:     systray.AddMenuItem("", ""),
	}
}

func (ow *OpenWindow) SetIcon(templateIconBytes []byte, regularIconBytes []byte) *OpenWindow {
	ow.item.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return ow
}

func (ow *OpenWindow) SetLocale(locale i18n.Locale) *OpenWindow {
	ow.item.SetTitle(locale.OpenWindow)
	ow.item.SetTooltip(locale.OpenWindow)
	return ow
}

func (ow *OpenWindow) Watch(listener OpenWindowListener) *OpenWindow {
	if ow.isWatched {
		return ow
	}

	ow.isWatched = true
	go func() {
		for {
			select {
			case <-ow.item.ClickedCh:
				listener.OnOpenWindow()
			case <-ow.chanStop:
				return
			}
		}
	}()
	return ow
}

func (ow *OpenWindow) StopWatch() *OpenWindow {
	if ow.isWatched {
		ow.chanStop <- struct{}{}
	}
	return ow
}

func (ow *OpenWindow) Click() *OpenWindow {
	ow.item.ClickedCh <- struct{}{}
	return ow
}
