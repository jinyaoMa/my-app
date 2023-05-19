package entity

import (
	"gorm.io/gorm"
)

type Option struct {
	Entity

	/* internal fields */
	Key            string `gorm:"size:128; unique; index; not null"`
	Value          string `gorm:"size:2048; default:''"`
	ValueEncrypted string `gorm:"size:2048; default:''"`
	Encrypted      bool   `gorm:"default:false"`

	/* relational fields */
}

func (o *Option) BeforeCreate(tx *gorm.DB) (err error) {
	if err = o.Entity.BeforeCreate(tx); err != nil {
		return
	}

	if o != nil {
		if err = o.encryptValue(tx); err != nil {
			return
		}
	}
	return
}

func (o *Option) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = o.Entity.BeforeUpdate(tx); err != nil {
		return
	}

	if o != nil {
		if err = o.encryptValue(tx); err != nil {
			return
		}
	}
	return
}

func (o *Option) AfterFind(tx *gorm.DB) (err error) {
	if err = o.Entity.AfterFind(tx); err != nil {
		return
	}

	if o != nil {
		if err = o.decryptValue(tx); err != nil {
			return
		}
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
