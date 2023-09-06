package main

import (
	"my-app/backend/cmd/my-tray/menuitems"
	"my-app/backend/pkg/tray"
)

func main() {
	tray.Run(menuitems.Root())
}
