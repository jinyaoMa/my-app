package app

import "my-app/backend/pkg/assetsio"

type Translation struct {
	Lang       assetsio.Lang         `json:"lang"`
	AppName    string                `json:"app_name"`
	OpenWindow string                `json:"open_window"`
	ColorTheme TranslationColorTheme `json:"color_theme"`
	Quit       string                `json:"quit"`
}

type TranslationColorTheme struct {
	Label  string `json:"label"`
	Title  string `json:"title"`
	System string `json:"system"`
	Light  string `json:"light"`
	Dark   string `json:"dark"`
}

// Metadata implements assetsio.ITranslation.
func (t *Translation) Metadata() assetsio.Lang {
	return t.Lang
}

func DefaultTranslation() *Translation {
	return &Translation{
		Lang: assetsio.Lang{
			Code: "[LangCode]",
			Text: "[LangText]",
		},
		AppName:    "[AppName]",
		OpenWindow: "[OpenWindow]",
		ColorTheme: TranslationColorTheme{
			Label:  "[ColorThemeLabel]",
			Title:  "[ColorThemeTitle]",
			System: "[ColorThemeSystem]",
			Light:  "[ColorThemeLight]",
			Dark:   "[ColorThemeDark]",
		},
		Quit: "[Quit]",
	}
}

func NewTranslation() assetsio.ITranslation {
	return &Translation{}
}
