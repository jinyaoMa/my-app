package app

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/vmodel"
	"my-app/backend/pkg/api"
	"my-app/backend/pkg/funcs"

	"github.com/gofiber/fiber/v2"
)

const (
	DefaultPortHttp  uint16 = 10080
	DefaultPortHttps uint16 = 10443
)

func StartAPI() bool {
	var err error
	var webPortHttp, webPortHttps, webDirCerts, webHostWhitelist *entity.Option

	webPortHttp, err = crudOption.GetByOptionName(vmodel.OptionNameWebPortHttp)
	if err != nil {
		webPortHttp = &entity.Option{
			Key:   vmodel.OptionNameWebPortHttp,
			Value: vmodel.OptionValueWebPortString(DefaultPortHttp),
		}
	}
	webPortHttps, err = crudOption.GetByOptionName(vmodel.OptionNameWebPortHttps)
	if err != nil {
		webPortHttps = &entity.Option{
			Key:   vmodel.OptionNameWebPortHttps,
			Value: vmodel.OptionValueWebPortString(DefaultPortHttps),
		}
	}
	webDirCerts, err = crudOption.GetByOptionName(vmodel.OptionNameWebDirCerts)
	if err != nil {
		dirCerts, _ := funcs.GetPathStartedFromExecutable("Certs")
		webDirCerts = &entity.Option{
			Key:   vmodel.OptionNameWebDirCerts,
			Value: dirCerts,
		}
	}
	webHostWhitelist, err = crudOption.GetByOptionName(vmodel.OptionNameWebHostWhitelist)
	if err != nil {
		webHostWhitelist = &entity.Option{
			Key:   vmodel.OptionNameWebHostWhitelist,
			Value: "",
		}
	}

	return web.Start(api.NewConfig(&api.Config{
		IsDev: cfg.IsDev,
		Log:   logger,
		Http: api.ConfigHttp{
			Port: vmodel.OptionValueWebPort(webPortHttp.Value, DefaultPortHttp),
		},
		Https: api.ConfigHttps{
			Port:          vmodel.OptionValueWebPort(webPortHttps.Value, DefaultPortHttps),
			HostWhitelist: vmodel.OptionValueCommaList(webHostWhitelist.Value),
			DirCerts:      webDirCerts.Value,
		},
		Setup: func(app *fiber.App) *fiber.App {
			return app
		},
	}))
}
