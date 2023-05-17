package entity

import (
	"gorm.io/gorm"
)

type Option struct {
	Entity
	Key            string `gorm:"size:100; unique; index"`
	Value          string `gorm:"size:100"`
	ValueEncrypted string `gorm:"size:255"`
	Encrypted      bool   `gorm:""`
}

func (o *Option) BeforeCreate(tx *gorm.DB) (err error) {
	if err = o.Entity.EntityBase.BeforeCreate(tx); err != nil {
		return
	}

	if o != nil {
		err = o.encryptValue(tx)
	}
	return
}

func (o *Option) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = o.Entity.EntityBase.BeforeUpdate(tx); err != nil {
		return
	}

	if o != nil {
		err = o.encryptValue(tx)
	}
	return
}

func (o *Option) AfterFind(tx *gorm.DB) (err error) {
	if err = o.Entity.EntityBase.AfterFind(tx); err != nil {
		return
	}

	if o != nil {
		err = o.decryptValue(tx)
	}
	return
}

func (o *Option) encryptValue(tx *gorm.DB) (err error) {
	if o.Encrypted && aes != nil {
		var ciphertext string
		ciphertext, err = aes.Encrypt(o.Value)
		if err != nil {
			return
		}
		o.ValueEncrypted = ciphertext
		tx.Statement.Omit("Value")
	}
	return
}

func (o *Option) decryptValue(tx *gorm.DB) (err error) {
	if o.Encrypted && aes != nil {
		var plaintext string
		plaintext, err = aes.Decrypt(o.ValueEncrypted)
		if err != nil {
			return
		}
		o.Value = plaintext
	}
	return
}
