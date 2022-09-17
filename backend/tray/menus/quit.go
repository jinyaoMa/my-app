package menus

import (
	"my-app/backend/pkg/i18n"

	"github.com/getlantern/systray"
)

type QuitListener struct {
	OnQuit func()
}

type Quit struct {
	isWatched bool
	chanStop  chan struct{}
	item      *systray.MenuItem
}

func NewQuit() *Quit {
	return &Quit{
		chanStop: make(chan struct{}, 1),
		item:     systray.AddMenuItem("", ""),
	}
}

func (q *Quit) SetIcon(templateIconBytes []byte, regularIconBytes []byte) *Quit {
	q.item.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return q
}

func (q *Quit) SetLocale() *Quit {
	locale := i18n.I18n().Locale()
	q.item.SetTitle(locale.Quit)
	q.item.SetTooltip(locale.Quit)
	return q
}

func (q *Quit) Watch(listener QuitListener) *Quit {
	if q.isWatched {
		return q
	}

	q.isWatched = true
	go func() {
		for {
			select {
			case <-q.item.ClickedCh:
				listener.OnQuit()
			case <-q.chanStop:
				return
			}
		}
	}()
	return q
}

func (q *Quit) StopWatch() *Quit {
	if q.isWatched {
		q.chanStop <- struct{}{}
	}
	return q
}

func (q *Quit) Click() *Quit {
	q.item.ClickedCh <- struct{}{}
	return q
}
