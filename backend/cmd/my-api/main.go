package main

import (
	"my-app/backend/internal/app"
	"my-app/backend/internal/crud"
	"my-app/backend/internal/vmodel"
	"my-app/backend/pkg/api"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {
	crudOption := crud.NewOption(app.DB())
	portHttp, _, _ := crudOption.GetUint16ByOptionName(vmodel.OptionNameWebPortHttp)
	portHttps, _, _ := crudOption.GetUint16ByOptionName(vmodel.OptionNameWebPortHttps)

	s := app.API()
	if s.Start(&api.Config{
		IsDev: true,
		Log:   app.LOG(),
		Http: api.ConfigHttp{
			Port: portHttp,
		},
		Https: api.ConfigHttps{
			Port: portHttps,
		},
		Setup: func(app *fiber.App) *fiber.App {
			return app
		},
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
