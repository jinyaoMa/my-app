package wserver

import "my-app/pkg/base"

type IWebServer[TOptions base.IOptions] interface {
	// IsRunning check if the server is running
	IsRunning() bool

	// IsStopping check if the server is stopping
	IsStopping() bool

	// HasErrors check if the server has errors
	HasErrors() bool

	// Start start the server with options
	Start(cfg TOptions) (ok bool)

	// Stop stop the server from running, before() callback before stopping
	Stop(before func()) (ok bool)
}
