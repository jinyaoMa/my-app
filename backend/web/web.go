package web

import (
	"context"
	"crypto/tls"
	"my-app/backend/app"
	"my-app/backend/web/static"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

var instance *web

type web struct {
	isRunning bool
	errGroup  errgroup.Group
	http      *http.Server // redirector
	https     *http.Server // server (tls)
}

func init() {
	app.App().Env().Log2File(func() {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	})
	gin.DefaultWriter = app.App().Log().Web().Writer()
	instance = &web{}
}

func Web() *web {
	return instance
}

func (w *web) IsRunning() bool {
	return w.isRunning
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
			app.App().Log().Web().Printf("server (http) shutdown error: %+v\n", err)
		}

		ctxHttps, cancelHttps := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelHttps()
		if err := w.https.Shutdown(ctxHttps); err != nil && err != http.ErrServerClosed {
			app.App().Log().Web().Printf("server (http/s) shutdown error: %+v\n", err)
		}

		if err := w.errGroup.Wait(); err != nil && err != http.ErrServerClosed {
			app.App().Log().Web().Printf("server running error: %+v\n", err)
		}

		w.isRunning = false
		return true
	}
	return false
}

func (w *web) reset() {
	cfg := app.App().Config().Web()

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
	var certificate tls.Certificate
	var err error
	dirCerts := ""
	dirCache, ok := certManager.Cache.(autocert.DirCache)
	if ok && string(dirCache) != "" {
		dirCerts = string(dirCache)
	} else {
		certificate, err = tls.X509KeyPair(static.Certs())
	}

	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		if dirCerts != "" {
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
