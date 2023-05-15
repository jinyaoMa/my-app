package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"my-app/backend/pkg/server/interfaces"
	"my-app/backend/pkg/server/options"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	options   *options.OServer
	mu        sync.Mutex
	isRunning bool
	hasErrors bool
	errGroup  errgroup.Group
	http      *http.Server // redirector
	https     *http.Server // server (tls)
}

// Start implements interfaces.IServer
func (s *Server) Start(opts *options.OServer) (ok bool) {
	if s.mu.TryLock() {
		defer s.mu.Unlock()
		if !s.isRunning {
			// stopped, can start
			s.options = opts
			return s.start()
		}
	}
	return false
}

// Stop implements interfaces.IServer
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

// HasErrors implements interfaces.IServer
func (s *Server) HasErrors() bool {
	return s.hasErrors
}

// IsRunning implements interfaces.IServer
func (s *Server) IsRunning() bool {
	return s.isRunning
}

func NewServer() interfaces.IServer {
	return &Server{}
}

func (s *Server) start() (ok bool) {
	s.hasErrors = false

	engine := gin.New()

	engine.UseH2C = true
	engine.Use(gin.Recovery())
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			var statusColor, methodColor, resetColor string
			if param.IsOutputColor() && s.options.IsDev {
				statusColor = param.StatusCodeColor()
				methodColor = param.MethodColor()
				resetColor = param.ResetColor()
			}
			if param.Latency > time.Minute {
				param.Latency = param.Latency.Truncate(time.Second)
			}
			return fmt.Sprintf("%s %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
				s.options.Logger.Prefix(),
				param.TimeStamp.Format("2006/01/02 - 15:04:05"),
				statusColor, param.StatusCode, resetColor,
				param.Latency,
				param.ClientIP,
				methodColor, param.Method, resetColor,
				param.Path,
				param.ErrorMessage,
			)
		},
		Output: s.options.Logger.Writer(),
	}))

	s.setup(engine)
	s.errGroup.Go(func() error {
		err := s.http.ListenAndServe()
		if err != nil {
			s.hasErrors = true
		}
		return err
	})
	s.errGroup.Go(func() error {
		err := s.https.ListenAndServeTLS("", "")
		if err != nil {
			s.hasErrors = true
		}
		return err
	})

	s.isRunning = true
	return true
}

func (s *Server) stop() (ok bool) {
	ctxHttp, cancelHttp := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelHttp()
	if err := s.http.Shutdown(ctxHttp); err != nil && err != http.ErrServerClosed {
		s.options.Logger.Printf("server (http) shutdown error: %+v\n", err)
		s.hasErrors = true
	}

	ctxHttps, cancelHttps := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelHttps()
	if err := s.https.Shutdown(ctxHttps); err != nil && err != http.ErrServerClosed {
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

func (s *Server) setup(engine *gin.Engine) {
	addrHttp := fmt.Sprintf(":%d", s.options.Http.Port)
	addrHttps := fmt.Sprintf(":%d", s.options.Https.Port)

	manager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(s.options.Https.DirCerts),
	}

	s.http = &http.Server{
		Addr: addrHttp,
		Handler: manager.HTTPHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, addrHttp, addrHttps, 1) + r.RequestURI
			http.Redirect(rw, r, target, http.StatusMovedPermanently)
		})),
	}

	tlsConfig := manager.TLSConfig()
	tlsConfig.GetCertificate = s.getSelfSignedOrLetsEncryptCert(manager)

	s.https = &http.Server{
		Addr:      addrHttps,
		Handler:   s.options.Setup(engine).Handler(),
		TLSConfig: tlsConfig,
	}
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
