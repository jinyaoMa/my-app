package menus

import (
	"my-app/backend/pkg/i18n"

	"github.com/getlantern/systray"
)

type ApiServiceListener struct {
	OnOpenSwagger func()
	OnStart       func() (ok bool)
	OnStop        func() (ok bool)
}

type ApiService struct {
	isWatched bool
	chanStop  chan struct{}
	isEnabled bool
	swagger   *systray.MenuItem
	start     *systray.MenuItem
	stop      *systray.MenuItem
}

func NewApiService() *ApiService {
	return &ApiService{
		chanStop: make(chan struct{}, 1),
		swagger:  systray.AddMenuItem("", ""),
		start:    systray.AddMenuItem("", ""),
		stop:     systray.AddMenuItem("", ""),
	}
}

func (as *ApiService) SetIconSwagger(templateIconBytes []byte, regularIconBytes []byte) *ApiService {
	as.swagger.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return as
}

func (as *ApiService) SetIconStart(templateIconBytes []byte, regularIconBytes []byte) *ApiService {
	as.start.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return as
}

func (as *ApiService) SetIconStop(templateIconBytes []byte, regularIconBytes []byte) *ApiService {
	as.stop.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return as
}

func (as *ApiService) SetLocale(locale i18n.Locale) *ApiService {
	as.start.SetTitle(locale.ApiService.Start)
	as.start.SetTooltip(locale.ApiService.Start)
	as.stop.SetTitle(locale.ApiService.Stop)
	as.stop.SetTooltip(locale.ApiService.Stop)
	as.swagger.SetTitle(locale.ApiService.Swagger)
	as.swagger.SetTooltip(locale.ApiService.Swagger)
	if as.isEnabled {
		as.start.Hide()
		as.stop.Show()
		as.swagger.Show()
	} else {
		as.start.Show()
		as.stop.Hide()
		as.swagger.Hide()
	}
	return as
}

func (as *ApiService) Watch(listener ApiServiceListener) *ApiService {
	if as.isWatched {
		return as
	}

	as.isWatched = true
	go func() {
		for {
			select {
			case <-as.start.ClickedCh:
				if listener.OnStart() {
					as.start.Hide()
					as.stop.Show()
					as.swagger.Show()
					as.isEnabled = true
				}
			case <-as.stop.ClickedCh:
				if listener.OnStop() {
					as.start.Show()
					as.stop.Hide()
					as.swagger.Hide()
					as.isEnabled = false
				}
			case <-as.swagger.ClickedCh:
				listener.OnOpenSwagger()
			case <-as.chanStop:
				return
			}
		}
	}()
	return as
}

func (as *ApiService) StopWatch() *ApiService {
	if as.isWatched {
		as.chanStop <- struct{}{}
	}
	return as
}

func (as *ApiService) ClickOpenSwagger() *ApiService {
	as.swagger.ClickedCh <- struct{}{}
	return as
}

func (as *ApiService) ClickStart() *ApiService {
	as.start.ClickedCh <- struct{}{}
	return as
}

func (as *ApiService) ClickStop() *ApiService {
	as.stop.ClickedCh <- struct{}{}
	return as
}
