package api

import (
	"my-app/backend/pkg/logger"

	"dario.cat/mergo"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	IsDev  bool
	Logger logger.Interface
	Http   ConfigHttp
	Https  ConfigHttps
	Setup  func(app *fiber.App) *fiber.App
}

type ConfigHttp struct {
	Port uint16
}

type ConfigHttps struct {
	Port          uint16
	HostWhitelist []string
	DirCerts      string
}

func DefaultConfig() *Config {
	return &Config{
		IsDev:  false,
		Logger: logger.New(logger.DefaultOption()),
		Http: ConfigHttp{
			Port: 10080,
		},
		Https: ConfigHttps{
			Port:          10443,
			HostWhitelist: []string{},
			DirCerts:      "",
		},
		Setup: func(app *fiber.App) *fiber.App {
			return app
		},
	}
}

func NewConfig(dst *Config) *Config {
	src := DefaultConfig()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
