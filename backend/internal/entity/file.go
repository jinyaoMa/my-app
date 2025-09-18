package entity

import (
	"errors"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type FileStatus int

const (
	FileStatusUploading FileStatus = iota + 1
	FileStatusPersisted
	FileStatusNotFound
	FileStatusForbidden
)

type File struct {
	db.Entity
	Oid       datatype.Oid `gorm:"index;not null;size:254;comment:File Oid;"`
	IsDir     bool         `gorm:"index;comment:Is Directory or Not;"`
	Name      string       `gorm:"index;comment:File Name;"`
	Checksums []string     `gorm:"serializer:json;comment:File Checksums;"`
	Size      uint64       `gorm:"comment:File Size in Bytes;"`

	Status   FileStatus `gorm:"index;comment:File Status;"`
	Readonly bool       `gorm:"comment:Is Read-only or Not;"`
	Hidden   bool       `gorm:"comment:Is Hidden or Visible;"`

	FileExtensionId *int64 `gorm:"comment:File Extension Id;"`
	FileExtension   *FileExtension

	FileUsers []FileUser
	Users     []User `gorm:"many2many:file_users;"`

	FileGroups []FileGroup
	Groups     []Group `gorm:"many2many:file_groups;"`
}

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	var countValid int64
	tx.Model(new(File)).Where(f.Oid).Count(&countValid)
	if countValid != int64(len(f.Oid)) {
		return errors.New("file oid contains invalid ids")
	}
	return nil
}

func (f *File) AfterCreate(tx *gorm.DB) (err error) {
	res := tx.Model(f).Update("oid", append(f.Oid, f.Id))
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (f *File) BeforeUpdate(tx *gorm.DB) (err error) {
	var countValid int64
	tx.Model(new(File)).Where(f.Oid).Count(&countValid)
	if countValid != int64(len(f.Oid)) {
		return errors.New("file oid contains invalid ids")
	}
	return nil
}

func (f *File) GetEntityM2MSetups() []db.EntityM2MSetup {
	return []db.EntityM2MSetup{
		{
			Model:     new(File),
			Field:     "Users",
			JoinTable: new(FileUser),
		},
		{
			Model:     new(File),
			Field:     "Groups",
			JoinTable: new(FileGroup),
		},
	}
}
