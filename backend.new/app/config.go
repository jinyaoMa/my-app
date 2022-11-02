package app

import (
	"my-app/backend.new/app/types"
	"my-app/backend.new/model"
	"my-app/backend.new/utils"

	"gorm.io/gorm"
)

type Config struct {
	db    *gorm.DB
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
			Name:  model.OptionNameDisplayLanguage,
			Value: "",
		},
		{
			Name:  model.OptionNameColorTheme,
			Value: types.ColorThemeDefault.ToString(),
		},
		{
			Name:  model.OptionNameFileLog,
			Value: utils.Utils().GetExecutablePath("MyApp.log"),
		},
		{
			Name:  model.OptionNameDirLanguages,
			Value: utils.Utils().GetExecutablePath("Languages"),
		},
		{
			Name:  model.OptionNameDirAssets,
			Value: utils.Utils().GetExecutablePath("Assets"),
		},
		{
			Name:  model.OptionNameDirUserData,
			Value: utils.Utils().GetExecutablePath("UserData"),
		},
		{
			Name:  model.OptionNameDirDocs,
			Value: utils.Utils().GetExecutablePath("Docs"),
		},
		{
			Name:  model.OptionNameWebAutoStart,
			Value: types.BoolFalse.ToString(),
		},
		{
			Name:  model.OptionNameWebPortHttp,
			Value: types.NewPort(":10080").ToString(),
		},
		{
			Name:  model.OptionNameWebPortHttps,
			Value: types.NewPort(":10443").ToString(),
		},
		{
			Name:  model.OptionNameWebDirCerts,
			Value: utils.Utils().GetExecutablePath("Certs"),
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
	} else {
		utils.Utils().PanicLogger().Fatalln("failed to load config")
	}

	return c
}

// Get get value of an option by the given option name
func (c *Config) Get(name model.OptionName) string {
	if v, ok := c._map[name]; ok {
		return v
	}
	return ""
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
