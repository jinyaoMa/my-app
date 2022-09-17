package menus

import (
	"my-app/backend/pkg/i18n"

	"github.com/getlantern/systray"
)

type DisplayLanguageListener struct {
	OnDisplayLanguageChanged func(lang string) (ok bool, complete func())
}

type DisplayLanguage struct {
	isWatched bool
	chanStop  chan struct{}
	title     *systray.MenuItem
	english   *systray.MenuItem
	chinese   *systray.MenuItem
}

func NewDisplayLanguage() *DisplayLanguage {
	return &DisplayLanguage{
		chanStop: make(chan struct{}, 1),
		title:    systray.AddMenuItem("", ""),
		english:  systray.AddMenuItem("", ""),
		chinese:  systray.AddMenuItem("", ""),
	}
}

func (dl *DisplayLanguage) SetIcon(templateIconBytes []byte, regularIconBytes []byte) *DisplayLanguage {
	dl.title.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return dl
}

func (dl *DisplayLanguage) SetLocale() *DisplayLanguage {
	locale := i18n.I18n().Locale()
	dl.title.SetTitle(locale.DisplayLanguage.Title)
	dl.title.SetTooltip(locale.DisplayLanguage.Title)
	dl.title.Disable()
	dl.english.SetTitle(i18n.I18n().GetLangText(i18n.En))
	dl.english.SetTooltip(i18n.I18n().GetLangText(i18n.En))
	dl.chinese.SetTitle(i18n.I18n().GetLangText(i18n.Zh))
	dl.chinese.SetTooltip(i18n.I18n().GetLangText(i18n.Zh))
	return dl
}

func (dl *DisplayLanguage) Watch(listener DisplayLanguageListener) *DisplayLanguage {
	if dl.isWatched {
		return dl
	}

	dl.isWatched = true
	go func() {
		for {
			select {
			case <-dl.english.ClickedCh:
				if ok, complete := listener.OnDisplayLanguageChanged(i18n.En); ok {
					dl.chinese.Uncheck()
					dl.english.Check()
					complete()
				}
			case <-dl.chinese.ClickedCh:
				if ok, complete := listener.OnDisplayLanguageChanged(i18n.Zh); ok {
					dl.chinese.Check()
					dl.english.Uncheck()
					complete()
				}
			case <-dl.chanStop:
				return
			}
		}
	}()
	return dl
}

func (dl *DisplayLanguage) StopWatch() *DisplayLanguage {
	if dl.isWatched {
		dl.chanStop <- struct{}{}
	}
	return dl
}

func (dl *DisplayLanguage) ClickEnglish() *DisplayLanguage {
	dl.english.ClickedCh <- struct{}{}
	return dl
}

func (dl *DisplayLanguage) ClickChinese() *DisplayLanguage {
	dl.chinese.ClickedCh <- struct{}{}
	return dl
}
