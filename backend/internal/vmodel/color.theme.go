package vmodel

import "github.com/wailsapp/wails/v2/pkg/options/windows"

func ColorTheme(t string) windows.Theme {
	switch t {
	case OptionValueColorThemeLight:
		return windows.Light
	case OptionValueColorThemeDark:
		return windows.Dark
	default:
		return windows.SystemDefault
	}
}

func ColorThemeString(t windows.Theme) string {
	switch t {
	case windows.Light:
		return OptionValueColorThemeLight
	case windows.Dark:
		return OptionValueColorThemeDark
	default:
		return OptionValueColorThemeSystem
	}
}
