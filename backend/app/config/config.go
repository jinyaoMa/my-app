package config

import (
	"my-app/backend/database/option"
	"my-app/backend/pkg/utils"
)

// option names
const (
	CfgDisplayLanguage = "DisplayLanguage"
	CfgColorTheme      = "ColorTheme"
	CfgLogPath         = "LogPath"
)

// display language option
const (
	DisplayLanguageEn = "en"
	DisplayLanguageZh = "zh"
)

// color theme option
const (
	ColorThemeSystem = "system"
	ColorThemeLight  = "light"
	ColorThemeDark   = "dark"
)

type Config struct {
	DisplayLanguage string
	ColorTheme      string
	LogPath         string
	*WebConfig
}

// default config
func DefaultConfig() *Config {
	return &Config{
		DisplayLanguage: DisplayLanguageEn,
		ColorTheme:      ColorThemeSystem,
		LogPath:         utils.GetExecutablePath("MyApp.log"),
		WebConfig: &WebConfig{
			AutoStart: "false",
			PortHttp:  ":10080",
			PortHttps: ":10443",
			DirCerts:  "",
		},
	}
}

// load config from database
func LoadConfig() *Config {
	cfg := DefaultConfig()

	var opts option.Options
	if opts.Find() {
		cfg.load(opts)
	} else {
		cfg.Save()
	}

	return cfg
}

// save all config into database
func (c *Config) Save() (ok bool) {
	var options option.Options
	for _, opt := range [][]string{
		{CfgDisplayLanguage, c.DisplayLanguage},
		{CfgColorTheme, c.ColorTheme},
		{CfgLogPath, c.LogPath},
		{CfgWebAutoStart, c.AutoStart},
		{CfgWebPortHttp, c.PortHttp},
		{CfgWebPortHttps, c.PortHttps},
		{CfgWebDirCerts, c.DirCerts},
	} {
		options = append(options, option.Option{
			Name:  opt[0],
			Value: opt[1],
		})
	}
	return options.Save()
}

// update option's value into database and set config's option value
func (c *Config) Update(name string, value string) (ok bool) {
	opt := &option.Option{
		Name: name,
	}
	return opt.Update(value) && c.set(name, value)
}

// load options into config and save new options into database
func (c *Config) load(options option.Options) (ok bool) {
	for _, option := range options {
		if !c.set(option.Name, option.Value) {
			return false
		}
	}

	optionPairs := map[string]string{
		CfgDisplayLanguage: c.DisplayLanguage,
		CfgColorTheme:      c.ColorTheme,
		CfgLogPath:         c.LogPath,
		CfgWebAutoStart:    c.AutoStart,
		CfgWebPortHttp:     c.PortHttp,
		CfgWebPortHttps:    c.PortHttps,
		CfgWebDirCerts:     c.DirCerts,
	}

	var newOptions option.Options
	for name, value := range optionPairs {
		optionNotExist := true
		for _, option := range options {
			// update exist options
			if option.Name == name {
				option.Value = value
				optionNotExist = false
				break
			}
		}
		if optionNotExist {
			// initialize new options
			newOptions = append(newOptions, option.Option{
				Name:  name,
				Value: value,
			})
		}
	}

	options = append(options, newOptions...)
	return options.Save()
}

// set config's option value
func (c *Config) set(name string, value string) (ok bool) {
	switch name {
	default:
		return false
	case CfgDisplayLanguage:
		c.DisplayLanguage = value
	case CfgColorTheme:
		c.ColorTheme = value
	case CfgLogPath:
		c.LogPath = value
	case CfgWebAutoStart:
		c.AutoStart = value
	case CfgWebPortHttp:
		c.PortHttp = value
	case CfgWebPortHttps:
		c.PortHttps = value
	case CfgWebDirCerts:
		c.DirCerts = value
	}
	return true
}

func (c *Config) Web() *WebConfig {
	return c.WebConfig
}
