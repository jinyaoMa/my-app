package app

import (
	"my-app/backend/api"
	"my-app/backend/internal/crud"
	"my-app/backend/pkg/funcs"
	"my-app/backend/pkg/web"
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

	return server.Start(web.NewConfig(&web.Config{
		IsDev: cfg.IsDev,
		Log:   logger,
		Http: web.ConfigHttp{
			Port: webPortHttp,
		},
		Https: web.ConfigHttps{
			Port:          webPortHttps,
			HostWhitelist: webHostWhitelist,
			DirCerts:      webDirCerts,
		},
		Setup: api.SETUP(),
	}))
}
