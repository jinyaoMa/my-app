package app

import (
	"sync"
)

var (
	once     sync.Once
	instance *Application
)

type Application struct {
	wlc *WailsLifeCycle
}

func App() *Application {
	once.Do(func() {
		instance = &Application{
			wlc: &WailsLifeCycle{},
		}
	})
	return instance
}

func (a *Application) WailsLifeCycle() *WailsLifeCycle {
	return a.wlc
}
