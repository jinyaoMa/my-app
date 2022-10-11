package keyring

import (
	"my-app/backend/database"
	"my-app/backend/pkg/utils"

	"gorm.io/gorm"
)

var aes *utils.AES

type Keyring struct {
	gorm.Model
	Alias      string `` // Keyring name alias
	Account    string `` // Keyring user account
	Password   string `` // Keyring user password
	More       string `` // more information about security questions, email, password history, etc.
	IsEncypted bool   `` // indicate if Password and More are encrypted
}

func init() {
	aes = utils.NewAES("test")
	database.DB().AutoMigrate(&Keyring{})
}

func (k *Keyring) BeforeCreate(tx *gorm.DB) (err error) {
	if !k.IsEncypted {
		if k.Password, err = aes.Encrypt(k.Password); err != nil {
			return
		}
		if k.More, err = aes.Encrypt(k.More); err != nil {
			return
		}
		k.IsEncypted = true
	}
	return
}

func (k *Keyring) AfterFind(tx *gorm.DB) (err error) {
	if k.IsEncypted {
		if k.Password, err = aes.Decrypt(k.Password); err != nil {
			return
		}
		if k.More, err = aes.Decrypt(k.More); err != nil {
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
