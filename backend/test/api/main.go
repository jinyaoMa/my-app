package main

import (
	"my-app/backend/pkg/api"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := api.New()
	if s.Start(&api.Config{}) {
		println("start")
	}

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received

	if s.Stop(func() {
		println("try to stop...")
	}) {
		println("exit")
	}
}
