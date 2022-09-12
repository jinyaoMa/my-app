package app

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	instance *Application
)

type Application struct {
}

func App() *Application {
	once.Do(func() {
		fmt.Println("Creating single instance now.")
		instance = &Application{}
	})
	fmt.Println("Single instance already created.")
	return instance
}
