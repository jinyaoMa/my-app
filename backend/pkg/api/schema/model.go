package schema

import (
	"time"

	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type ModelId struct {
	Id string `json:"id" required:"false" doc:"Model Id (Base36)"`
}

func (i ModelId) GetId() (id datatype.Id, err error) {
	return datatype.ParseIdFromB36(i.Id)
}

func (i ModelId) GetIdB36() string {
	return i.Id
}

type ModelTime struct {
	CreatedAt time.Time `json:"createdAt" doc:"Created At"`
	UpdatedAt time.Time `json:"updatedAt" doc:"Updated At"`
	DeletedAt time.Time `json:"deletedAt" doc:"Deleted At"`
}

type ModelReserved struct {
	Reserved bool `json:"reserved" doc:"Is System Reserved"`
}

type ModelHmac struct {
	HmacOk bool `json:"hmacOk" doc:"Is Hmac Ok"`
}
