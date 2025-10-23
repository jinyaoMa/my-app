package main

import (
	"net/http"
	"strings"
	"time"

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
			application.NewService(new(services.FileService)),
			application.NewService(new(services.GreetService)),
			application.NewService(new(services.UserService)),
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
		Windows: application.WindowsOptions{
			WebviewUserDataPath: app.CFG.Get().WebviewUserDataPath,
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
			ActivationPolicy: application.ActivationPolicyAccessory,
		},
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID: "cn.majinyao.my-app.single-instance",
			OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
				if WindowMain != nil {
					WindowMain.EmitEvent("secondInstanceLaunched", data)
					WindowMain.Restore()
					WindowMain.Focus()
				}
			},
			AdditionalData: map[string]string{
				"launchtime": time.Now().Local().String(),
			},
			ExitCode:      0,
			EncryptionKey: [32]byte([]byte("Ml!Zjj@Lfw#Wqq$Wxb%Mjy^&*()_+1234567890-=")[:32]),
		},
	})
	return App
}
