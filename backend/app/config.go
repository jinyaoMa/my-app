package app

import (
	"my-app/backend/model"
	"my-app/backend/pkg/i18n"
	"my-app/backend/pkg/utils"
)

const (
	CfgTotalNumberOfOptions = 7
	CfgDisplayLanguage      = "Config.DisplayLanguage"
	CfgColorTheme           = "Config.ColorTheme"
	CfgLogPath              = "Config.LogPath"
	CfgWebAutoStart         = "Config.Web.AutoStart"
	CfgWebPortHttp          = "Config.Web.PortHttp"
	CfgWebPortHttps         = "Config.Web.PortHttps"
	CfgWebDirCerts          = "Config.Web.DirCerts"
)

const (
	ColorThemeSystem = "system"
	ColorThemeLight  = "light"
	ColorThemeDark   = "dark"
)

type Config struct {
	DisplayLanguage string
	ColorTheme      string
	LogPath         string
	Web             *WebConfig
}

type WebConfig struct {
	AutoStart string
	PortHttp  string
	PortHttps string
	DirCerts  string
}

func DefaultConfig() *Config {
	return &Config{
		DisplayLanguage: i18n.En,
		ColorTheme:      ColorThemeSystem,
		LogPath:         utils.GetExecutablePath("MyApp.log"),
		Web: &WebConfig{
			AutoStart: "true",
			PortHttp:  ":10080",
			PortHttps: ":10443",
			DirCerts:  "",
		},
	}
}

func LoadConfig() *Config {
	cfg := DefaultConfig()

	var options model.MyOptions
	result := options.Load()
	if result.Error != nil {
		instance.AppLog().Fatalf("fail to load options: %+v\n", result.Error)
	}
	if result.RowsAffected == 0 {
		// options not yet generated and stored
		cfg.saveOptions(options)
	} else if result.RowsAffected < CfgTotalNumberOfOptions {
		cfg.loadOptions(options)
		cfg.updateOptions(options)
	} else {
		cfg.loadOptions(options)
	}
	return cfg
}

func (c *Config) updateOptions(options model.MyOptions) {
	optionPairs := [][]string{
		{CfgDisplayLanguage, c.DisplayLanguage},
		{CfgColorTheme, c.ColorTheme},
		{CfgLogPath, c.LogPath},
		{CfgWebAutoStart, c.Web.AutoStart},
		{CfgWebPortHttp, c.Web.PortHttp},
		{CfgWebPortHttps, c.Web.PortHttps},
		{CfgWebDirCerts, c.Web.DirCerts},
	}

	var newOptions model.MyOptions
	for _, pair := range optionPairs {
		optionNotExist := true
		for _, option := range options {
			// update exist options
			if option.Name == pair[0] {
				option.Value = pair[1]
				optionNotExist = false
				break
			}
		}
		if optionNotExist {
			// initialize new options
			newOptions = append(newOptions, model.MyOption{
				Name:  pair[0],
				Value: pair[1],
			})
		}
	}

	options = append(options, newOptions...)

	result := options.Save()
	if result.Error != nil {
		instance.AppLog().Fatalf("fail to update options: %+v\n", result.Error)
	}
}

func (c *Config) saveOptions(options model.MyOptions) {
	options = append(options, model.MyOption{
		Name:  CfgDisplayLanguage,
		Value: c.DisplayLanguage,
	})
	options = append(options, model.MyOption{
		Name:  CfgColorTheme,
		Value: c.ColorTheme,
	})
	options = append(options, model.MyOption{
		Name:  CfgLogPath,
		Value: c.LogPath,
	})
	options = append(options, model.MyOption{
		Name:  CfgWebAutoStart,
		Value: c.Web.AutoStart,
	})
	options = append(options, model.MyOption{
		Name:  CfgWebPortHttp,
		Value: c.Web.PortHttp,
	})
	options = append(options, model.MyOption{
		Name:  CfgWebPortHttps,
		Value: c.Web.PortHttps,
	})
	options = append(options, model.MyOption{
		Name:  CfgWebDirCerts,
		Value: c.Web.DirCerts,
	})

	result := options.Save()
	if result.Error != nil {
		instance.AppLog().Fatalf("fail to save options: %+v\n", result.Error)
	}
}

func (c *Config) loadOptions(options model.MyOptions) {
	for _, option := range options {
		switch option.Name {
		case CfgDisplayLanguage:
			c.DisplayLanguage = option.Value
		case CfgColorTheme:
			c.ColorTheme = option.Value
		case CfgLogPath:
			c.LogPath = option.Value
		case CfgWebAutoStart:
			c.Web.AutoStart = option.Value
		case CfgWebPortHttp:
			c.Web.PortHttp = option.Value
		case CfgWebPortHttps:
			c.Web.PortHttps = option.Value
		case CfgWebDirCerts:
			c.Web.DirCerts = option.Value
		}
	}
}
