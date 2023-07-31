package app

import "my-app/backend/pkg/assetsio"

type Translation struct {
	Lang       *assetsio.Lang `json:"lang"`
	AppName    string         `json:"app_name"`
	OpenWindow string         `json:"open_window"`
	Quit       string         `json:"quit"`
}

// Metadata implements assetsio.ITranslation.
func (t *Translation) Metadata() *assetsio.Lang {
	return t.Lang
}

func DefaultTranslation() *Translation {
	return &Translation{
		Lang: &assetsio.Lang{
			Code: "[LangCode]",
			Text: "[LangText]",
		},
		AppName:    "[AppName]",
		OpenWindow: "[OpenWindow]",
		Quit:       "[Quit]",
	}
}

func NewTranslation() assetsio.ITranslation {
	return &Translation{}
}
