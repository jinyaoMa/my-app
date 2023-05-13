package options

import (
	"my-app/backend/pkg/logger"
	iLogger "my-app/backend/pkg/logger/interfaces"
	"my-app/backend/pkg/logger/options"

	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
)

type OServer struct {
	IsDev  bool
	Logger iLogger.ILogger
	Http   *OServerHttp
	Https  *OServerHttps
	Setup  func(engine *gin.Engine) *gin.Engine
}

type OServerHttp struct {
	Port uint16
}

type OServerHttps struct {
	Port     uint16
	DirCerts string
}

func DefaultOServer() *OServer {
	return &OServer{
		IsDev:  false,
		Logger: logger.NewLogger(options.DefaultOLogger()),
		Http: &OServerHttp{
			Port: 10080,
		},
		Https: &OServerHttps{
			Port:     10443,
			DirCerts: "",
		},
		Setup: func(engine *gin.Engine) *gin.Engine {
			return engine
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
