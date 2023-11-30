package i18n

type ITranslation interface {
	GetLocale() string
	SetLocale(locale string)
	GetDisplayName() string
	SetDisplayName(name string)
}
