package server

type Interface interface {
	// IsRunning check if the server is running
	IsRunning() bool

	// HasErrors check if the server has errors
	HasErrors() bool

	// Start start the server with options
	Start(opts *Option) (ok bool)

	// Stop stop the server from running, before() callback before stopping
	Stop(before func()) (ok bool)
}
