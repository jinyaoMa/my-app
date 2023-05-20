package interfaces

import "my-app/backend/pkg/database/entity"

type ICrudFileCategory interface {
	ICrud[*entity.FileCategory]
}
