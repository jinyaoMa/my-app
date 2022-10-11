package keyring

import (
	"my-app/backend/database"
)

type Keyrings []Keyring

func (ks *Keyrings) Find() (ok bool) {
	result := database.DB().Find(ks)
	return result.RowsAffected > 0
}
