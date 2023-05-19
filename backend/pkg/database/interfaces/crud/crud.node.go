package crud

import "my-app/backend/pkg/database/entity"

type ICrudNode interface {
	ICrud[*entity.Node]
}
