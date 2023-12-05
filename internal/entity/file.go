package entity

import (
	"fmt"
	"my-app/pkg/db"
	"path/filepath"
	"time"
)

type File struct {
	db.Entity[*File]
	IsDirectory bool      `gorm:"default:false"`
	Path        string    `gorm:"index"` // the directory that this file/folder placed in
	Name        string    `gorm:"index"` // file/folder name
	Size        uint64    `gorm:"default:0"`
	Hidden      bool      `gorm:"default:false"`
	VisitedAt   time.Time `gorm:""`

	// checksum format `{sha1:160bit:40hex}{xxh3:128bit:32hex}{size:64bit:16hex}`
	// 40+32+16=88
	Checksum string `gorm:"size:88; index"`

	/* belongs to */
	UserID          int64         `gorm:""`
	User            User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FileExtensionID int64         `gorm:""`
	FileExtension   FileExtension `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (file *File) UPath() string {
	return filepath.Join("/", fmt.Sprint(file.UserID), file.Path, file.Name)
}
