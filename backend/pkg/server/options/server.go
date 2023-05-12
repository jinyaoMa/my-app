package options

import (
	"my-app/backend/pkg/logger"
	"my-app/backend/pkg/logger/interfaces"
	"my-app/backend/pkg/logger/options"

	"github.com/imdario/mergo"
)

type OServer struct {
	Logger   interfaces.ILogger
	UseHttps bool
}

func DefaultOServer() *OServer {
	return &OServer{
		Logger:   logger.NewLogger(options.DefaultOLogger()),
		UseHttps: true,
	}
}

func NewOServer(dst *OServer) *OServer {
	src := DefaultOServer()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
