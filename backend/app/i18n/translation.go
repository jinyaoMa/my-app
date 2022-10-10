package i18n

type Translation struct {
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
		En    string `json:"en"`
		Zh    string `json:"zh"`
	} `json:"display_language"`
	ColorTheme struct {
		Label  string `json:"label"`
		Title  string `json:"title"`
		System string `json:"system"`
		Light  string `json:"light"`
		Dark   string `json:"dark"`
	} `json:"color_theme"`
	WebService struct {
		Label     string `json:"label"`
		Enabled   string `json:"enabled"`
		Disabled  string `json:"disabled"`
		VitePress string `json:"vitepress"`
		Swagger   string `json:"swagger"`
		Start     string `json:"start"`
		Stop      string `json:"stop"`
	} `json:"web_service"`
	QuitDialog struct {
		Message       string `json:"message"`
		DefaultButton string `json:"default_button"`
		CancelButton  string `json:"cancel_button"`
	} `json:"quit_dialog"`
}
