package interfaces

import (
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
)

type IFileCategoryService interface {
	interfaces.ICrudService[*entity.FileCategory]
}
