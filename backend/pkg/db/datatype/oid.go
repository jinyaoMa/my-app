package datatype

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm/schema"
	"majinyao.cn/my-app/backend/pkg/utils"
)

const OidDelimiter = "."

func ParseOid(s string) (Oid, error) {
	if s == "" {
		return nil, nil
	}
	if !strings.HasPrefix(s, OidDelimiter) {
		return nil, fmt.Errorf("db datatype oid: oid %s is not valid", s)
	}

	s = s[1:]
	return utils.SliceMap(strings.Split(s, OidDelimiter), func(s string) int64 {
		v, _ := utils.ConvertB36ToInt64(s)
		return v
	}), nil
}

type Oid []int64

func (o Oid) String() string {
	if o == nil {
		return ""
	}

	return OidDelimiter + strings.Join(utils.SliceMap(o, func(v int64) string {
		return utils.ConvertInt64ToB36(v)
	}), OidDelimiter)
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (o *Oid) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) error {
	switch value := dbValue.(type) {
	case string:
		oid, err := ParseOid(value)
		if err != nil {
			return err
		}
		*o = oid
	default:
		return fmt.Errorf("db datatype oid: unsupported data %#v", dbValue)
	}
	return nil
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// fieldValue: current field's value of the dst
func (o Oid) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error) {
	return o.String(), nil
}
