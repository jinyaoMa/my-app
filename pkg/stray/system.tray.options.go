package stray

import (
	"my-app/pkg/base"
	"my-app/pkg/i18n"
)

type SystemTrayOptions[TTranslation i18n.ITranslation] struct {
	base.Options
	TemplateIcon func(translation TTranslation) (templateIconBytes []byte, regularIconBytes []byte)
	Title        func(translation TTranslation) string
	Tooltip      func(translation TTranslation) string
	Menu         []*MenuItem[TTranslation]
}

func DefaultSystemTrayOptions[TTranslation i18n.ITranslation]() *SystemTrayOptions[TTranslation] {
	return &SystemTrayOptions[TTranslation]{
		TemplateIcon: func(translation TTranslation) (templateIconBytes []byte, regularIconBytes []byte) {
			return
		},
		Title: func(translation TTranslation) string {
			return ""
		},
		Tooltip: func(translation TTranslation) string {
			return ""
		},
	}
}

func NewSystemTrayOptions[TTranslation i18n.ITranslation](dst *SystemTrayOptions[TTranslation]) (*SystemTrayOptions[TTranslation], error) {
	return base.MergeOptions(DefaultSystemTrayOptions[TTranslation](), dst)
}
