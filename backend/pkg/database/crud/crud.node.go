package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/crud/interfaces"
	"my-app/backend/pkg/database/entity"
)

type CrudNode struct {
	interfaces.ICrud[*entity.Node]
}

func NewCrudNode(database *database.Database) interfaces.ICrudNode {
	return &CrudNode{
		ICrud: NewCrud[*entity.Node](database),
	}
}
