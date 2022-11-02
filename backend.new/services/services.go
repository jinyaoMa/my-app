package services

import (
	"my-app/backend.new/app"
	"my-app/backend.new/services/general"
	"my-app/backend.new/services/settings"
	"sync"
)

var (
	instance *services
	once     sync.Once
)

type services struct {
	_all     []interface{}
	general  *general.Service
	settings *settings.Service
}

func Services() *services {
	once.Do(func() {
		general := general.NewService()
		settings := settings.NewService()

		instance = &services{
			_all: []interface{}{
				general,
				settings,
			},
			general:  general,
			settings: settings,
		}
		app.App().Log().Services().Println("SERVICES INSTANCE INITIALIZED")
	})
	return instance
}

func (s *services) All() []interface{} {
	return s._all
}

func (s *services) General() *general.Service {
	return s.general
}

func (s *services) Settings() *settings.Service {
	return s.settings
}
