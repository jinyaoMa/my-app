package interfaces

import (
	"my-app/backend/internal/entity"
	"my-app/backend/pkg/database/interfaces"
)

type INodeService interface {
	interfaces.ICrudService[*entity.Node]
}
