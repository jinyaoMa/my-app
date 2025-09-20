package schema

import (
	"time"

	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type ModelIdGetter interface {
	GetId() (id datatype.Id, err error)
	GetIdB36() string
}

type ModelId struct {
	Id string `json:"id" doc:"Model Id (Base36)"`
}

func (i ModelId) GetId() (id datatype.Id, err error) {
	return datatype.ParseIdFromB36(i.Id)
}

func (i ModelId) GetIdB36() string {
	return i.Id
}

type ModelTime struct {
	CreatedAt time.Time  `json:"createAt" doc:"Created At"`
	UpdatedAt time.Time  `json:"updateAt" doc:"Updated At"`
	DeletedAt *time.Time `json:"deleteAt,omitempty" doc:"Deleted At"`
}
