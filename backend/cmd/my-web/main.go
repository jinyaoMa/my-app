package main

import (
	"my-app/backend/internal/app"
	"my-app/backend/internal/crud"
	"my-app/backend/internal/vmodel"
	"my-app/backend/pkg/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {
	optionService := crud.NewOptionService(app.Db())
	portHttp, _, _ := optionService.GetUint16ByOptionName(vmodel.OptionNameWebPortHttp)
	portHttps, _, _ := optionService.GetUint16ByOptionName(vmodel.OptionNameWebPortHttps)

	s := app.Web()
	if s.Start(&server.Option{
		IsDev:  true,
		Logger: app.Log(),
		Http: server.OptionHttp{
			Port: portHttp,
		},
		Https: server.OptionHttps{
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
