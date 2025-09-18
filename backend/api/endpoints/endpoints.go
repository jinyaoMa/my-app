package endpoints

import (
	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
)

func New(scheme string, ops []huma.Operation) endpoint.Register {
	return new(endpoints).init(scheme, ops)
}

type endpoints struct {
	Scheme string
	Ops    []huma.Operation
}

func (e *endpoints) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		e.RegisterList(api),
		e.RegisterTags(api),
	)
	return
}

func (e *endpoints) init(scheme string, ops []huma.Operation) *endpoints {
	e.Scheme = scheme
	e.Ops = ops
	return e
}
