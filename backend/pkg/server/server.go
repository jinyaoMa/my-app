package server

import (
	"my-app/backend/pkg/server/interfaces"
	"my-app/backend/pkg/server/options"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	options   *options.OServer
	isRunning bool
	errGroup  errgroup.Group
	http      *http.Server // redirector
	https     *http.Server // server (tls)
}

// Start implements interfaces.IServer
func (*Server) Start(opts options.OServer) (ok bool) {
	panic("unimplemented")
}

// Stop implements interfaces.IServer
func (*Server) Stop() (ok bool) {
	panic("unimplemented")
}

// IsRunning implements interfaces.IServer
func (s *Server) IsRunning() bool {
	return s.isRunning
}

func NewServer(opts *options.OServer) interfaces.IServer {
	opts = options.NewOServer(opts)

	return &Server{
		options: opts,
	}
}
