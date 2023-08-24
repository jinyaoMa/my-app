package entity

import (
	"my-app/backend/pkg/db"

	"gorm.io/gorm/schema"
)

type UserFilePermission int

const (
	UserFilePermissionRead UserFilePermission = iota
	UserFilePermissionReadWrite
	UserFilePermissionReadExecute
	UserFilePermissionReadWriteExecute
)

type UserFile struct {
	db.EntityBase
	/* internal fields */
	Permission UserFilePermission `gorm:"default:0"`
	/* relational fields */
	UserID int64 `gorm:"primaryKey"`
	FileID int64 `gorm:"primaryKey"`
}

// TableName implements schema.Tabler
func (*UserFile) TableName() string {
	return "users_files"
}

func NewUserFile() schema.Tabler {
	return new(UserFile)
}
