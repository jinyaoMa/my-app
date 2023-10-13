package entity

import (
	"errors"
	"my-app/backend/pkg/db"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

var (
	//md5+sha256+crc32+size =>(hex encoded) 128bit/4+256bit/4+32bit/4+64bit/4=32+64+8+16=123bytes
	regexpChecksum = regexp.MustCompile(`^[0-9a-f]{32}\+[0-9a-f]{64}\+[0-9a-f]{8}\+[0-9a-f]{1,16}$`)
)

type File struct {
	db.Entity

	/* internal fields */
	IsDirectory bool      `gorm:"default:false"`
	Path        string    `gorm:"index"`
	Name        string    `gorm:""`
	Size        uint64    `gorm:"default:0"`
	ReadOnly    bool      `gorm:"default:false"`
	Hidden      bool      `gorm:"default:false"`
	VisitedAt   time.Time `gorm:""`

	// md5+sha256+crc32+size =>(hex encoded) 128bit/4+256bit/4+32bit/4+64bit/4=32+64+8+16=123bytes
	// for file, md5, sha256 and crc32 are hashed using file's data
	Checksum string `gorm:"size:123; index"`

	/* relational fields */
	UserID          *int64 `gorm:""`
	FileExtensionID *int64 `gorm:""`
}

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	if err = f.Entity.BeforeCreate(tx); err != nil {
		return
	}

	if f != nil {
		if err = f.validateIfExist(tx); err != nil {
			return
		}
		if err = f.validateName(tx); err != nil {
			return
		}
	}
	return
}

func (f *File) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = f.Entity.BeforeUpdate(tx); err != nil {
		return
	}

	if f != nil {
		if !tx.Statement.Changed("ReadOnly") && f.ReadOnly {
			return errors.New("File is set to ReadOnly")
		}
		if tx.Statement.Changed("Path") || tx.Statement.Changed("Name") {
			if err = f.validateIfExist(tx); err != nil {
				return
			}
			if err = f.validateName(tx); err != nil {
				return
			}
		}
		if tx.Statement.Changed("Checksum") {
			if err = f.validateChecksum(tx); err != nil {
				return
			}
		}
	}
	return
}

func (f *File) BeforeDelete(tx *gorm.DB) (err error) {
	if err = f.Entity.BeforeDelete(tx); err != nil {
		return
	}

	if f != nil {
		var readonly bool
		readonly, err = f.validateIfReadOnly(tx)
		if err != nil {
			return
		}
		if readonly {
			return errors.New("File is set to ReadOnly")
		}
	}
	return
}

func (f *File) AfterFind(tx *gorm.DB) (err error) {
	if err = f.Entity.AfterFind(tx); err != nil {
		return
	}

	if f != nil {
		tx.Statement.UpdateColumn("VisitedAt", time.Now())
	}
	return
}

func (f *File) validateIfReadOnly(tx *gorm.DB) (readonly bool, err error) {
	var file *File
	result := tx.First(file, f.ID)
	if result.Error != nil {
		err = errors.New("File not exists")
	}
	readonly = file.ReadOnly
	return
}

func (f *File) validateIfExist(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&File{}).Where(&File{
		Path: f.Path,
		Name: f.Name,
	}).Count(&count)
	if count > 0 {
		err = errors.New("File Path and Name exists")
	}
	return
}

func (f *File) validateName(tx *gorm.DB) (err error) {
	f.Name = strings.TrimSpace(f.Name)
	if f.Name == "" {
		err = errors.New("File.Name cannot be empty or spaces")
	} else if len(f.Name) > 256 {
		err = errors.New("File.Name length cannot be more than 512")
	}
	return
}

func (f *File) validateChecksum(tx *gorm.DB) (err error) {
	f.Checksum = strings.TrimSpace(f.Checksum)
	f.Checksum = strings.ToLower(f.Checksum)
	if len(f.Checksum) < 182 && regexpChecksum.MatchString(f.Checksum) {
		err = errors.New("File.Checksum format should be `md5+sha256+crc32+size` hex encoded")
	}
	return
}
