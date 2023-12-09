package i18n

import (
	"encoding/json"
	"io"
	"io/fs"
	"my-app/pkg/base"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/exp/slices"
	"golang.org/x/text/language"
)

type I18n[TTranslation ITranslation] struct {
	options        *I18nOptions[TTranslation]
	availLocales   []string
	translations   []TTranslation
	translationMap map[string]TTranslation // locale: translation
	currentLocale  string
	mutex          sync.Mutex
}

func (i18n *I18n[TTranslation]) Load() error {
	availLocales := make([]string, 0, 18)
	translations := make([]TTranslation, 0, 18)
	translationMap := make(map[string]TTranslation)
	err := filepath.WalkDir(i18n.options.APath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		locale := base.GetFilenameWithoutExtension(path)
		tag, err := language.Parse(locale)
		if err != nil {
			return nil
		}

		if file, err := os.Open(path); err == nil {
			defer file.Close()
			data, err := io.ReadAll(file)
			if err != nil {
				return nil
			}

			translation := *new(TTranslation)
			err = json.Unmarshal(data, translation)
			if err != nil {
				return nil
			}

			locale = translation.GetLocale()
			if tag.String() != locale {
				return nil
			}

			availLocales = append(availLocales, locale)
			translations = append(translations, translation)
			translationMap[locale] = translation
		}
		return nil
	})
	if err != nil {
		return err
	}

	i18n.mutex.Lock()
	defer i18n.mutex.Unlock()
	i18n.availLocales = availLocales
	i18n.translations = translations
	i18n.translationMap = translationMap
	if i18n.options.Locale == "" {
		if len(i18n.availLocales) > 0 {
			i18n.currentLocale = i18n.availLocales[0]
		}
	} else {
		i18n.currentLocale = i18n.options.Locale
	}
	return nil
}

func (i18n *I18n[TTranslation]) Locale(locales ...string) (currentLocale string) {
	if len(locales) == 1 {
		i18n.mutex.Lock()
		defer i18n.mutex.Unlock()
		if slices.Contains(i18n.availLocales, locales[0]) {
			i18n.currentLocale = locales[0]
		}
	}
	return i18n.currentLocale
}

func (i18n *I18n[TTranslation]) T() (translation TTranslation) {
	if translation, ok := i18n.translationMap[i18n.currentLocale]; ok {
		return translation
	}
	return i18n.options.Placeholder
}

func NewI18n[TTranslation ITranslation](options *I18nOptions[TTranslation]) (i18n *I18n[TTranslation], err error) {
	options, err = NewI18nOptions(options)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(options.APath, os.ModeDir); err != nil {
		return nil, err
	}

	return &I18n[TTranslation]{
		options:        options,
		availLocales:   make([]string, 0, 18),
		translations:   make([]TTranslation, 0, 18),
		translationMap: make(map[string]TTranslation),
	}, nil
}
