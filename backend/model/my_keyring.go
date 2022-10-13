package model

import (
	"my-app/backend/pkg/utils"

	"gorm.io/gorm"
)

type Keyring struct {
	gorm.Model
	Alias    string `` // Keyring name alias
	Account  string `` // Keyring user account
	Password string `` // Keyring user password
	More     string `` // more information about security questions, email, password history, etc.
	UserID   uint

	AES        *utils.AES `gorm:"-:all"` // aes tool used to encrypt/decrypt
	IsEncypted bool       `gorm:"-:all"` // indicate if Password and More are encrypted
}

func (k *Keyring) BeforeCreate(tx *gorm.DB) (err error) {
	if !k.IsEncypted {
		if k.Password, err = k.AES.Encrypt(k.Password); err != nil {
			return
		}
		if k.More, err = k.AES.Encrypt(k.More); err != nil {
			return
		}
		k.IsEncypted = true
	}
	return
}

func (k *Keyring) AfterFind(tx *gorm.DB) (err error) {
	if k.Password, err = k.AES.Decrypt(k.Password); err != nil {
		return
	}
	if k.More, err = k.AES.Decrypt(k.More); err != nil {
		return
	}
	k.IsEncypted = false
	return
}

func (k *Keyring) Find() (ok bool) {
	tx := db.Where(k).Find(k)
	return tx.RowsAffected > 0
}

func (k *Keyring) Create() (ok bool) {
	tx := db.Create(k)
	return tx.RowsAffected == 1
}

type Keyrings []Keyring

func (ks *Keyrings) Find(selects ...string) (ok bool) {
	var tx *gorm.DB
	if len(selects) > 0 {
		tx = db.Select(selects).Find(ks)
	} else {
		tx = db.Find(ks)
	}
	return tx.RowsAffected > 0
}

func (ks *Keyrings) FindByUser(userid uint, selects ...string) (ok bool) {
	var tx *gorm.DB
	if len(selects) > 0 {
		tx = db.Select(selects).Where(Keyring{
			UserID: userid,
		}).Find(ks)
	} else {
		tx = db.Where(Keyring{
			UserID: userid,
		}).Find(ks)
	}
	return tx.RowsAffected > 0
}
