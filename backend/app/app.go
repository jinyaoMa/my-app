package app

import (
	"sync"
)

var (
	once     sync.Once
	instance *app
)

type app struct {
	config *Config
}

func App() *app {
	once.Do(func() {
		instance = &app{
			config: DefaultConfig(),
		}
	})
	return instance
}

func (a *app) Config() *Config {
	return a.config
}

func (a *app) WebConfig() *WebConfig {
	return a.config.Web
}
