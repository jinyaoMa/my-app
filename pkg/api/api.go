package api

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

type API struct {
	config     *Config
	mu         sync.Mutex
	isRunning  bool
	isStopping bool
	hasErrors  bool
	errGroup   errgroup.Group
	http       *http.Server // redirector
	https      *fiber.App   // server (tls)
}

// Start implements IAPI
func (a *API) Start(cfg *Config) (ok bool) {
	if a.mu.TryLock() {
		defer a.mu.Unlock()
		if !a.isRunning {
			// stopped, can start
			a.config = NewConfig(cfg)
			return a.start()
		}
	}
	return false
}

// Stop implements IAPI
func (a *API) Stop(before func()) (ok bool) {
	if a.mu.TryLock() {
		defer a.mu.Unlock()
		if a.isRunning {
			// running, can stop
			a.isStopping = true
			before()
			ok = a.stop()
			a.isStopping = false
			return
		}
	}
	return false
}

// HasErrors implements IAPI
func (a *API) HasErrors() bool {
	return a.hasErrors
}

// IsRunning implements IAPI
func (a *API) IsRunning() bool {
	return a.isRunning
}

// IsStopping implements IAPI
func (a *API) IsStopping() bool {
	return a.isStopping
}

func (a *API) start() (ok bool) {
	a.hasErrors = false

	var ln net.Listener
	if ln, ok = a.setup(); !ok {
		return
	}

	a.errGroup.Go(func() error {
		err := a.http.ListenAndServe()
		if err != nil {
			a.hasErrors = true
		}
		return err
	})
	a.errGroup.Go(func() error {
		err := a.https.Listener(ln)
		if err != nil {
			a.hasErrors = true
		}
		return err
	})

	a.isRunning = true
	return true
}

func (a *API) stop() (ok bool) {
	ctxHttp, cancelHttp := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelHttp()
	if err := a.http.Shutdown(ctxHttp); err != nil && err != http.ErrServerClosed {
		a.config.Log.Printf("server (http) shutdown error: %+v\n", err)
		a.hasErrors = true
	}

	if err := a.https.Shutdown(); err != nil && err != http.ErrServerClosed {
		a.config.Log.Printf("server (http/s) shutdown error: %+v\n", err)
		a.hasErrors = true
	}

	if err := a.errGroup.Wait(); err != nil && err != http.ErrServerClosed {
		a.config.Log.Printf("server running error: %+v\n", err)
		a.hasErrors = true
	}

	a.isRunning = false
	return true
}

func (a *API) setup() (ln net.Listener, ok bool) {
	addrHttp := fmt.Sprintf(":%d", a.config.Http.Port)
	addrHttps := fmt.Sprintf(":%d", a.config.Https.Port)

	manager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(a.config.Https.HostWhitelist...),
		Cache:      autocert.DirCache(a.config.Https.DirCerts),
	}

	a.http = &http.Server{
		Addr: addrHttp,
		Handler: manager.HTTPHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, addrHttp, addrHttps, 1) + r.RequestURI
			http.Redirect(rw, r, target, http.StatusMovedPermanently)
		})),
	}

	tlsConfig := &tls.Config{
		GetCertificate: a.getSelfSignedOrLetsEncryptCert(manager),
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

	a.https = fiber.New()
	timeFormat := time.RFC3339Nano
	if a.config.IsDev {
		timeFormat = "15:04:05"
	}
	a.https.Use(logger.New(logger.Config{
		Output:        a.config.Log.Writer(),
		Format:        a.config.Log.Prefix() + " ${time} | ${status} - ${latency} ${method} ${path}",
		TimeFormat:    timeFormat,
		TimeZone:      "Local",
		DisableColors: !a.config.IsDev,
	}))
	a.config.Setup(a.https)

	return ln, true
}

// getSelfSignedOrLetsEncryptCert override tlsConfig.GetCertificate to enable self-signed certs
func (a *API) getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		keyFile := filepath.Join(a.config.Https.DirCerts, hello.ServerName+".key")
		crtFile := filepath.Join(a.config.Https.DirCerts, hello.ServerName+".crt")
		certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
		if err != nil {
			// fallback to default cert
			return certManager.GetCertificate(hello)
		}
		return &certificate, err
	}
}

func New() *API {
	return &API{}
}

func NewIAPI() IAPI {
	return &API{}
}
