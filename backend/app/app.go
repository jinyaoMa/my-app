package app

import (
	"sync"
)

var (
	once     sync.Once
	instance *app
)

type app struct {
	wlc *WailsLifeCycle
}

func App() *app {
	once.Do(func() {
		instance = &app{
			wlc: &WailsLifeCycle{},
		}
	})
	return instance
}

func (a *app) WailsLifeCycle() *WailsLifeCycle {
	return a.wlc
}
