package interfaces

import "my-app/backend/pkg/database/entity"

type ICrudUserPassword interface {
	ICrud[*entity.UserPassword]
}
