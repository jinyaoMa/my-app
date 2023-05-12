package interfaces

import "my-app/backend/pkg/server/options"

type IServer interface {
	// IsRunning check if the server is running
	IsRunning() bool

	// Start start the server with options
	Start(opts *options.OServer) (ok bool)

	// Start stop the server from running
	Stop() (ok bool)
}
