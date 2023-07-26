package interfaces

import (
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
)

type IUserFileService interface {
	interfaces.ICrudService[*entity.UserFile]
}
