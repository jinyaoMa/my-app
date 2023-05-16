package crud

import "my-app/backend/pkg/database/entity"

type ICrudLog interface {
	ICrud[*entity.Log]
}
