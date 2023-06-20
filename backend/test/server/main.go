package main

import (
	"my-app/backend/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := server.New()
	s.Start(&server.Option{})

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received

	s.Stop()
}
