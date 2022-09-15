package web

import (
	"context"
	"crypto/tls"
	"embed"
	"log"
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
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := w.http.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
			log.Printf("server (http) shutdown error: %+v\n", err)
		}
		if err := w.https.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
			log.Printf("server (http/s) shutdown error: %+v\n", err)
		}
		if err := w.errGroup.Wait(); err != nil && err != http.ErrServerClosed {
			log.Printf("server running error: %+v\n", err)
		}

		w.isRunning = false
		return true
	}
	return false
}

func (w *web) reset() {
	portHttp := ":10080"
	portHttps := ":10443"

	manager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
	}

	w.http = &http.Server{
		Addr: portHttp,
		Handler: manager.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, portHttp, portHttps, 1) + r.RequestURI
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		})),
	}

	tlsConfig := manager.TLSConfig()
	tlsConfig.GetCertificate = w.getSelfSignedOrLetsEncryptCert(manager)

	w.https = &http.Server{
		Addr:      portHttps,
		Handler:   router(),
		TLSConfig: tlsConfig,
	}
}

func (s *web) getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		var certificate tls.Certificate
		var err error
		dirCache, ok := certManager.Cache.(autocert.DirCache)
		if ok {
			keyFile := filepath.Join(string(dirCache), hello.ServerName+".key")
			crtFile := filepath.Join(string(dirCache), hello.ServerName+".crt")
			certificate, err = tls.LoadX509KeyPair(crtFile, keyFile)
		} else {
			key, _ := certs.ReadFile("certs/localhost.key")
			crt, _ := certs.ReadFile("certs/localhost.crt")
			certificate, err = tls.X509KeyPair(crt, key)
		}
		if err != nil {
			return certManager.GetCertificate(hello)
		}
		return &certificate, err
	}
}
