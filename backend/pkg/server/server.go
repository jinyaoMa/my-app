package server

import (
	"context"
	"fmt"
	"my-app/backend/pkg/server/interfaces"
	"my-app/backend/pkg/server/options"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	*options.OServer
	mu        sync.Mutex
	isRunning bool
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
			s.OServer = opts
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

// IsRunning implements interfaces.IServer
func (s *Server) IsRunning() bool {
	return s.isRunning
}

func (s *Server) start() (ok bool) {
	engine := gin.New()

	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			var statusColor, methodColor, resetColor string
			if param.IsOutputColor() {
				statusColor = param.StatusCodeColor()
				methodColor = param.MethodColor()
				resetColor = param.ResetColor()
			}
			if param.Latency > time.Minute {
				param.Latency = param.Latency.Truncate(time.Second)
			}
			return fmt.Sprintf("%s %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
				s.OServer.Logger.Prefix(),
				param.TimeStamp.Format("2006/01/02 - 15:04:05"),
				statusColor, param.StatusCode, resetColor,
				param.Latency,
				param.ClientIP,
				methodColor, param.Method, resetColor,
				param.Path,
				param.ErrorMessage,
			)
		},
		Output: s.OServer.Logger.Writer(),
	}))

	return false
}

func (s *Server) stop() (ok bool) {
	ctxHttp, cancelHttp := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelHttp()
	if err := s.http.Shutdown(ctxHttp); err != nil && err != http.ErrServerClosed {
		s.OServer.Logger.Printf("server (http) shutdown error: %+v\n", err)
	}

	ctxHttps, cancelHttps := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelHttps()
	if err := s.https.Shutdown(ctxHttps); err != nil && err != http.ErrServerClosed {
		s.OServer.Logger.Printf("server (http/s) shutdown error: %+v\n", err)
	}

	if err := s.errGroup.Wait(); err != nil && err != http.ErrServerClosed {
		s.OServer.Logger.Printf("server running error: %+v\n", err)
	}

	return true
}

func NewServer() interfaces.IServer {
	return &Server{}
}
