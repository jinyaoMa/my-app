package wserver

import (
	"my-app/pkg/base"

	"github.com/gofiber/fiber/v2"
)

type WebServerOptions struct {
	base.Options
	IsDev bool
	Http  WebServerOptionsHttp
	Https WebServerOptionsHttps
	Setup func(app *fiber.App) *fiber.App
}

type WebServerOptionsHttp struct {
	Port uint16
}

type WebServerOptionsHttps struct {
	WebServerOptionsHttp
	HostWhitelist []string
	DirCerts      string
}

func DefaultWebServerOptions() *WebServerOptions {
	return &WebServerOptions{
		IsDev: false,
		Http: WebServerOptionsHttp{
			Port: 10080,
		},
		Https: WebServerOptionsHttps{
			WebServerOptionsHttp: WebServerOptionsHttp{
				Port: 10443,
			},
			HostWhitelist: []string{},
			DirCerts:      "",
		},
		Setup: func(app *fiber.App) *fiber.App {
			return app
		},
	}
}

func NewWebServerOptions(dst *WebServerOptions) (*WebServerOptions, error) {
	return base.SimpleMerge(DefaultWebServerOptions(), dst)
}
