package app

import (
	"my-app/backend.new/app/types"
	"my-app/backend.new/model"
	"my-app/backend.new/utils"

	"gorm.io/gorm"
)

type config struct {
	db          *gorm.DB
	options     model.Options
	optionPairs map[types.ConfigName]string
}

// default config
func DefaultConfig(db *gorm.DB) *config {
	c := &config{
		db: db,
		// default config options
		options: model.Options{
			{
				Name:  types.ConfigNameDisplayLanguage,
				Value: "",
			},
			{
				Name:  types.ConfigNameColorTheme,
				Value: types.ColorThemeSystem.ToString(),
			},
			{
				Name:  types.ConfigNameLogFile,
				Value: utils.Utils().GetExecutablePath(utils.Utils().GetExecutableFileName("log")),
			},
			{
				Name:  types.ConfigNameDirLanguages,
				Value: utils.Utils().GetExecutablePath("Languages"),
			},
			{
				Name:  types.ConfigNameDirAssets,
				Value: utils.Utils().GetExecutablePath("Assets"),
			},
			{
				Name:  types.ConfigNameDirUserData,
				Value: utils.Utils().GetExecutablePath("UserData"),
			},
			{
				Name:  types.ConfigNameDirDocs,
				Value: utils.Utils().GetExecutablePath("Docs"),
			},
			{
				Name:  types.ConfigNameWebAutoStart,
				Value: types.BooleanFalse,
			},
			{
				Name:  types.ConfigNameWebPortHttp,
				Value: types.Port(10080).ToString(),
			},
			{
				Name:  types.ConfigNameWebPortHttps,
				Value: types.Port(10443).ToString(),
			},
			{
				Name:  types.ConfigNameWebDirCerts,
				Value: utils.Utils().GetExecutablePath("Certs"),
			},
		},
	}

	return c.generateOptionPairs()
}

// load config from database
func LoadConfig(db *gorm.DB) *config {
	c := DefaultConfig(db)

	if c.options.FindAndSave(db) {
		for _, opt := range c.options {
			c.optionPairs[opt.Name] = opt.Value
		}
	} else {
		utils.Utils().Panic("failed to load config: db error")
	}

	return c
}

// Get get value of an option by the given option name
func (c *config) Get(name types.ConfigName) string {
	if v, ok := c.optionPairs[name]; ok {
		return v
	}
	return ""
}

// Set set new value of an option by the given option name
func (c *config) Set(name types.ConfigName, newValue string) (ok bool) {
	// update db
	opt := model.Option{
		Name:  name,
		Value: newValue,
	}
	if ok = opt.FindByNameAndSave(c.db); !ok {
		return
	}

	// update config
	for _, opt := range c.options {
		if opt.Name == name {
			opt.Value = newValue
			c.optionPairs[name] = newValue
			return ok && true
		}
	}
	return
}

// List get the option list of config
func (c *config) Options() model.Options {
	return c.options
}

// Map get the option pairs of config
func (c *config) OptionPairs() map[types.ConfigName]string {
	return c.optionPairs
}

// generateMap generate option pairs from the option list
func (c *config) generateOptionPairs() *config {
	c.optionPairs = make(map[types.ConfigName]string)
	for _, opt := range c.options {
		c.optionPairs[opt.Name] = opt.Value
	}
	return c
}
