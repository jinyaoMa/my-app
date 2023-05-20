package interfaces

import "my-app/backend/pkg/database/entity"

type ICrudUser interface {
	ICrud[*entity.User]
}
