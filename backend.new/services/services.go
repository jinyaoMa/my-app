package services

import (
	"my-app/backend.new/services/settings"
	"sync"
)

var (
	instance *services
	once     sync.Once
)

type services struct {
	_all     []interface{}
	settings *settings.Service
}

func Services() *services {
	once.Do(func() {
		settings := settings.NewService()

		instance = &services{
			_all: []interface{}{
				settings,
			},
			settings: settings,
		}
	})
	return instance
}

func (s *services) All() []interface{} {
	return s._all
}

func (s *services) Settings() *settings.Service {
	return s.settings
}
