package i18n

import (
	"my-app/pkg/base"
	"path/filepath"
)

type I18nOptions[TTranslation ITranslation] struct {
	base.Options
	APath       string       // language asset absolute path
	Placeholder TTranslation // translation placeholder
}

func DefaultI18nOptions[TTranslation ITranslation]() (*I18nOptions[TTranslation], error) {
	xDir, _, err := base.GetExecutable()
	if err != nil {
		return nil, err
	}

	return &I18nOptions[TTranslation]{
		APath:       filepath.Join(xDir, "Languages/"),
		Placeholder: *new(TTranslation),
	}, nil
}

func NewI18nOptions[TTranslation ITranslation](dst *I18nOptions[TTranslation]) (*I18nOptions[TTranslation], error) {
	src, err := DefaultI18nOptions[TTranslation]()
	if err != nil {
		return dst, err
	}
	return base.MergeOptions(src, dst)
}
