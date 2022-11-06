package model

import (
	"my-app/backend/app/types"

	"gorm.io/gorm"
)

type Option struct {
	gorm.Model
	Name  types.ConfigName `gorm:"unique"` // Option name
	Value string           ``              // Option value associated with name
}

func (o *Option) FindByName(db *gorm.DB) (ok bool) {
	tx := db.Where(Option{
		Name: o.Name,
	}).Find(o)
	return tx.RowsAffected > 0
}

func (o *Option) Save(db *gorm.DB) (ok bool) {
	tx := db.Save(o)
	return tx.RowsAffected == 1
}

func (o *Option) FindByNameAndSave(db *gorm.DB) (ok bool) {
	found := &Option{
		Name: o.Name,
	}
	if found.FindByName(db) { // option already exists in db
		o.Model = found.Model
	}
	// insert or save the option
	return o.Save(db)
}

type Options []*Option

func (os *Options) IndexOf(o *Option) int {
	for i, opt := range *os {
		if opt.Name == o.Name {
			return i
		}
	}
	return -1
}

func (os *Options) Contains(o *Option) bool {
	return os.IndexOf(o) >= 0
}

func (os *Options) Find(db *gorm.DB) (ok bool) {
	tx := db.Find(os)
	return tx.RowsAffected > 0
}

func (os *Options) Save(db *gorm.DB) (ok bool) {
	tx := db.Save(os)
	return tx.Error == nil
}

func (os *Options) FindAndSave(db *gorm.DB) (ok bool) {
	var founds Options
	if founds.Find(db) { // already some options in db
		for _, opt := range *os {
			if i := founds.IndexOf(opt); i >= 0 {
				// update the option with stored value
				opt.Model = founds[i].Model
				opt.Value = founds[i].Value
			}
		}
	}
	// insert or save all options
	return os.Save(db)
}
