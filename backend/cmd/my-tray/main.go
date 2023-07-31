package main

import (
	"my-app/backend/cmd/my-tray/menuitem"
	"my-app/backend/pkg/tray"
)

func main() {
	tray.Run(menuitem.Root())
}
