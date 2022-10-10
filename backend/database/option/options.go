package option

import (
	"my-app/backend/database"
)

type Options []Option

func (os *Options) Find() (ok bool) {
	result := database.DB().Find(os)
	return result.RowsAffected > 0
}

func (os *Options) Save() (ok bool) {
	result := database.DB().Save(os)
	return result.Error == nil
}
