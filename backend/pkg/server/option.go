package server

import (
	"my-app/backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
)

type Option struct {
	IsDev  bool
	Logger logger.Interface
	Http   OptionHttp
	Https  OptionHttps
	Setup  func(engine *gin.Engine) *gin.Engine
}

type OptionHttp struct {
	Port uint16
}

type OptionHttps struct {
	Port     uint16
	DirCerts string
}

func DefaultOption() *Option {
	return &Option{
		IsDev:  false,
		Logger: logger.New(logger.DefaultOption()),
		Http: OptionHttp{
			Port: 10080,
		},
		Https: OptionHttps{
			Port:     10443,
			DirCerts: "",
		},
		Setup: func(engine *gin.Engine) *gin.Engine {
			return engine
		},
	}
}

func NewOption(dst *Option) *Option {
	src := DefaultOption()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
