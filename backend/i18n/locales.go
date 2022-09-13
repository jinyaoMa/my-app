package i18n

import (
	"embed"
	"encoding/json"
	"path/filepath"
)

//go:embed locales
var locales embed.FS

const (
	dirLocales = "locales"
)

type locale struct {
	Lang struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"lang"`
	AppName         string `json:"appname"`
	DisplayLanguage string `json:"display_language"`
	ColorTheme      struct {
		Title string `json:"title"`
		Light string `json:"light"`
		Dark  string `json:"dark"`
	} `json:"color_theme"`
	ApiService struct {
		Start   string `json:"start"`
		Stop    string `json:"stop"`
		Swagger string `json:"swagger"`
	} `json:"api_service"`
	Quit       string `json:"quit"`
	QuitDialog struct {
		Message       string `json:"message"`
		DefaultButton string `json:"default_button"`
		CancelButton  string `json:"cancel_button"`
	} `json:"quit_dialog"`
}

type localeMap map[string]locale

func load() (tm localeMap, availableLanguages []string) {
	tm = make(localeMap)
	files, _ := locales.ReadDir(dirLocales)
	for _, f := range files {
		if !f.IsDir() { // load only locale JSON file
			t := locale{}
			data, _ := locales.ReadFile(filepath.Join(dirLocales, f.Name()))
			json.Unmarshal(data, &t)
			tm[t.Lang.Code] = t
			availableLanguages = append(availableLanguages, t.Lang.Code)
		}
	}
	return
}
