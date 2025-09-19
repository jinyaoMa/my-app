package db

import (
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type EntityM2MSetup struct {
	Model     any
	Field     string
	JoinTable any
}

type EntityM2MSetupsGetter interface {
	GetEntityM2MSetups() []EntityM2MSetup
}

type EntityIdGetter interface {
	GetId() datatype.Id
}

type Entity struct {
	Id        datatype.Id    `gorm:"primaryKey;comment:Id;"`
	CreatedAt time.Time      `gorm:"<-:create;autoCreateTime;comment:Created At;"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:Updated At;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:Deleted At;"`
}

func (e Entity) GetId() datatype.Id {
	return e.Id
}

func (e *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	s, ok := dbcontext.GetSnowflake(tx)
	if ok {
		e.Id = datatype.Id(s.Generate())
	}
	return nil
}
