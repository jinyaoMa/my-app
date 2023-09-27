package main

import (
	"my-app/backend/api"
	"my-app/backend/internal/app"
	"my-app/backend/internal/crud"
	"my-app/backend/pkg/funcs"
	"my-app/backend/pkg/web"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	crudOption := crud.NewOption(app.DB())
	portHttp, _, _ := crudOption.GetOrSaveUint16ByOptionName(crud.OptionNameWebPortHttp, 10080)
	portHttps, _, _ := crudOption.GetOrSaveUint16ByOptionName(crud.OptionNameWebPortHttps, 10443)
	dirCerts, _ := funcs.GetPathStartedFromExecutable("Certs")

	s := app.SERVER()
	if s.Start(&web.Config{
		IsDev: true,
		Log:   app.LOG(),
		Http: web.ConfigHttp{
			Port: portHttp,
		},
		Https: web.ConfigHttps{
			Port:     portHttps,
			DirCerts: dirCerts,
		},
		Setup: api.Setup(),
	}) {
		println("start")
	}

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received

	if s.Stop(func() {
		println("try to stop...")
	}) {
		println("exit")
	}
}
