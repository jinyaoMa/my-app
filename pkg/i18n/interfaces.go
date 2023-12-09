package i18n

type ITranslation interface {
	GetLocale() string
	SetLocale(locale string)
	GetDisplayName() string
	SetDisplayName(name string)
}

type II18n[TTranslation ITranslation] interface {
	Load() error
	Locale(locales ...string) (currentLocale string)
	Translation(locales ...string) (translation TTranslation)
	Translations(locales ...string) (translations []TTranslation)
}
