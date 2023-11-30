package i18n

import (
	"io/fs"
	"os"
	"path/filepath"
)

type I18n[TTranslation ITranslation] struct {
	options      *I18nOptions
	availLocales []string
	translations []TTranslation
}

func (i18n *I18n[TTranslation]) Load() error {
	return filepath.WalkDir(i18n.options.APath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
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
		options: options,
	}, nil
}
