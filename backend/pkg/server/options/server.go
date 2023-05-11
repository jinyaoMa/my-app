package options

import "github.com/imdario/mergo"

type OServer struct {
	UseHttps bool
}

func DefaultOServer() *OServer {
	return &OServer{
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
