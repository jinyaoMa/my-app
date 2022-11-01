package services

import (
	"my-app/backend.new/services/app"
	"my-app/backend.new/services/settings"
	"sync"
)

var (
	instance *services
	once     sync.Once
)

type services struct {
	_all     []interface{}
	app      *app.Service
	settings *settings.Service
}

func Services() *services {
	once.Do(func() {
		app := app.NewService()
		settings := settings.NewService()

		instance = &services{
			_all: []interface{}{
				app,
				settings,
			},
			app:      app,
			settings: settings,
		}
	})
	return instance
}

func (s *services) All() []interface{} {
	return s._all
}

func (s *services) App() *app.Service {
	return s.app
}

func (s *services) Settings() *settings.Service {
	return s.settings
}
