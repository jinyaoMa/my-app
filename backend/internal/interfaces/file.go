package interfaces

import (
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
)

type IFileService interface {
	interfaces.ICrudService[*entity.File]
}
