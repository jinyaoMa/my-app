package crud

import (
	"my-app/internal/entity"
	"my-app/pkg/crypto"
	"my-app/pkg/db"

	"gorm.io/gorm"
)

type Option struct {
	db.ICRUD[*entity.Option]
	session *gorm.DB
	crypto  crypto.ICrypto
}

func (option *Option) Save(entity *entity.Option) (affected int64, err error) {
	if entity.Encrypted {
		cipher, err := option.crypto.Encrypt(entity.Value)
		if err != nil {
			return 0, err
		}
		entity.Value = cipher
	}
	return option.ICRUD.Save(entity)
}

func (option *Option) SaveAll(entities []*entity.Option) (affected int64, err error) {
	count := len(entities)
	for i := 0; i < count; i++ {
		if entities[i].Encrypted {
			cipher, err := option.crypto.Encrypt(entities[i].Value)
			if err != nil {
				return 0, err
			}
			entities[i].Value = cipher
		}
	}
	return option.ICRUD.SaveAll(entities)
}

func NewOption(session *gorm.DB, crypto crypto.ICrypto) (option *Option, iOption IOption) {
	_, crud := db.NewCRUD[*entity.Option](session)
	option = &Option{
		ICRUD:   crud,
		session: session,
		crypto:  crypto,
	}
	return option, option
}
