package option

import (
	"my-app/backend/database"

	"gorm.io/gorm"
)

type Option struct {
	gorm.Model
	Name  string `gorm:"unique"` // Option name
	Value string ``              // Option value associated with name
}

func init() {
	database.DB().AutoMigrate(&Option{})
}

func (o *Option) Find() (ok bool) {
	result := database.DB().Where(o).Find(o)
	return result.RowsAffected > 0
}

func (o *Option) Update(newValue string) (ok bool) {
	result := database.DB().Model(o).Where(o).Updates(Option{
		Value: newValue,
	})
	return result.RowsAffected == 1
}
