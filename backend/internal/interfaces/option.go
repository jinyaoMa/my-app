package interfaces

import (
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
)

type IOptionService interface {
	interfaces.ICrudService[*entity.Option]
}
