package app

import "my-app/backend/pkg/aio"

type Translation struct {
	Lang            aio.Lang                   `json:"lang"`
	AppName         string                     `json:"app_name"`
	OpenWindow      string                     `json:"open_window"`
	DisplayLanguage TranslationDisplayLanguage `json:"display_language"`
	ColorTheme      TranslationColorTheme      `json:"color_theme"`
	Quit            string                     `json:"quit"`
}

type TranslationColorTheme struct {
	Label  string `json:"label"`
	Title  string `json:"title"`
	System string `json:"system"`
	Light  string `json:"light"`
	Dark   string `json:"dark"`
}

type TranslationDisplayLanguage struct {
	Label string `json:"label"`
	Title string `json:"title"`
}

// Metadata implements aio.ITranslation.
func (t *Translation) Metadata() aio.Lang {
	return t.Lang
}

func DefaultTranslation() *Translation {
	return &Translation{
		Lang: aio.Lang{
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

func NewTranslation() *Translation {
	return &Translation{}
}

func NewITranslation() aio.ITranslation {
	return NewTranslation()
}
