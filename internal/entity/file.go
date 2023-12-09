package entity

import (
	"errors"
	"fmt"
	"my-app/pkg/db"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

type File struct {
	db.Entity[*File]
	IsDir     bool      `gorm:"default:false"`
	Checksum  string    `gorm:"index"` // checksum format `{sha1:160bit:40hex}{xxh3:128bit:32hex}{size:64bit:16hex}`
	Path      string    `gorm:"index"` // the directory that this file/folder placed in; directory should end with '/'
	Name      string    `gorm:"index"` // file/folder name
	Size      uint64    `gorm:"default:0"`
	Hidden    bool      `gorm:"default:false"`
	VisitedAt time.Time `gorm:""`

	/* belongs to */
	UserID          uint64        `gorm:""`
	User            User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FileExtensionID uint64        `gorm:""`
	FileExtension   FileExtension `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (file *File) BeforeCreate(tx *gorm.DB) (err error) {
	if !strings.HasSuffix(file.Path, "/") {
		e := fmt.Sprintf("file path %s should end with '/'", file.Path)
		return errors.New(e)
	}
	return file.Entity.BeforeCreate(tx)
}

func (file *File) BeforeUpdate(tx *gorm.DB) (err error) {
	if !strings.HasSuffix(file.Path, "/") {
		e := fmt.Sprintf("file path %s should end with '/'", file.Path)
		return errors.New(e)
	}
	return file.Entity.BeforeUpdate(tx)
}

func (file *File) UPath() string {
	return filepath.Join("/", fmt.Sprint(file.UserID), file.Path, file.Name)
}
