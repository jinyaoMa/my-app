package services

import (
	"my-app/backend.new/services/general"
	"my-app/backend.new/services/settings"
	"sync"
)

var _services = &services{}

type services struct {
	once     sync.Once
	all      []interface{}
	general  *general.Service
	settings *settings.Service
}

// entry of all services
func Services() *services {
	_services.once.Do(func() {
		_services.general = general.NewService()
		_services.settings = settings.NewService()

		_services.all = []interface{}{
			_services.general,
			_services.settings,
		}
	})
	return _services
}

func (ss *services) All() []interface{} {
	return ss.all
}

func (ss *services) General() *general.Service {
	return ss.general
}

func (ss *services) Settings() *settings.Service {
	return ss.settings
}
