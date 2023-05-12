package options

import (
	"my-app/backend/pkg/logger"
	"my-app/backend/pkg/logger/interfaces"
	"my-app/backend/pkg/logger/options"

	"github.com/imdario/mergo"
)

type OServer struct {
	IsDev  bool
	Logger interfaces.ILogger
	Ports  *OServerPorts
}

type OServerPorts struct {
	Http  uint16
	Https uint16
}

func DefaultOServer() *OServer {
	return &OServer{
		IsDev:  false,
		Logger: logger.NewLogger(options.DefaultOLogger()),
		Ports: &OServerPorts{
			Http:  8091,
			Https: 0,
		},
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
