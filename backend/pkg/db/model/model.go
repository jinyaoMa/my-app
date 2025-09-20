package model

import (
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IdGetter interface {
	GetId() datatype.Id
	IsTransient() bool
}

type Model struct {
	Id        datatype.Id    `gorm:"primaryKey;comment:Id;"`
	CreatedAt time.Time      `gorm:"<-:create;autoCreateTime;comment:Created At;"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:Updated At;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:Deleted At;"`
}

func (m Model) GetId() datatype.Id {
	return m.Id
}

func (m Model) IsTransient() bool {
	return m.Id.Int64() == 0
}

func (m *Model) BeforeSave(tx *gorm.DB) (err error) {
	if m.IsTransient() {
		s, ok := dbcontext.GetSnowflake(tx)
		if ok {
			m.Id = datatype.Id(s.Generate())
		}
	}
	return nil
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if m.IsTransient() {
		s, ok := dbcontext.GetSnowflake(tx)
		if ok {
			m.Id = datatype.Id(s.Generate())
		}
	}
	return nil
}
