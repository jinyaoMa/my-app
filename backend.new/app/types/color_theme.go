package types

const (
	ColorThemeDefault = ColorTheme("default")
	ColorThemeLight   = ColorTheme("light")
	ColorThemeDark    = ColorTheme("dark")
)

type ColorTheme string

func ParseColorTheme(theme string) ColorTheme {
	switch theme {
	case ColorThemeLight.ToString():
		return ColorThemeLight
	case ColorThemeDark.ToString():
		return ColorThemeDark
	}
	return ColorThemeDefault
}

func (ct ColorTheme) ToString() string {
	return string(ct)
}
