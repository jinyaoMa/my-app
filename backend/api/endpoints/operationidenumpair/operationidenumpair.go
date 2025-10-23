package operationidenumpair

import (
	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
)

func New(res *apir.Resources) endpoint.Register {
	return new(OperationIdEnumPair).Init(res)
}

type OperationIdEnumPair struct {
	*apir.Resources
}

func (p *OperationIdEnumPair) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		p.RegisterList(api),
		p.RegisterToFlag(api),
		p.RegisterFromFlag(api),
	)
	return
}

func (p *OperationIdEnumPair) Init(res *apir.Resources) *OperationIdEnumPair {
	p.Resources = res
	return p
}
