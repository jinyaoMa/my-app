package crud

import (
	"strings"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
	"majinyao.cn/my-app/backend/pkg/utils"
)

const FilterParamsMaxLength = 20

type Filters []Filter

func (f Filters) Apply(tx *gorm.DB, paramsMaxLength int) *gorm.DB {
	for _, filter := range f {
		tx = filter.Apply(tx, paramsMaxLength)
	}
	return tx
}

type Filter struct {
	Or      bool          `json:"or" required:"false" doc:"Or or And"`
	Type    FilterType    `json:"type" required:"true" doc:"Filter Type: equal, notEqual, lessThan, lessThanOrEqual, greaterThan, greaterThanOrEqual, contains, notContains, startsWith, notStartsWith, endsWith, notEndsWith, null, notNull, between, notBetween, in, notIn"`
	Field   string        `json:"field" required:"true" doc:"Filter Field Name"`
	Params  []any         `json:"params" required:"false" doc:"Filter Condition Parameters"`
	Special FilterSpecial `json:"special" required:"false" doc:"Filter Special: idString, sensitive, secret, password, oid"`
}

func (f *Filter) Apply(tx *gorm.DB, paramsMaxLength int) *gorm.DB {
	if len(f.Params) > paramsMaxLength || !regField.MatchString(f.Field) {
		return tx
	}

	switch f.Type {
	case FilterTypeContains, FilterTypeNotContains:
		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				f.Params[i] = "%" + FixLikeKeyword(v) + "%"
			}
		}
	case FilterTypeStartsWith, FilterTypeNotStartsWith:
		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				f.Params[i] = FixLikeKeyword(v) + "%"
			}
		}
	case FilterTypeEndsWith, FilterTypeNotEndsWith:
		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				f.Params[i] = "%" + FixLikeKeyword(v)
			}
		}
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
	case FilterTypeContains, FilterTypeStartsWith, FilterTypeEndsWith:
		condition = f.Field + " LIKE ? ESCAPE '\\'"
	case FilterTypeNotContains, FilterTypeNotStartsWith, FilterTypeNotEndsWith:
		condition = f.Field + " NOT LIKE ? ESCAPE '\\'"
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
			return tx.Or(condition)
		} else {
			return tx.Where(condition)
		}
	}

	switch f.Special {
	case FilterSpecialIdString:
		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				if id, err := datatype.ParseIdFromB36(v); err == nil {
					f.Params[i] = id
				}
			}
		}
	case FilterSpecialSensitive:
		cipher, ok := dbcontext.GetCipher(tx)
		if !ok {
			return tx
		}

		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				f.Params[i] = cipher.EncryptBase64(v)
			}
		}
	case FilterSpecialSecret:
		hasher, ok := dbcontext.GetHasher(tx)
		if !ok {
			return tx
		}

		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				f.Params[i] = hasher.HashBase64(v)
			}
		}
	case FilterSpecialPassword:
		keygen, ok := dbcontext.GetKeygen(tx)
		if !ok {
			return tx
		}

		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				f.Params[i] = keygen.DeriveBase64(v, true)
			}
		}
	case FilterSpecialOid:
		var oid []int64
		for i := range f.Params {
			if v, ok := f.Params[i].(string); ok {
				for id := range strings.SplitSeq(v, datatype.OidDelimiter) {
					if id, err := utils.ConvertB36ToInt64(id); err == nil {
						oid = append(oid, id)
					}
				}
			}
		}
		f.Params = []any{datatype.Oid(oid)}
	}

	if f.Or {
		return tx.Or(condition, f.Params...)
	}
	return tx.Where(condition, f.Params...)
}
