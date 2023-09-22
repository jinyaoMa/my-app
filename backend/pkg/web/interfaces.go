package web

type IWeb interface {
	// IsRunning check if the server is running
	IsRunning() bool

	// IsStopping check if the server is stopping
	IsStopping() bool

	// HasErrors check if the server has errors
	HasErrors() bool

	// Start start the server with options
	Start(cfg *Config) (ok bool)

	// Stop stop the server from running, before() callback before stopping
	Stop(before func()) (ok bool)
}
