package i18n

import (
	"embed"
	"encoding/json"
)

//go:embed locales
var locales embed.FS

const (
	En = "en"
	Zh = "zh"
)

type Locale struct {
	Lang struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"lang"`
	AppName         string `json:"appname"`
	OpenWindow      string `json:"open_window"`
	Quit            string `json:"quit"`
	DisplayLanguage struct {
		Label string `json:"label"`
		Title string `json:"title"`
	} `json:"display_language"`
	ColorTheme struct {
		Label  string `json:"label"`
		Title  string `json:"title"`
		System string `json:"system"`
		Light  string `json:"light"`
		Dark   string `json:"dark"`
	} `json:"color_theme"`
	ApiService struct {
		Label    string `json:"label"`
		Enabled  string `json:"enabled"`
		Disabled string `json:"disabled"`
		Swagger  string `json:"swagger"`
		Start    string `json:"start"`
		Stop     string `json:"stop"`
		Dialog   struct {
			Start string `json:"start"`
			Stop  string `json:"stop"`
		} `json:"dialog"`
	} `json:"api_service"`
	QuitDialog struct {
		Message       string `json:"message"`
		DefaultButton string `json:"default_button"`
		CancelButton  string `json:"cancel_button"`
	} `json:"quit_dialog"`
}

func load() (localeMap map[string]Locale, availableLanguages []string) {
	dirLocales := "locales"
	var al []string
	lm := make(map[string]Locale)
	files, _ := locales.ReadDir(dirLocales)
	for _, f := range files {
		if !f.IsDir() { // load only locale JSON file
			t := Locale{}
			data, _ := locales.ReadFile(dirLocales + "/" + f.Name()) // embed use slash as separator
			json.Unmarshal(data, &t)
			lm[t.Lang.Code] = t
			al = append(al, t.Lang.Code)
		}
	}
	return lm, al
}
