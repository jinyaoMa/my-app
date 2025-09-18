package i18n

type Options struct {
	Fallback   string `json:"fallback"`   // fallback language if the language is not found
	Directory  string `json:"directory"`  // path to the Language files
	DefineJson string `json:"defineJson"` // define.json file name, default is "define.json"
}
