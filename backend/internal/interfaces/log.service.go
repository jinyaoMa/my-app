package interfaces

import (
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
)

type ILogService interface {
	interfaces.ICrudService[*entity.Log]
}
