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
	localeMap          map[string]Locale
	availableLanguages []string
	currentLanguage    string
}

func I18n() *i18n {
	once.Do(func() {
		localeMap, availableLanguages := load()

		currentLanguage := ""
		if len(availableLanguages) > 0 {
			currentLanguage = availableLanguages[0]
		}

		instance = &i18n{
			localeMap:          localeMap,
			availableLanguages: availableLanguages,
			currentLanguage:    currentLanguage,
		}
	})
	return instance
}

func (i *i18n) Change(lang string) *i18n {
	if slices.Contains(i.availableLanguages, lang) {
		i.currentLanguage = lang
	}
	return i
}

func (i *i18n) Get(lang string) (l Locale) {
	return i.localeMap[lang]
}

func (i *i18n) GetCurrentLanguage() string {
	return i.currentLanguage
}

func (i *i18n) GetLangText(lang string) string {
	return i.localeMap[lang].Lang.Text
}

func (i *i18n) Locale() (l Locale) {
	return i.Get(i.currentLanguage)
}
