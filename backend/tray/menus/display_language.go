package menus

import (
	"my-app/backend/pkg/i18n"

	"github.com/getlantern/systray"
)

type DisplayLanguageListener struct {
	OnDisplayLanguageChanged func(lang string) (ok bool)
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

func (dl *DisplayLanguage) SetLocale(locale i18n.Locale) *DisplayLanguage {
	dl.title.SetTitle(locale.DisplayLanguage)
	dl.title.SetTooltip(locale.DisplayLanguage)
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
				if listener.OnDisplayLanguageChanged(i18n.En) {
					dl.chinese.Uncheck()
					dl.english.Check()
				}
			case <-dl.chinese.ClickedCh:
				if listener.OnDisplayLanguageChanged(i18n.Zh) {
					dl.chinese.Check()
					dl.english.Uncheck()
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
