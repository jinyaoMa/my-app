package i18n

import (
	"sync"

	"golang.org/x/exp/slices"
)

var (
	once     sync.Once
	instance *i18n
)

type i18n struct {
	localeMap          localeMap
	availableLanguages []string
	currentLocale      string
}

func I18n() *i18n {
	once.Do(func() {
		localeMap, availableLanguages := load()

		currentLocale := ""
		if len(availableLanguages) > 0 {
			currentLocale = availableLanguages[0]
		}

		instance = &i18n{
			localeMap:          localeMap,
			availableLanguages: availableLanguages,
			currentLocale:      currentLocale,
		}
	})
	return instance
}

func (i *i18n) Change(lang string) (ok bool) {
	if slices.Contains(i.availableLanguages, lang) {
		i.currentLocale = lang
		return true
	}
	return false
}

func (i *i18n) Locale() (l locale) {
	return i.localeMap[i.currentLocale]
}
