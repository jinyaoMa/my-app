package wserver

import (
	"log"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
)

type WebServer struct {
	options    *WebServerOptions
	mutex      sync.Mutex
	isRunning  bool
	isStopping bool
	hasErrors  bool
	errGroup   errgroup.Group
	http       *http.Server // redirector
	https      *fiber.App   // server (tls)
	logger     *log.Logger
}

// HasErrors implements IWebServer.
func (webServer *WebServer) HasErrors() bool {
	return webServer.hasErrors
}

// IsRunning implements IWebServer.
func (webServer *WebServer) IsRunning() bool {
	return webServer.isRunning
}

// IsStopping implements IWebServer.
func (webServer *WebServer) IsStopping() bool {
	return webServer.isStopping
}

// Start implements IWebServer.
func (webServer *WebServer) Start(options *WebServerOptions) (ok bool) {
	if webServer.mutex.TryLock() {
		defer webServer.mutex.Unlock()
		if !webServer.isRunning {
			// stopped, can start
			if options, err := NewWebServerOptions(options); err == nil {
				return webServer.start(options)
			}
		}
	}
	return false
}

// Stop implements IWebServer.
func (webServer *WebServer) Stop(before func()) (ok bool) {
	if webServer.mutex.TryLock() {
		defer webServer.mutex.Unlock()
		if webServer.isRunning {
			// running, can stop
			webServer.isStopping = true
			before()
			ok = webServer.stop()
			webServer.isStopping = false
			return
		}
	}
	return false
}

func NewWebServer(logger *log.Logger) (webServer *WebServer, iWebServer IWebServer) {
	webServer = &WebServer{
		logger: logger,
	}
	return webServer, webServer
}
