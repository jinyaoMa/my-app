package crud

import "my-app/backend/pkg/database/entity"

type ICrudOption interface {
	ICrud[*entity.Option]
}