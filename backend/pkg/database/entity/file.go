package entity

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

var (
	regexpChecksum = regexp.MustCompile(`^[0-9a-f]{32}\:[0-9a-f]{128}\:\d$`)
)

type File struct {
	Entity

	/* internal fields */
	IsDirectory bool      `gorm:"default:false"`
	Path        string    `gorm:"size:4096"`
	Name        string    `gorm:"size:1024"`
	Size        int64     `gorm:"default:0"`
	ReadOnly    bool      `gorm:"default:false"`
	Hidden      bool      `gorm:"default:false"`
	VisitedAt   time.Time `gorm:""`

	// md5:sha512:File.Size => 32 + 128 = 160 hex digits + 2[:] + 20[int64] = 182 (size)
	// for file, md5 and sha512 are hashed using file's data
	// for directory, md5 and sha512 are hashed using File.Path+File.Name
	Checksum string `gorm:"size:182; unique; index"`

	/* relational fields */
	UserID          int64   `gorm:""`
	AccessableUsers []*User `gorm:"many2many:users_files"`
	FileExtensionID int64   `gorm:""`
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
		if err = f.validateIfExist(tx); err != nil {
			return
		}
		if tx.Statement.Changed("Name") {
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

func (f *File) AfterFind(tx *gorm.DB) (err error) {
	if err = f.Entity.AfterFind(tx); err != nil {
		return
	}

	if f != nil {
		tx.Statement.UpdateColumn("VisitedAt", time.Now())
	}
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
		err = errors.New("File.Checksum format should be `md5:sha512:File.Size`")
	}
	return
}
