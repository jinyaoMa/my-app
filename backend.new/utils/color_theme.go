package utils

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	ColorThemeSystem ColorTheme = "system"
	ColorThemeLight  ColorTheme = "light"
	ColorThemeDark   ColorTheme = "dark"
)

type ColorTheme string

func (ct ColorTheme) ToString() string {
	return string(ct)
}

func (ct ColorTheme) Change(ctx context.Context, newValue ColorTheme) ColorTheme {
	ct = newValue
	switch ct {
	case ColorThemeSystem:
		runtime.WindowSetSystemDefaultTheme(ctx)
	case ColorThemeLight:
		runtime.WindowSetLightTheme(ctx)
	case ColorThemeDark:
		runtime.WindowSetDarkTheme(ctx)
	}
	return ct
}
