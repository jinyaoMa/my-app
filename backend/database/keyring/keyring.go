package keyring

import (
	"my-app/backend/database"

	"gorm.io/gorm"
)

type Keyring struct {
	gorm.Model
	Alias      string `` // Keyring name alias
	Account    string `` // Keyring user account
	Password   string `` // Keyring user password
	More       string `` // more information about security questions, email, password history, etc.
	IsEncypted bool   `` // indicate if Password and More are encrypted
}

func init() {
	database.DB().AutoMigrate(&Keyring{})
}

func (k *Keyring) BeforeCreate(tx *gorm.DB) (err error) {
	if !k.IsEncypted {
		if k.Password, err = database.AES().Encrypt(k.Password); err != nil {
			return
		}
		if k.More, err = database.AES().Encrypt(k.More); err != nil {
			return
		}
		k.IsEncypted = true
	}
	return
}

func (k *Keyring) AfterFind(tx *gorm.DB) (err error) {
	if k.IsEncypted {
		if k.Password, err = database.AES().Decrypt(k.Password); err != nil {
			return
		}
		if k.More, err = database.AES().Decrypt(k.More); err != nil {
			return
		}
		k.IsEncypted = false
	}
	return
}

func (k *Keyring) Create() (ok bool) {
	result := database.DB().Create(k)
	return result.RowsAffected == 1
}

func (k *Keyring) Find() (ok bool) {
	result := database.DB().Where(k).Find(k)
	return result.RowsAffected > 0
}
