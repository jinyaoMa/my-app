package db

import (
	"time"

	"gorm.io/gorm"
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
	GetId() int64
	GetIdString() string
}

type Entity struct {
	Id        int64          `gorm:"primaryKey;comment:Id;"`
	CreatedAt time.Time      `gorm:"<-:create;autoCreateTime;comment:Created At;"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:Updated At;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:Deleted At;"`
}

func (e Entity) GetId() int64 {
	return e.Id
}

func (e Entity) GetIdString() string {
	return ConvertIdToString(e.Id)
}

func (e *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	s, ok := GetSnowflake(tx)
	if ok {
		e.Id = s.Generate()
	}
	return nil
}
