package model

import "gorm.io/gorm"

type MyOption struct {
	gorm.Model
	Name  string `gorm:"unique"` // Option name
	Value string ``              // Option value associated with name
}

func (mo *MyOption) Update(newValue string) *gorm.DB {
	return db.Model(mo).Where(mo).Updates(MyOption{
		Value: newValue,
	})
}

type MyOptions []MyOption

func (mos *MyOptions) Load() *gorm.DB {
	return db.Find(mos)
}

func (mos *MyOptions) Save() *gorm.DB {
	return db.Save(mos)
}
