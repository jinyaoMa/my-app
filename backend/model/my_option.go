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
	tx := db.Where(o).Find(o)
	return tx.RowsAffected > 0
}

func (o *Option) Update(newValue string) (ok bool) {
	tx := db.Model(&Option{}).Where(o).Updates(Option{
		Value: newValue,
	})
	return tx.RowsAffected == 1
}

type Options []Option

func (os *Options) Find() (ok bool) {
	tx := db.Find(os)
	return tx.RowsAffected > 0
}

func (os *Options) Save() (ok bool) {
	tx := db.Save(os)
	return tx.Error == nil
}
