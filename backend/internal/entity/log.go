package entity

import "my-app/backend/pkg/db"

type Log struct {
	db.Entity

	/* internal fields */
	Message string `gorm:"size:4096; default:''"`

	/* relational fields */
}
