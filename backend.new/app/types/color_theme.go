package types

const (
	ColorThemeSystem = ColorTheme("system")
	ColorThemeLight  = ColorTheme("light")
	ColorThemeDark   = ColorTheme("dark")
)

type ColorTheme string

func ParseColorTheme(theme string) ColorTheme {
	switch theme {
	case ColorThemeLight.ToString():
		return ColorThemeLight
	case ColorThemeDark.ToString():
		return ColorThemeDark
	}
	return ColorThemeSystem
}

func (ct ColorTheme) ToString() string {
	return string(ct)
}
