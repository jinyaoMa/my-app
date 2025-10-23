package datatype

import (
	"context"
	"fmt"
	"reflect"

	"gorm.io/gorm/schema"
	"majinyao.cn/my-app/backend/pkg/utils"
)

func ParseIdFromB36(b36 string) (Id, error) {
	v, err := utils.ConvertB36ToInt64(b36)
	if err != nil {
		return 0, fmt.Errorf("db datatype id: id %s is not valid", b36)
	}
	return Id(v), nil
}

type Id int64

func (i Id) B36String() string {
	return utils.ConvertInt64ToB36(int64(i))
}

func (i Id) Int64() int64 {
	return int64(i)
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (i *Id) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) error {
	if dbValue == nil {
		return nil
	}

	switch value := dbValue.(type) {
	case int64:
		*i = Id(value)
	default:
		return fmt.Errorf("db datatype oid: unsupported data %#v", dbValue)
	}
	return nil
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// fieldValue: current field's value of the dst
func (i Id) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error) {
	return i.Int64(), nil
}
