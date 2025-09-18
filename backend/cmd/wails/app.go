package main

import (
	"net/http"
	"strings"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/services/log"
	"majinyao.cn/my-app/backend/cmd/wails/services"
	"majinyao.cn/my-app/backend/internal/app"
)

const (
	AppName        = "My Application"
	AppDescription = "My Application with many features"
)

var App *application.App

func newApp() *application.App {
	App = application.New(application.Options{
		Name:        AppName,
		Description: AppDescription,
		Services: []application.Service{
			application.NewService(log.New()),
			application.NewService(&services.GreetService{}),
			application.NewService(&services.UserService{}),
		},
		Assets: application.AssetOptions{
			Middleware: func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// Let Wails handle the `/wails/` route
					if strings.HasPrefix(r.URL.Path, "/wails/") {
						next.ServeHTTP(w, r)
						return
					}

					// Let API handle everything else
					app.API.ServeHTTP(w, r)
				})
			},
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
			ActivationPolicy: application.ActivationPolicyAccessory,
		},
	})
	return App
}
