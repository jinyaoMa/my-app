package main

import (
	_ "embed"
	// "jinyaoma/node/internal/app"

	"time"

	"majinyao.cn/my-app/backend/internal/app"
)

func main() {
	a := newApp()
	_ = newSystray(a)
	_ = newMainWindow(a)

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			a.Event.Emit("time", now)
			time.Sleep(time.Second)
		}
	}()

	if err := a.Run(); err != nil {
		app.LOG.Fatal(err)
	}
}
