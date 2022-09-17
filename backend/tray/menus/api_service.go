package menus

import (
	"my-app/backend/pkg/i18n"

	"github.com/getlantern/systray"
)

type ApiServiceListener struct {
	OnOpenSwagger func()
	OnStart       func() (ok bool, complete func())
	OnStop        func() (ok bool, complete func())
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

func (as *ApiService) IsEnabled() bool {
	return as.isEnabled
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

func (as *ApiService) SetLocale() *ApiService {
	locale := i18n.I18n().Locale()
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
				if ok, complete := listener.OnStart(); ok {
					as.start.Hide()
					as.stop.Show()
					as.swagger.Show()
					as.isEnabled = true
					complete()
				}
			case <-as.stop.ClickedCh:
				if ok, complete := listener.OnStop(); ok {
					as.start.Show()
					as.stop.Hide()
					as.swagger.Hide()
					as.isEnabled = false
					complete()
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
