package crud

import "my-app/backend/pkg/database/entity"

type ICrudFileExtension interface {
	ICrud[*entity.FileExtension]
}
