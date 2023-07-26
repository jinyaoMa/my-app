package interfaces

import (
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
)

type IUserService interface {
	interfaces.ICrudService[*entity.User]
}
