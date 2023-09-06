package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type Node struct {
	*db.CRUD[*entity.Node]
}

func NewNode(dbs *db.DB) *Node {
	return &Node{
		CRUD: db.NewCRUD[*entity.Node](dbs),
	}
}

func NewINode(dbs *db.DB) interfaces.ICRUDNode {
	return NewNode(dbs)
}
