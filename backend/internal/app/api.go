package app

import (
	app "my-app/backend/api"
	"my-app/backend/internal/crud"
	"my-app/backend/pkg/api"
	"my-app/backend/pkg/funcs"
)

const (
	DefaultPortHttp  uint16 = 10080
	DefaultPortHttps uint16 = 10443
)

func StartAPI() bool {
	dirCerts, err := funcs.GetPathStartedFromExecutable("Certs")
	if err != nil {
		panic(err)
	}

	webPortHttp, _, err := crudOption.GetOrSaveUint16ByOptionName(crud.OptionNameWebPortHttp, DefaultPortHttp)
	if err != nil {
		panic(err)
	}
	webPortHttps, _, err := crudOption.GetOrSaveUint16ByOptionName(crud.OptionNameWebPortHttps, DefaultPortHttps)
	if err != nil {
		panic(err)
	}
	webDirCerts, _, err := crudOption.GetOrSaveByOptionName(crud.OptionNameWebDirCerts, dirCerts, true)
	if err != nil {
		panic(err)
	}
	webHostWhitelist, _, err := crudOption.GetOrSaveStringsByOptionName(crud.OptionNameWebHostWhitelist, []string{})
	if err != nil {
		panic(err)
	}

	return web.Start(api.NewConfig(&api.Config{
		IsDev: cfg.IsDev,
		Log:   logger,
		Http: api.ConfigHttp{
			Port: webPortHttp,
		},
		Https: api.ConfigHttps{
			Port:          webPortHttps,
			HostWhitelist: webHostWhitelist,
			DirCerts:      webDirCerts,
		},
		Setup: app.SETUP(),
	}))
}
