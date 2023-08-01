package interfaces

import (
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
)

type IOptionService interface {
	interfaces.ICrudService[*entity.Option]

	GetByOptionName(name string) (opt *entity.Option, err error)
	GetUint16ByOptionName(name string) (value uint16, opt *entity.Option, err error)
}
