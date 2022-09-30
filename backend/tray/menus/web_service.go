package menus

import (
	"my-app/backend/pkg/i18n"

	"github.com/getlantern/systray"
)

type WebServiceListener struct {
	OnOpenVitePress func()
	OnOpenSwagger   func()
	OnStart         func() (ok bool, complete func())
	OnStop          func() (ok bool, complete func())
}

type WebService struct {
	isWatched bool
	chanStop  chan struct{}
	isEnabled bool
	vitepress *systray.MenuItem
	swagger   *systray.MenuItem
	start     *systray.MenuItem
	stop      *systray.MenuItem
}

func NewWebService() *WebService {
	return &WebService{
		chanStop:  make(chan struct{}, 1),
		vitepress: systray.AddMenuItem("", ""),
		swagger:   systray.AddMenuItem("", ""),
		start:     systray.AddMenuItem("", ""),
		stop:      systray.AddMenuItem("", ""),
	}
}

func (ws *WebService) IsEnabled() bool {
	return ws.isEnabled
}

func (ws *WebService) SetIconVitePress(templateIconBytes []byte, regularIconBytes []byte) *WebService {
	ws.vitepress.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return ws
}

func (ws *WebService) SetIconSwagger(templateIconBytes []byte, regularIconBytes []byte) *WebService {
	ws.swagger.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return ws
}

func (ws *WebService) SetIconStart(templateIconBytes []byte, regularIconBytes []byte) *WebService {
	ws.start.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return ws
}

func (ws *WebService) SetIconStop(templateIconBytes []byte, regularIconBytes []byte) *WebService {
	ws.stop.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return ws
}

func (ws *WebService) SetLocale() *WebService {
	locale := i18n.I18n().Locale()
	ws.start.SetTitle(locale.WebService.Start)
	ws.start.SetTooltip(locale.WebService.Start)
	ws.stop.SetTitle(locale.WebService.Stop)
	ws.stop.SetTooltip(locale.WebService.Stop)
	ws.vitepress.SetTitle(locale.WebService.VitePress)
	ws.vitepress.SetTooltip(locale.WebService.VitePress)
	ws.swagger.SetTitle(locale.WebService.Swagger)
	ws.swagger.SetTooltip(locale.WebService.Swagger)
	if ws.isEnabled {
		ws.start.Hide()
		ws.stop.Show()
		ws.vitepress.Show()
		ws.swagger.Show()
	} else {
		ws.start.Show()
		ws.stop.Hide()
		ws.vitepress.Hide()
		ws.swagger.Hide()
	}
	return ws
}

func (ws *WebService) Watch(listener WebServiceListener) *WebService {
	if ws.isWatched {
		return ws
	}

	ws.isWatched = true
	go func() {
		for {
			select {
			case <-ws.start.ClickedCh:
				if ok, complete := listener.OnStart(); ok {
					ws.start.Hide()
					ws.stop.Show()
					ws.vitepress.Show()
					ws.swagger.Show()
					ws.isEnabled = true
					complete()
				}
			case <-ws.stop.ClickedCh:
				if ok, complete := listener.OnStop(); ok {
					ws.start.Show()
					ws.stop.Hide()
					ws.vitepress.Hide()
					ws.swagger.Hide()
					ws.isEnabled = false
					complete()
				}
			case <-ws.vitepress.ClickedCh:
				listener.OnOpenVitePress()
			case <-ws.swagger.ClickedCh:
				listener.OnOpenSwagger()
			case <-ws.chanStop:
				return
			}
		}
	}()
	return ws
}

func (ws *WebService) StopWatch() *WebService {
	if ws.isWatched {
		ws.chanStop <- struct{}{}
	}
	return ws
}

func (ws *WebService) ClickOpenVitePress() *WebService {
	ws.vitepress.ClickedCh <- struct{}{}
	return ws
}

func (ws *WebService) ClickOpenSwagger() *WebService {
	ws.swagger.ClickedCh <- struct{}{}
	return ws
}

func (ws *WebService) ClickStart() *WebService {
	ws.start.ClickedCh <- struct{}{}
	return ws
}

func (ws *WebService) ClickStop() *WebService {
	ws.stop.ClickedCh <- struct{}{}
	return ws
}
