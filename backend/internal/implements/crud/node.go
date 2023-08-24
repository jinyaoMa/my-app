package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type CRUDNode struct {
	*db.CRUD[*entity.Node]
}

func NewCRUDNode(dbs *db.DB) interfaces.ICRUDNode {
	return &CRUDNode{
		CRUD: db.NewCRUD[*entity.Node](dbs),
	}
}
