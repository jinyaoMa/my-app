package app

import (
	"my-app/backend/i18n"
	"my-app/backend/tray/menus"
)

const (
	CfgLanguage    = "Config.Language"
	CfgTheme       = "Config.Theme"
	CfgWebPortHttp = "Config.Web.PortHttp"
	CfgPortHttps   = "Config.Web.PortHttps"
	CfgWebDirCerts = "Config.Web.DirCerts"
)

type Config struct {
	Language string
	Theme    string
	Web      *WebConfig
}

type WebConfig struct {
	PortHttp  string
	PortHttps string
	DirCerts  string
}

func DefaultConfig() *Config {
	return &Config{
		Language: i18n.En,
		Theme:    menus.ColorThemeSystem,
		Web: &WebConfig{
			PortHttp:  ":10080",
			PortHttps: ":10443",
			DirCerts:  "",
		},
	}
}
