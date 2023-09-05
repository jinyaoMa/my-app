package vmodel

import (
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// option values
const (
	OptionValueColorThemeSystem = "system"
	OptionValueColorThemeLight  = "light"
	OptionValueColorThemeDark   = "dark"
)

func OptionValueColorTheme(t string, def windows.Theme) windows.Theme {
	switch t {
	case OptionValueColorThemeLight:
		return windows.Light
	case OptionValueColorThemeDark:
		return windows.Dark
	case OptionValueColorThemeSystem:
		return windows.SystemDefault
	default:
		return def
	}
}

func OptionValueColorThemeString(t windows.Theme, def string) string {
	switch t {
	case windows.Light:
		return OptionValueColorThemeLight
	case windows.Dark:
		return OptionValueColorThemeDark
	case windows.SystemDefault:
		return OptionValueColorThemeSystem
	default:
		return def
	}
}

func OptionValueBool(b string) bool {
	return b == "1"
}

func OptionValueBoolString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func OptionValueWebPort(p string, def uint16) uint16 {
	v, err := strconv.ParseUint(p, 10, 16)
	if err != nil {
		return def
	}
	return uint16(v)
}

func OptionValueWebPortString(p uint16) string {
	return strconv.FormatUint(uint64(p), 10)
}

func OptionValueCommaList(s string) []string {
	return strings.Split(s, ",")
}

func OptionValueCommaListString(s ...string) string {
	return strings.Join(s, ",")
}
