package app

import (
	"my-app/backend.new/model"
	"my-app/backend.new/utils"

	"gorm.io/gorm"
)

type Config struct {
	db *gorm.DB

	_list model.Options
	_map  map[model.OptionName]string
}

// default config
func DefaultConfig(db *gorm.DB) *Config {
	c := &Config{
		db: db,
	}

	// default config as list
	c._list = model.Options{
		{
			Name:  model.OptionDisplayLanguage,
			Value: "en",
		},
		{
			Name:  model.OptionColorTheme,
			Value: utils.ColorThemeSystem.ToString(),
		},
		{
			Name:  model.OptionFileLog,
			Value: utils.Utils().GetExecutablePath("MyApp.log"),
		},
		{
			Name:  model.OptionDirLanguages,
			Value: utils.Utils().GetExecutablePath("Languages"),
		},
		{
			Name:  model.OptionDirAssets,
			Value: utils.Utils().GetExecutablePath("Assets"),
		},
		{
			Name:  model.OptionDirUserData,
			Value: utils.Utils().GetExecutablePath("UserData"),
		},
		{
			Name:  model.OptionDirDocs,
			Value: utils.Utils().GetExecutablePath("Docs"),
		},
		{
			Name:  model.OptionWebAutoStart,
			Value: "false",
		},
		{
			Name:  model.OptionWebPortHttp,
			Value: ":10080",
		},
		{
			Name:  model.OptionWebPortHttps,
			Value: ":10443",
		},
		{
			Name:  model.OptionWebDirCerts,
			Value: "",
		},
	}
	c.generateMap()

	return c
}

// load config from database
func LoadConfig(db *gorm.DB) *Config {
	c := DefaultConfig(db)

	if c._list.FindAndSave(db) {
		for _, opt := range c._list {
			c._map[opt.Name] = opt.Value
		}
	}

	return c
}

// Get get value of an option by the given option name
func (c *Config) Get(name model.OptionName) string {
	return c._map[name]
}

// Set set new value of an option by the given option name
func (c *Config) Set(name model.OptionName, newValue string) (ok bool) {
	if c.db == nil {
		return false
	}

	// update db
	opt := model.Option{
		Name:  name,
		Value: newValue,
	}
	if ok = opt.FindByNameAndSave(c.db); !ok {
		return
	}

	// update config
	for _, opt := range c._list {
		if opt.Name == name {
			opt.Value = newValue
			c._map[name] = newValue
			return ok && true
		}
	}
	return
}

// List get the list of config
func (c *Config) List() model.Options {
	return c._list
}

// Map get the map of config
func (c *Config) Map() map[model.OptionName]string {
	return c._map
}

// generateMap generate map from the list
func (c *Config) generateMap() {
	c._map = make(map[model.OptionName]string)
	for _, opt := range c._list {
		c._map[opt.Name] = opt.Value
	}
}
