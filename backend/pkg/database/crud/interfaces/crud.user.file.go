package interfaces

import "my-app/backend/pkg/database/entity"

type ICrudUserFile interface {
	ICrud[*entity.UserFile]
}
