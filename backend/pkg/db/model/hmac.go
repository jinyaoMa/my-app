package model

import (
	"errors"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type ModelHmac struct {
	Hmac   datatype.Secret `gorm:"<-:update;comment:HMAC;"`
	HmacOk bool            `gorm:"-"`
}

func (h *ModelHmac) Update(tx *gorm.DB, model any, data func() string) (err error) {
	hasher, ok := dbcontext.GetHasher(tx)
	if !ok {
		return errors.New("db context does not contain hasher")
	}

	if err = tx.Model(model).
		UpdateColumn("hmac", hasher.HashBase64(data())).Error; err != nil {
		return
	}
	h.HmacOk = true
	return
}

func (h *ModelHmac) Verify(tx *gorm.DB, data func() string) (err error) {
	h.HmacOk, err = h.Hmac.Verify(tx, datatype.Secret{
		Data: data(),
	})
	return
}
