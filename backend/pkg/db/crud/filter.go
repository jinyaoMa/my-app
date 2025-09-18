package crud

import (
	"regexp"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db"
)

var regFilterField = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_\.]*$`)

type Filter struct {
	Or      bool          `json:"or" required:"false" doc:"Or or And"`
	Type    FilterType    `json:"type" doc:"Filter Type"`
	Field   string        `json:"field" doc:"Filter Field Name"`
	Params  []any         `json:"params" doc:"Filter Condition Parameters"`
	Special FilterSpecial `json:"special" required:"false" doc:"Filter Special"`
}

func (f *Filter) Apply(tx *gorm.DB) *gorm.DB {
	if !regFilterField.MatchString(f.Field) {
		return tx
	}

	condition := ""
	switch f.Type {
	case FilterTypeEqual:
		condition = f.Field + " = ?"
	case FilterTypeNotEqual:
		condition = f.Field + " <> ?"
	case FilterTypeLessThan:
		condition = f.Field + " < ?"
	case FilterTypeLessThanOrEqual:
		condition = f.Field + " <= ?"
	case FilterTypeGreaterThan:
		condition = f.Field + " > ?"
	case FilterTypeGreaterThanOrEqual:
		condition = f.Field + " >= ?"
	case FilterTypeLike:
		condition = f.Field + " LIKE ?"
	case FilterTypeNotLike:
		condition = f.Field + " NOT LIKE ?"
	case FilterTypeNull:
		condition = f.Field + " IS NULL"
	case FilterTypeNotNull:
		condition = f.Field + " IS NOT NULL"
	case FilterTypeBetween:
		condition = f.Field + " BETWEEN ? AND ?"
	case FilterTypeNotBetween:
		condition = f.Field + " NOT BETWEEN ? AND ?"
	case FilterTypeIn:
		condition = f.Field + " IN ?"
	case FilterTypeNotIn:
		condition = f.Field + " NOT IN ?"
	default:
		return tx
	}

	if f.Type == FilterTypeNull || f.Type == FilterTypeNotNull {
		if f.Or {
			tx = tx.Or(condition)
		} else {
			tx = tx.Where(condition)
		}
		return tx
	}

	if f.Or {
		tx = tx.Or(condition, f.Params...)
		return tx
	}

	switch f.Special {
	case FilterSpecialIdString:
		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				if id, err := db.ConvertStringToId(v); err == nil {
					f.Params[i] = id
				}
			}
		}
	case FilterSpecialEncrypted:
		cipher, ok := db.GetCipher(tx)
		if !ok {
			return tx
		}

		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				f.Params[i] = cipher.EncryptBase64(v)
			}
		}
	case FilterSpecialHashed:
		hasher, ok := db.GetHasher(tx)
		if !ok {
			return tx
		}

		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				f.Params[i] = hasher.HashBase64(v)
			}
		}
	case FilterSpecialPassword:
		keygen, ok := db.GetKeygen(tx)
		if !ok {
			return tx
		}

		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				f.Params[i] = keygen.DeriveBase64(v, true)
			}
		}
	}
	tx = tx.Where(condition, f.Params...)
	return tx
}
