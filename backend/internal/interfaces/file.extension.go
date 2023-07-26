package interfaces

import (
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
)

type IFileExtensionService interface {
	interfaces.ICrudService[*entity.FileExtension]
}
