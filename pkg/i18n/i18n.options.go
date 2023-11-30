package i18n

import (
	"my-app/pkg/base"
	"path/filepath"
)

type I18nOptions struct {
	base.Options
	APath string // language asset absolute path
}

func DefaultI18nOptions() (*I18nOptions, error) {
	xDir, _, err := base.GetExecutable()
	if err != nil {
		return nil, err
	}

	return &I18nOptions{
		APath: filepath.Join(xDir, "Languages/"),
	}, nil
}

func NewI18nOptions(dst *I18nOptions) (*I18nOptions, error) {
	src, err := DefaultI18nOptions()
	if err != nil {
		return dst, err
	}
	return base.MergeOptions(src, dst)
}
