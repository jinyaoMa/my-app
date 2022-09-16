package app

import (
	"log"
	"my-app/backend/model"
	"my-app/backend/pkg/i18n"
)

const (
	CfgTotalNumberOfOptions = 5
	CfgLanguage             = "Config.Language"
	CfgTheme                = "Config.Theme"
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
		Theme:    ColorThemeSystem,
		Web: &WebConfig{
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
		log.Fatalf("fail to load options: %+v\n", result.Error)
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
		{CfgLanguage, c.Language},
		{CfgTheme, c.Theme},
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
		log.Fatalf("fail to update options: %+v\n", result.Error)
	}
}

func (c *Config) saveOptions(options model.MyOptions) {
	options = append(options, model.MyOption{
		Name:  CfgLanguage,
		Value: c.Language,
	})
	options = append(options, model.MyOption{
		Name:  CfgTheme,
		Value: c.Theme,
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
		log.Fatalf("fail to save options: %+v\n", result.Error)
	}
}

func (c *Config) loadOptions(options model.MyOptions) {
	for _, option := range options {
		switch option.Name {
		case CfgLanguage:
			c.Language = option.Value
		case CfgTheme:
			c.Theme = option.Value
		case CfgWebPortHttp:
			c.Web.PortHttp = option.Value
		case CfgWebPortHttps:
			c.Web.PortHttps = option.Value
		case CfgWebDirCerts:
			c.Web.DirCerts = option.Value
		}
	}
}
