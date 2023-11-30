package wserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"golang.org/x/crypto/acme/autocert"
)

func (webServer *WebServer) start(options *WebServerOptions) (ok bool) {
	webServer.options = options
	webServer.hasErrors = false

	var ln net.Listener
	if ln, ok = webServer.setup(); !ok {
		return
	}

	webServer.errGroup.Go(func() error {
		err := webServer.http.ListenAndServe()
		if err != nil {
			webServer.hasErrors = true
		}
		return err
	})
	webServer.errGroup.Go(func() error {
		err := webServer.https.Listener(ln)
		if err != nil {
			webServer.hasErrors = true
		}
		return err
	})

	webServer.isRunning = true
	return true
}

func (webServer *WebServer) stop() (ok bool) {
	ctxHttp, cancelHttp := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelHttp()
	if err := webServer.http.Shutdown(ctxHttp); err != nil && err != http.ErrServerClosed {
		webServer.logger.Printf("server (http) shutdown error: %+v\n", err)
		webServer.hasErrors = true
	}

	if err := webServer.https.Shutdown(); err != nil && err != http.ErrServerClosed {
		webServer.logger.Printf("server (http/s) shutdown error: %+v\n", err)
		webServer.hasErrors = true
	}

	if err := webServer.errGroup.Wait(); err != nil && err != http.ErrServerClosed {
		webServer.logger.Printf("server running error: %+v\n", err)
		webServer.hasErrors = true
	}

	webServer.isRunning = false
	return true
}

func (webServer *WebServer) setup() (ln net.Listener, ok bool) {
	addrHttp := fmt.Sprintf(":%d", webServer.options.Http.Port)
	addrHttps := fmt.Sprintf(":%d", webServer.options.Https.Port)

	manager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(webServer.options.Https.HostWhitelist...),
		Cache:      autocert.DirCache(webServer.options.Https.DirCerts),
	}

	webServer.http = &http.Server{
		Addr: addrHttp,
		Handler: manager.HTTPHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, addrHttp, addrHttps, 1) + r.RequestURI
			http.Redirect(rw, r, target, http.StatusMovedPermanently)
		})),
	}

	tlsConfig := &tls.Config{
		GetCertificate: webServer.getSelfSignedOrLetsEncryptCert(manager),
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

	webServer.https = fiber.New()
	timeFormat := time.RFC3339Nano
	if webServer.options.IsDev {
		timeFormat = "15:04:05"
	}
	webServer.https.Use(logger.New(logger.Config{
		Output:        webServer.logger.Writer(),
		Format:        webServer.logger.Prefix() + " ${time} | ${status} - ${latency} ${method} ${path}",
		TimeFormat:    timeFormat,
		TimeZone:      "Local",
		DisableColors: !webServer.options.IsDev,
	}))
	webServer.options.Setup(webServer.https)

	return ln, true
}

// getSelfSignedOrLetsEncryptCert override tlsConfig.GetCertificate to enable self-signed certs
func (webServer *WebServer) getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		keyFile := filepath.Join(webServer.options.Https.DirCerts, hello.ServerName+".key")
		crtFile := filepath.Join(webServer.options.Https.DirCerts, hello.ServerName+".crt")
		certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
		if err != nil {
			// fallback to default cert
			return certManager.GetCertificate(hello)
		}
		return &certificate, err
	}
}
