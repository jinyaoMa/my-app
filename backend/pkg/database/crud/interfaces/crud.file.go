package interfaces

import "my-app/backend/pkg/database/entity"

type ICrudFile interface {
	ICrud[*entity.File]
}
