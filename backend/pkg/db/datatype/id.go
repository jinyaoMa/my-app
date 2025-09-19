package datatype

import (
	"context"
	"encoding/hex"
	"fmt"
	"reflect"

	"gorm.io/gorm/schema"
	"majinyao.cn/my-app/backend/pkg/utils"
)

func ParseIdFromHex(hexStr string) (Id, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return 0, err
	}
	return Id(utils.ConvertBytesToInt64(buf)), nil
}

type Id int64

func (i Id) HexString() string {
	return utils.ConvertInt64ToHex(int64(i))
}

func (i Id) Int64() int64 {
	return int64(i)
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (o *Id) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) (err error) {
	switch value := dbValue.(type) {
	case int64:
		*o = Id(value)
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
