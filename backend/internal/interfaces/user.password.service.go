package interfaces

import (
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
)

type IUserPasswordService interface {
	interfaces.ICrudService[*entity.UserPassword]
}
