package operationidenumpair

import (
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
)

func New(scheme string, tx *gorm.DB) endpoint.Register {
	return new(OperationIdEnumPair).Init(scheme, tx)
}

type OperationIdEnumPair struct {
	Db     *gorm.DB
	Scheme string
}

func (p *OperationIdEnumPair) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops, p.RegisterList(api))
	return
}

func (p *OperationIdEnumPair) Init(scheme string, tx *gorm.DB) *OperationIdEnumPair {
	p.Db = tx
	p.Scheme = scheme
	return p
}
