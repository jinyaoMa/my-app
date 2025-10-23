package endpoints

import (
	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
)

func New(ops []huma.Operation) endpoint.Register {
	return new(endpoints).Init(ops)
}

type endpoints struct {
	Ops    []huma.Operation
}

func (e *endpoints) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		e.RegisterList(api),
		e.RegisterTags(api),
	)
	return
}

func (e *endpoints) Init(ops []huma.Operation) *endpoints {
	e.Ops = ops
	return e
}
