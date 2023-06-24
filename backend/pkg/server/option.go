package server

import (
	"my-app/backend/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/imdario/mergo"
)

type Option struct {
	IsDev  bool
	Logger logger.Interface
	Http   OptionHttp
	Https  OptionHttps
	Setup  func(app *fiber.App) *fiber.App
}

type OptionHttp struct {
	Port uint16
}

type OptionHttps struct {
	Port          uint16
	HostWhitelist []string
	DirCerts      string
}

func DefaultOption() *Option {
	return &Option{
		IsDev:  false,
		Logger: logger.New(logger.DefaultOption()),
		Http: OptionHttp{
			Port: 10080,
		},
		Https: OptionHttps{
			Port:          10443,
			HostWhitelist: []string{},
			DirCerts:      "",
		},
		Setup: func(app *fiber.App) *fiber.App {
			return app
		},
	}
}

func NewOption(dst *Option) *Option {
	src := DefaultOption()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
