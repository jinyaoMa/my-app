package entity

import "my-app/backend/pkg/db"

type Node struct {
	db.Entity

	/* internal fields */
	Name string `gorm:"size:64; default:''"`
	UUID string `gorm:""`

	/* relational fields */
}
