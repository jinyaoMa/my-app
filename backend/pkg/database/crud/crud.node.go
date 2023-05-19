package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces/crud"
)

type CrudNode struct {
	crud.ICrud[*entity.Node]
}

func NewCrudNode(database *database.Database) crud.ICrudNode {
	return &CrudNode{
		ICrud: NewCrud[*entity.Node](database),
	}
}
