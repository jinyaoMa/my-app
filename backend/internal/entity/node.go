package entity

import "my-app/backend/pkg/database/entity"

type Node struct {
	entity.Entity

	/* internal fields */
	Name string `gorm:"size:64; default:''"`
	UUID string `gorm:""`

	/* relational fields */
}
