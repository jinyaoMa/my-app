package web

import (
	"context"
	"crypto/tls"
	"embed"
	"my-app/backend/app"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

//go:embed certs
var certs embed.FS

var (
	once     sync.Once
	instance *web
)

type web struct {
	isRunning bool
	errGroup  errgroup.Group
	http      *http.Server // redirector
	https     *http.Server // server (tls)
}

func Web() *web {
	once.Do(func() {
		instance = &web{}
	})
	return instance
}

func (w *web) Start() (ok bool) {
	if w.isRunning {
		return false
	}

	w.reset()
	w.errGroup.Go(func() error {
		return w.http.ListenAndServe()
	})
	w.errGroup.Go(func() error {
		return w.https.ListenAndServeTLS("", "")
	})

	w.isRunning = true
	return true
}

func (w *web) Stop() (ok bool) {
	if w.isRunning {
		ctxHttp, cancelHttp := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelHttp()
		if err := w.http.Shutdown(ctxHttp); err != nil && err != http.ErrServerClosed {
			app.App().WebLog().Printf("server (http) shutdown error: %+v\n", err)
		}

		ctxHttps, cancelHttps := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelHttps()
		if err := w.https.Shutdown(ctxHttps); err != nil && err != http.ErrServerClosed {
			app.App().WebLog().Printf("server (http/s) shutdown error: %+v\n", err)
		}

		if err := w.errGroup.Wait(); err != nil && err != http.ErrServerClosed {
			app.App().WebLog().Printf("server running error: %+v\n", err)
		}

		w.isRunning = false
		return true
	}
	return false
}

func (w *web) reset() {
	cfg := app.App().WebConfig()

	manager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(cfg.DirCerts),
	}

	w.http = &http.Server{
		Addr: cfg.PortHttp,
		Handler: manager.HTTPHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, cfg.PortHttp, cfg.PortHttps, 1) + r.RequestURI
			http.Redirect(rw, r, target, http.StatusMovedPermanently)
		})),
	}

	tlsConfig := manager.TLSConfig()
	tlsConfig.GetCertificate = w.getSelfSignedOrLetsEncryptCert(manager)

	w.https = &http.Server{
		Addr:      cfg.PortHttps,
		Handler:   router(),
		TLSConfig: tlsConfig,
	}
}

func (s *web) getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	dirCerts := "certs"

	var certificate tls.Certificate
	var err error
	isCustomDirCerts := false
	dirCache, ok := certManager.Cache.(autocert.DirCache)
	if ok && string(dirCache) != "" {
		dirCerts = string(dirCache)
		isCustomDirCerts = true
	} else {
		key, _ := certs.ReadFile(dirCerts + "/localhost.key") // embed use slash as separator
		crt, _ := certs.ReadFile(dirCerts + "/localhost.crt") // embed use slash as separator
		certificate, err = tls.X509KeyPair(crt, key)
	}

	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		if isCustomDirCerts {
			keyFile := filepath.Join(dirCerts, hello.ServerName+".key")
			crtFile := filepath.Join(dirCerts, hello.ServerName+".crt")
			certificate, err = tls.LoadX509KeyPair(crtFile, keyFile)
		}
		if err != nil {
			return certManager.GetCertificate(hello)
		}
		return &certificate, err
	}
}