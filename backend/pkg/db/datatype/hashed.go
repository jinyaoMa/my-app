package datatype

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type Hashed string

func (s Hashed) VerifyBase64(tx *gorm.DB, data string) (ok bool, err error) {
	hasher, ok := dbcontext.GetHasher(tx)
	if !ok {
		return false, errors.New("db context does not contain hasher")
	}
	return hasher.VerifyBase64(string(s), data), nil
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (h *Hashed) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) (err error) {
	switch value := dbValue.(type) {
	case string:
		*h = Hashed(value)
	default:
		return fmt.Errorf("db datatype hashed: unsupported data %#v", dbValue)
	}
	return nil
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// fieldValue: current field's value of the dst
func (h Hashed) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error) {
	hasher, ok := dbcontext.GetHasherFromContext(ctx)
	if !ok {
		return nil, errors.New("db context does not contain hasher")
	}
	return hasher.HashBase64(string(h)), nil
}
