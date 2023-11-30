package i18n

import "os"

type I18n[TTranslation ITranslation] struct {
	options *I18nOptions
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
