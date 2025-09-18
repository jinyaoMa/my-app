package i18n

type Locale struct {
	Code string `json:"code"` // language code
	Text string `json:"text"` // language display name
	File string `json:"file"` // language file name
}
