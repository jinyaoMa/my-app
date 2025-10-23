package crud

import "gorm.io/gorm"

type Joins []string

func (j Joins) Apply(tx *gorm.DB, maxLength int) *gorm.DB {
	length := len(j)
	if length > 0 && length > maxLength {
		return tx
	}
	for _, join := range j {
		if regField.MatchString(join) {
			tx = tx.Joins(join)
		}
	}
	return tx
}
