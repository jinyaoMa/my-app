package schema

import (
	"time"

	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type EntityIdGetter interface {
	GetId() (id datatype.Id, err error)
	GetIdHexString() string
	IsTransient() bool
}

type EntityId struct {
	Id string `json:"id" doc:"Entity Id (Hex)"`
}

func (i EntityId) GetId() (id datatype.Id, err error) {
	return datatype.ParseIdFromHex(i.Id)
}

func (i EntityId) GetIdHexString() string {
	return i.Id
}

func (i EntityId) IsTransient() bool {
	return i.Id == ""
}

type EntityTime struct {
	CreatedAt time.Time  `json:"createAt" doc:"Created At"`
	UpdatedAt time.Time  `json:"updateAt" doc:"Updated At"`
	DeletedAt *time.Time `json:"deleteAt,omitempty" doc:"Deleted At"`
}
