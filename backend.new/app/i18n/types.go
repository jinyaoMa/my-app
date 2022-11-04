package i18n

const LanguagePlaceholder = Language("")

var TranslationPlaceholder = Translation{
	Lang: TranslationLang{
		Code: LanguagePlaceholder.ToString(),
		Text: LanguagePlaceholder.ToString(),
	},
	AppName:    "[AppName]",
	OpenWindow: "[OpenWindow]",
	Quit:       "[Quit]",
	DisplayLanguage: TranslationDisplayLanguage{
		Label: "[Label]",
		Title: "[Title]",
	},
	ColorTheme: TranslationColorTheme{
		Label:  "[Label]",
		Title:  "[Title]",
		System: "[System]",
		Light:  "[Light]",
		Dark:   "[Dark]",
	},
	WebService: TranslationWebService{
		Label:     "[Label]",
		Enabled:   "[Enabled]",
		Disabled:  "[Disabled]",
		VitePress: "[VitePress]",
		Swagger:   "[Swagger]",
		Start:     "[Start]",
		Stop:      "[Stop]",
	},
	QuitDialog: TranslationQuitDialog{
		Message:       "[Message]",
		DefaultButton: "[DefaultButton]",
		CancelButton:  "[CancelButton]",
	},
}

type Translation struct {
	Lang            TranslationLang            `json:"lang"`
	AppName         string                     `json:"appname"`
	OpenWindow      string                     `json:"open_window"`
	Quit            string                     `json:"quit"`
	DisplayLanguage TranslationDisplayLanguage `json:"display_language"`
	ColorTheme      TranslationColorTheme      `json:"color_theme"`
	WebService      TranslationWebService      `json:"web_service"`
	QuitDialog      TranslationQuitDialog      `json:"quit_dialog"`
}

type TranslationLang struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type TranslationDisplayLanguage struct {
	Label string `json:"label"`
	Title string `json:"title"`
}

type TranslationColorTheme struct {
	Label  string `json:"label"`
	Title  string `json:"title"`
	System string `json:"system"`
	Light  string `json:"light"`
	Dark   string `json:"dark"`
}

type TranslationWebService struct {
	Label     string `json:"label"`
	Enabled   string `json:"enabled"`
	Disabled  string `json:"disabled"`
	VitePress string `json:"vitepress"`
	Swagger   string `json:"swagger"`
	Start     string `json:"start"`
	Stop      string `json:"stop"`
}

type TranslationQuitDialog struct {
	Message       string `json:"message"`
	DefaultButton string `json:"default_button"`
	CancelButton  string `json:"cancel_button"`
}

type Language string

func (l Language) ToString() string {
	return string(l)
}
