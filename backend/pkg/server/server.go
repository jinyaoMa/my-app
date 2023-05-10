package server

import (
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

func NewServer(opts *Options) *Server {
	opts = NewOptions(opts)

	return &Server{
		options: opts,
	}
}
