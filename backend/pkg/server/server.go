package server

import (
	"my-app/backend/pkg/server/interfaces"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	options   *Options
	isRunning bool
	errGroup  errgroup.Group
	http      *http.Server // redirector
	https     *http.Server // server (tls)
}

// IsRunning implements interfaces.IServer
func (s *Server) IsRunning() bool {
	return s.isRunning
}

func NewServer(opts *Options) interfaces.IServer {
	opts = NewOptions(opts)

	return &Server{
		options: opts,
	}
}
