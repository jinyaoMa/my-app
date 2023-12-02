package i18n

import (
	"encoding/json"
	"io"
	"io/fs"
	"my-app/pkg/base"
	"os"
	"path/filepath"

	"golang.org/x/text/language"
)

type I18n[TTranslation ITranslation] struct {
	options        *I18nOptions
	availLocales   []string
	translations   []TTranslation
	translationMap map[string]TTranslation // locale: translation
}

func (i18n *I18n[TTranslation]) Load() error {
	return filepath.WalkDir(i18n.options.APath, func(path string, d fs.DirEntry, err error) error {
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

			i18n.availLocales = append(i18n.availLocales, locale)
			i18n.translations = append(i18n.translations, translation)
			i18n.translationMap[locale] = translation
		}
		return nil
	})
}

func NewI18n[TTranslation ITranslation](options *I18nOptions) (i18n *I18n[TTranslation], err error) {
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
