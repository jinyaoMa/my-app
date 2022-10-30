package model

import (
	"gorm.io/gorm"
)

// option names
const (
	OptionDisplayLanguage OptionName = "DisplayLanguage"
	OptionColorTheme      OptionName = "ColorTheme"

	OptionFileLog      OptionName = "FileLog"
	OptionDirLanguages OptionName = "DirLanguages"
	OptionDirAssets    OptionName = "DirAssets"
	OptionDirUserData  OptionName = "DirUserData"
	OptionDirDocs      OptionName = "DirDocs"

	OptionWebAutoStart OptionName = "Web.AutoStart"
	OptionWebPortHttp  OptionName = "Web.PortHttp"
	OptionWebPortHttps OptionName = "Web.PortHttps"
	OptionWebDirCerts  OptionName = "Web.DirCerts"
)

type OptionName string

func (on OptionName) ToString() string {
	return string(on)
}

type Option struct {
	gorm.Model
	Name  OptionName `gorm:"unique"` // Option name
	Value string     ``              // Option value associated with name
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
		// update the old value
		found.Value = o.Value
		// assign current option back
		o = found
	}
	// insert or save the option
	return o.Save(db)
}

type Options []*Option

func (os Options) IndexOf(o *Option) int {
	for i, opt := range os {
		if opt.Name == o.Name {
			return i
		}
	}
	return -1
}

func (os Options) Contains(o *Option) bool {
	return os.IndexOf(o) >= 0
}

func (os Options) Find(db *gorm.DB) (ok bool) {
	tx := db.Find(os)
	return tx.RowsAffected > 0
}

func (os Options) Save(db *gorm.DB) (ok bool) {
	tx := db.Save(os)
	return tx.Error == nil
}

func (os Options) FindAndSave(db *gorm.DB) (ok bool) {
	var founds Options
	if founds.Find(db) { // already some options in db
		for _, opt := range os {
			if i := founds.IndexOf(opt); i >= 0 {
				// update the old value
				founds[i].Value = opt.Value
			} else {
				// insert new option
				founds = append(founds, opt)
			}
		}
		// assign current options back
		os = founds
	}
	// insert or save all options
	return os.Save(db)
}
