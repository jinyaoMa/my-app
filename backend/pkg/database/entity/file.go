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
	IsDirectory bool   `gorm:"default:false"`
	Path        string `gorm:"size:4096"`
	Name        string `gorm:"size:512"`
	Size        int64  `gorm:"default:0"`
	Extension   FileExtension
	VisitedAt   time.Time

	// md5:sha512:File.Size => 32 + 128 = 160 hex digits + 2[:] + 20[int64] = 182 (size)
	// for file, md5 and sha512 are hashed using file's data
	// for directory, md5 and sha512 are hashed using File.Path+File.Name
	Checksum string `gorm:"size:182; unique; index"`
}

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	if err = f.Entity.BeforeCreate(tx); err != nil {
		return
	}

	if f != nil {
		f.validateName(tx)
	}
	return
}

func (f *File) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = f.Entity.BeforeUpdate(tx); err != nil {
		return
	}

	if f != nil {
		if tx.Statement.Changed("Name") {
			f.validateName(tx)
		}
		if tx.Statement.Changed("Checksum") {
			f.validateChecksum(tx)
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
