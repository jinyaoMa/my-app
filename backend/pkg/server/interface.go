package server

type Interface interface {
	// IsRunning check if the server is running
	IsRunning() bool

	// HasErrors check if the server has errors
	HasErrors() bool

	// Start start the server with options
	Start(opts *Option) (ok bool)

	// Start stop the server from running
	Stop(stopping func(), stopped func(hasError bool)) (ok bool)
}
