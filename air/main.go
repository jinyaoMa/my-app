package main

import (
	"fmt"
	"my-app/backend/app"
	"my-app/backend/web"
)

func main() {
	go fmt.Printf("\n[GIN-debug]  ➜  Local: https://localhost%s/swagger/index.html \n\n", app.App().Config().PortHttps)
	web.Web().Air()
}
