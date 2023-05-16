package entity

import (
	iSnowflake "my-app/backend/pkg/snowflake/interfaces"

	"gorm.io/gorm"
)

type Option struct {
	Entity
	Key       string `gorm:"size:100; unique; index"`
	Value     string `gorm:"size:255"`
	Encrypted bool   `gorm:""`
}

func (o *Option) BeforeCreate(tx *gorm.DB) (err error) {
	if err = o.Entity.EntityBase.BeforeCreate(tx); err != nil {
		return
	}

	o.encryptValue()
	return
}

func (o *Option) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = o.Entity.EntityBase.BeforeUpdate(tx); err != nil {
		return
	}

	o.encryptValue()
	return
}

func (o *Option) AfterFind(tx *gorm.DB) (err error) {
	if err = o.Entity.EntityBase.AfterFind(tx); err != nil {
		return
	}

	o.decryptValue()
	return
}

func (o *Option) encryptValue() {
	if o.Encrypted {
		o.Value = "" // encrypt
	}
}

func (o *Option) decryptValue() {
	if o.Encrypted {
		o.Value = "" // decrypt
	}
}

func NewOption(snowflake iSnowflake.ISnowflake, option *Option) *Option {
	option.snowflake = snowflake
	return option
}
