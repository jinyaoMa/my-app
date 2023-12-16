package entity

import (
	"my-app/pkg/crypto"
	"my-app/pkg/db"
)

type Option struct {
	db.Entity[*Option]
	Name      string `gorm:"unique; index; not null" json:"name"`
	Value     string `gorm:"default:''" json:"value"` // may be encrypted
	Encrypted bool   `gorm:"default:false" json:"encrypted"`
}

func (option *Option) Encrypt(crypto crypto.ICrypto) (err error) {
	if !option.Encrypted {
		ciphertext, err := crypto.Encrypt(option.Value)
		if err != nil {
			return err
		}
		option.Value = ciphertext
		option.Encrypted = true
	}
	return
}

func (option *Option) Decrypt(crypto crypto.ICrypto) (err error) {
	if option.Encrypted {
		plaintext, err := crypto.Decrypt(option.Value)
		if err != nil {
			return err
		}
		option.Value = plaintext
		option.Encrypted = false
	}
	return
}
