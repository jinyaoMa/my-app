package app

import (
	"log"
	"my-app/backend/model"
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
		log.Println(model.MyOption{})
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
