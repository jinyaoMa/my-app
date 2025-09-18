package datatype

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm/schema"
	"majinyao.cn/my-app/backend/pkg/db"
)

type Encrypted string

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (e *Encrypted) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) (err error) {
	cipher, ok := db.GetCipherFromContext(ctx)
	if !ok {
		return errors.New("db context does not contain cipher")
	}

	switch value := dbValue.(type) {
	case string:
		if plaintext, err := cipher.DecryptBase64(value); err == nil {
			*e = Encrypted(plaintext)
		} else {
			*e = Encrypted(value)
		}
	default:
		return fmt.Errorf("db datatype encrypted: unsupported data %#v", dbValue)
	}
	return nil
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// fieldValue: current field's value of the dst
func (e Encrypted) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error) {
	cipher, ok := db.GetCipherFromContext(ctx)
	if !ok {
		return nil, errors.New("db context does not contain cipher")
	}
	return cipher.EncryptBase64(string(e)), nil
}
