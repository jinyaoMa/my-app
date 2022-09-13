package tray

import (
	_ "embed"
	"sync"

	"github.com/getlantern/systray"
)

//go:embed icons/icon.ico
var icon []byte

var (
	once     sync.Once
	instance *tray
)

type tray struct {
}

func Tray() *tray {
	once.Do(func() {
		instance = &tray{}
		systray.Register(instance.onReady, nil)
	})
	return instance
}

func (st *tray) onReady() {
	systray.SetTemplateIcon(icon, icon)
	systray.SetTitle("Webview example")
}
