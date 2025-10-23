package crud

import "gorm.io/gorm"

type Includes []Include

func (i Includes) Apply(tx *gorm.DB, maxSize int, maxLength int, maxSortLength int) *gorm.DB {
	length := len(i)
	if length > 0 && length > maxLength {
		return tx
	}
	for _, include := range i {
		tx = include.Apply(tx, maxSize, maxLength, maxSortLength)
	}
	return tx
}

type Include struct {
	Choice
	Association string                     `json:"association" required:"true" doc:"Included Association Name"`
	Sorts       Sorts                      `json:"sorts" required:"false" doc:"List Sort Conditions"`
	Includes    Includes                   `json:"includes" required:"false" doc:"Included Associations"`
	Size        int                        `json:"size" required:"false" doc:"Size Limit"`
	IgnoreSize  bool                       `json:"-"` // ignore size limit
	Condition   func(tx *gorm.DB) *gorm.DB `json:"-"` // additional condition
}

func (i *Include) Apply(tx *gorm.DB, maxSize int, maxLength int, maxSortLength int) *gorm.DB {
	tx = tx.Preload(i.Association, func(tx *gorm.DB) *gorm.DB {
		tx = i.Includes.Apply(tx, maxSize, maxLength, maxSortLength)
		tx = i.Sorts.Apply(tx, maxSortLength)
		if i.Condition != nil {
			tx = i.Condition(tx)
		}
		if i.IgnoreSize {
			return i.Choice.Apply(tx)
		}

		limit := i.Size
		if limit < 1 || limit > maxSize {
			limit = 1
		}
		return i.Choice.Apply(tx.Limit(limit))
	})
	return tx
}
