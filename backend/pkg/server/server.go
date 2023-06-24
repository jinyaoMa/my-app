package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	options   *Option
	mu        sync.Mutex
	isRunning bool
	hasErrors bool
	errGroup  errgroup.Group
	http      *http.Server // redirector
	https     *fiber.App   // server (tls)
}

// Start implements Interface
func (s *Server) Start(opts *Option) (ok bool) {
	if s.mu.TryLock() {
		defer s.mu.Unlock()
		if !s.isRunning {
			// stopped, can start
			s.options = NewOption(opts)
			return s.start()
		}
	}
	return false
}

// Stop implements Interface
func (s *Server) Stop() (ok bool) {
	if s.mu.TryLock() {
		defer s.mu.Unlock()
		if s.isRunning {
			// running, can stop
			return s.stop()
		}
	}
	return false
}

// HasErrors implements Interface
func (s *Server) HasErrors() bool {
	return s.hasErrors
}

// IsRunning implements Interface
func (s *Server) IsRunning() bool {
	return s.isRunning
}

func (s *Server) start() (ok bool) {
	s.hasErrors = false

	var ln net.Listener
	if ln, ok = s.setup(); !ok {
		return
	}

	s.errGroup.Go(func() error {
		err := s.http.ListenAndServe()
		if err != nil {
			s.hasErrors = true
		}
		return err
	})
	s.errGroup.Go(func() error {
		err := s.https.Listener(ln)
		if err != nil {
			s.hasErrors = true
		}
		return err
	})

	s.isRunning = true
	return true
}

func (s *Server) stop() (ok bool) {
	ctxHttp, cancelHttp := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelHttp()
	if err := s.http.Shutdown(ctxHttp); err != nil && err != http.ErrServerClosed {
		s.options.Logger.Printf("server (http) shutdown error: %+v\n", err)
		s.hasErrors = true
	}

	if err := s.https.Shutdown(); err != nil && err != http.ErrServerClosed {
		s.options.Logger.Printf("server (http/s) shutdown error: %+v\n", err)
		s.hasErrors = true
	}

	if err := s.errGroup.Wait(); err != nil && err != http.ErrServerClosed {
		s.options.Logger.Printf("server running error: %+v\n", err)
		s.hasErrors = true
	}

	s.isRunning = false
	return true
}

func (s *Server) setup() (ln net.Listener, ok bool) {
	addrHttp := fmt.Sprintf(":%d", s.options.Http.Port)
	addrHttps := fmt.Sprintf(":%d", s.options.Https.Port)

	manager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(s.options.Https.HostWhitelist...),
		Cache:      autocert.DirCache(s.options.Https.DirCerts),
	}

	s.http = &http.Server{
		Addr: addrHttp,
		Handler: manager.HTTPHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, addrHttp, addrHttps, 1) + r.RequestURI
			http.Redirect(rw, r, target, http.StatusMovedPermanently)
		})),
	}

	tlsConfig := &tls.Config{
		GetCertificate: s.getSelfSignedOrLetsEncryptCert(manager),
		// By default NextProtos contains the "h2"
		// This has to be removed since Fasthttp does not support HTTP/2
		// Or it will cause a flood of PRI method logs
		// http://webconcepts.info/concepts/http-method/PRI
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}

	ln, err := tls.Listen("tcp", addrHttps, tlsConfig)
	if err != nil {
		return nil, false
	}

	s.https = fiber.New()
	timeFormat := time.RFC3339Nano
	if s.options.IsDev {
		timeFormat = "15:04:05"
	}
	s.https.Use(logger.New(logger.Config{
		Output:        s.options.Logger.Writer(),
		Format:        s.options.Logger.Prefix() + " ${time} | ${status} - ${latency} ${method} ${path}",
		TimeFormat:    timeFormat,
		TimeZone:      "Local",
		DisableColors: !s.options.IsDev,
	}))
	s.options.Setup(s.https)

	return ln, true
}

// getSelfSignedOrLetsEncryptCert override tlsConfig.GetCertificate to enable self-signed certs
func (s *Server) getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		keyFile := filepath.Join(s.options.Https.DirCerts, hello.ServerName+".key")
		crtFile := filepath.Join(s.options.Https.DirCerts, hello.ServerName+".crt")
		certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
		if err != nil {
			// fallback to default cert
			return certManager.GetCertificate(hello)
		}
		return &certificate, err
	}
}

func New() Interface {
	return &Server{}
}
