package main

import (
	"my-app/backend/internal/menuitems"
	"my-app/backend/pkg/tray"
)

func main() {
	tray.Run(menuitems.Root())
}
