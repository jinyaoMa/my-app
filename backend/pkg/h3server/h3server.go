package h3server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/quic-go/quic-go/http3"
	"golang.org/x/sync/errgroup"
)

type IH3Server interface {
	IsRunning() bool
	Addr() string
	SecureAddr() string
	OnRun(cb func(h1 *http.Server, h2 *http.Server, h3 *http3.Server))
	OnShutddown(cb func(h1 *http.Server, h2 *http.Server, h3 *http3.Server))
	Run(port uint16, securePort uint16, certFile string, keyFile string, onStart ...func()) (err error)
	Shutdown(ctx context.Context) (err error)
}

func New(router http.Handler) IH3Server {
	return new(h3server).init(router)
}

type h3server struct {
	isRunning   atomic.Bool
	addr        string
	secureAddr  string
	router      http.Handler
	http        *http.Server
	https       *http.Server
	http3       *http3.Server
	onRuns      []func(h1 *http.Server, h2 *http.Server, h3 *http3.Server)
	onShutdowns []func(h1 *http.Server, h2 *http.Server, h3 *http3.Server)
}

func (s *h3server) Addr() string {
	return s.addr
}

func (s *h3server) SecureAddr() string {
	return s.secureAddr
}

func (s *h3server) IsRunning() bool {
	return s.isRunning.Load()
}

func (s *h3server) OnRun(cb func(h1 *http.Server, h2 *http.Server, h3 *http3.Server)) {
	if cb != nil {
		s.onRuns = append(s.onRuns, cb)
	}
}
func (s *h3server) OnShutddown(cb func(h1 *http.Server, h2 *http.Server, h3 *http3.Server)) {
	if cb != nil {
		s.onShutdowns = append(s.onShutdowns, cb)
	}
}

func (s *h3server) Run(port uint16, securePort uint16, certFile string, keyFile string, onStart ...func()) (err error) {
	if s.isRunning.Swap(true) {
		return errors.New("server is already running")
	}
	for _, f := range onStart {
		if f != nil {
			f()
		}
	}
	defer func() {
		s.isRunning.Store(false)
		for _, onShutdown := range s.onShutdowns {
			onShutdown(s.http, s.https, s.http3)
		}
	}()

	s.addr = fmt.Sprintf(":%d", port)
	s.secureAddr = fmt.Sprintf(":%d", securePort)
	router := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		s.http3.SetQUICHeaders(rw.Header())
		s.router.ServeHTTP(rw, r)
	})

	s.http = &http.Server{
		Addr: s.addr,
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, s.addr, s.secureAddr, 1) + r.RequestURI
			http.Redirect(rw, r, target, http.StatusMovedPermanently)
		}),
	}
	s.https = &http.Server{
		Addr:    s.secureAddr,
		Handler: router,
	}
	s.http3 = &http3.Server{
		Addr:    s.secureAddr,
		Handler: router,
	}

	var g errgroup.Group
	g.Go(func() error {
		errServe := s.http.ListenAndServe()
		if !errors.Is(errServe, http.ErrServerClosed) {
			s.https.Close()
			s.http3.Close()
		}
		return errServe
	})
	g.Go(func() error {
		errServe := s.https.ListenAndServeTLS(certFile, keyFile)
		if !errors.Is(errServe, http.ErrServerClosed) {
			s.http.Close()
			s.http3.Close()
		}
		return errServe
	})
	g.Go(func() error {
		errServe := s.http3.ListenAndServeTLS(certFile, keyFile)
		if !errors.Is(errServe, http.ErrServerClosed) {
			s.http.Close()
			s.https.Close()
		}
		return errServe
	})
	for _, onRun := range s.onRuns {
		onRun(s.http, s.https, s.http3)
	}
	err = g.Wait()
	return
}

func (s *h3server) Shutdown(ctx context.Context) (err error) {
	if !s.isRunning.Load() {
		return errors.New("server is not running")
	}

	// Shutdown signal with grace period of 30 seconds
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
	}

	// Trigger graceful shutdown
	var g errgroup.Group
	g.Go(func() error {
		return s.http.Shutdown(ctx)
	})
	g.Go(func() error {
		return s.https.Shutdown(ctx)
	})
	g.Go(func() error {
		return s.http3.Shutdown(ctx)
	})
	return g.Wait()
}

func (s *h3server) init(router http.Handler) *h3server {
	s.router = router
	return s
}
