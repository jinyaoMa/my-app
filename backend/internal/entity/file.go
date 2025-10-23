package entity

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type FileStatus string

const (
	FileStatusUploading FileStatus = "uploading"
	FileStatusPersisted FileStatus = "persisted"
	FileStatusLost      FileStatus = "lost"
	FileStatusForbidden FileStatus = "forbidden"
)

type File struct {
	model.Model
	Oid       datatype.Oid  `gorm:"index;not null;size:254;comment:File Oid;"`
	IsDir     sql.NullBool  `gorm:"index;comment:Is Directory or Not;"`
	Name      string        `gorm:"index;not null;check:name <> '';comment:File Name;"`
	Checksums []string      `gorm:"serializer:json;comment:File Checksums;"`
	Size      sql.NullInt64 `gorm:"index;comment:File Size in Bytes;"`

	Status   FileStatus   `gorm:"index;not null;comment:File Status;"`
	Readonly sql.NullBool `gorm:"index;comment:Is Read-only or Not;"`
	Hidden   sql.NullBool `gorm:"index;comment:Is Hidden or Visible;"`

	FileExtensionId *datatype.Id `gorm:"comment:File Extension Id;"`
	FileExtension   *FileExtension

	FileUsers []FileUser
	Users     []User `gorm:"many2many:file_users;"`

	FileGroups []FileGroup
	Groups     []Group `gorm:"many2many:file_groups;"`
}

// AfterCreate implements model.Setup.
func (f *File) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (f *File) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (f *File) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (f *File) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
// Subtle: this method shadows the method (Model).BeforeCreate of File.Model.
func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	if err = f.Model.BeforeCreate(tx); err != nil {
		return
	}
	if len(f.Oid) > 0 {
		var countValid int64
		tx.Model(new(File)).Where([]int64(f.Oid)).Count(&countValid)
		if countValid != int64(len(f.Oid)) {
			return errors.New("file oid contains invalid ids")
		}
	}
	f.Oid = append(f.Oid, f.Id.Int64())
	return
}

// BeforeDelete implements model.Setup.
func (f *File) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// BeforeUpdate implements model.Setup.
func (f *File) BeforeUpdate(tx *gorm.DB) (err error) {
	lenOid := len(f.Oid)
	if lenOid == 0 {
		return errors.New("file oid is empty")
	}
	if f.Oid[lenOid-1] != f.Id.Int64() {
		return errors.New("file oid last element must be self id")
	}

	var countValid int64
	if tx.Statement.Unscoped {
		tx.Unscoped().Model(new(File)).Where([]int64(f.Oid)).Count(&countValid)
	} else {
		tx.Model(new(File)).Where([]int64(f.Oid)).Count(&countValid)
	}
	if countValid != int64(len(f.Oid)) {
		return errors.New("file oid contains invalid ids")
	}
	return
}

// Many2ManySetups implements model.Setup.
func (f *File) Many2ManySetups() (setups []model.Many2ManySetup) {
	return []model.Many2ManySetup{
		{
			Field:     "Users",
			JoinTable: new(FileUser),
		},
		{
			Field:     "Groups",
			JoinTable: new(FileGroup),
		},
	}
}
