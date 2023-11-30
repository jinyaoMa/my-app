package i18n

type Translation struct {
	Locale      string `json:"locale"` // locale must be the same as the json source filename without extension
	DisplayName string `json:"displayName"`
}

// GetDisplayName implements ITranslation.
func (translation *Translation) GetDisplayName() string {
	return translation.DisplayName
}

// GetLocale implements ITranslation.
func (translation *Translation) GetLocale() string {
	return translation.Locale
}

// SetDisplayName implements ITranslation.
func (translation *Translation) SetDisplayName(name string) {
	translation.DisplayName = name
}

// SetLocale implements ITranslation.
func (translation *Translation) SetLocale(locale string) {
	translation.Locale = locale
}

func NewTranslation(translation *Translation) ITranslation {
	return translation
}
