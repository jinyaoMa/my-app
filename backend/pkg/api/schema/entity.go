package schema

import (
	"time"

	"majinyao.cn/my-app/backend/pkg/db"
)

type EntityIdGetter interface {
	GetId() (id int64, err error)
	GetIdString() string
	IsTransient() bool
}

type EntityId struct {
	Id string `json:"id" doc:"Entity Id"`
}

func (i EntityId) GetId() (id int64, err error) {
	return db.ConvertStringToId(i.Id)
}

func (i EntityId) GetIdString() string {
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
