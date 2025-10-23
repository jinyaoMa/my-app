package entity

import (
	"database/sql"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type FileUser struct {
	model.Model

	FileId datatype.Id `gorm:"comment:File Id;"`
	File   File

	UserId datatype.Id `gorm:"uniqueIndex:idx_file_users_user_avatar;comment:User Id;"`
	User   User

	NoCreate sql.NullBool `gorm:"comment:No Create;"`
	NoRead   sql.NullBool `gorm:"comment:No Read;"`
	NoUpdate sql.NullBool `gorm:"comment:No Update;"`
	NoDelete sql.NullBool `gorm:"comment:No Delete;"`
	IsAvatar sql.NullBool `gorm:"uniqueIndex:idx_file_users_user_avatar;comment:Is Used as Avatar or Not;"`
}

// AfterCreate implements model.Setup.
func (u *FileUser) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (u *FileUser) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (u *FileUser) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (u *FileUser) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
// Subtle: this method shadows the method (Model).BeforeCreate of FileUser.Model.
func (u *FileUser) BeforeCreate(tx *gorm.DB) (err error) {
	return u.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
func (u *FileUser) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// BeforeUpdate implements model.Setup.
func (u *FileUser) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

// Many2ManySetups implements model.Setup.
func (u *FileUser) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
}
