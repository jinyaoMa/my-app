package model

import (
	"gorm.io/gorm"
)

type Option struct {
	gorm.Model
	Name  string `gorm:"unique"` // Option name
	Value string ``              // Option value associated with name
}

func (o *Option) Find() (ok bool) {
	result := db.Where(o).Find(o)
	return result.RowsAffected > 0
}

func (o *Option) Update(newValue string) (ok bool) {
	result := db.Model(o).Where(o).Updates(Option{
		Value: newValue,
	})
	return result.RowsAffected == 1
}

type Options []Option

func (os *Options) Find() (ok bool) {
	result := db.Find(os)
	return result.RowsAffected > 0
}

func (os *Options) Save() (ok bool) {
	result := db.Save(os)
	return result.Error == nil
}
