package tray

import (
	"context"
	"embed"
	"my-app/backend/pkg/utils"
	"my-app/backend/tray/menus"

	"github.com/getlantern/systray"
)

// web service menu item
const (
	MniWebServiceVitePress = "vitepress"
	MniWebServiceSwagger   = "swagger"
	MniWebServiceStart     = "start"
	MniWebServiceStop      = "stop"
)

//go:embed icons
var icons embed.FS

var instance *tray

type tray struct {
	ctx             context.Context // ctx from wails
	icons           *utils.EmbedFs
	openWindow      *menus.MenuItem
	webService      *menus.MenuSwitch
	displayLanguage *menus.MenuSelect
	colorTheme      *menus.MenuSelect
	quit            *menus.MenuItem
}

func init() {
	instance = &tray{
		icons: utils.NewEmbedFs(icons, "icons"),
	}
	systray.Register(instance.onReady, nil)
}

func Tray(ctxs ...context.Context) *tray {
	if len(ctxs) > 0 {
		instance.ctx = ctxs[0]
	}
	return instance
}
