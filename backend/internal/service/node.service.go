package service

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
)

type NodeService struct {
	interfaces.INodeService
}

func NewNodeService(db *database.Database) interfaces.INodeService {
	return &NodeService{
		INodeService: database.NewCrudService[*entity.Node](db),
	}
}
