package datatype

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm/schema"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type Sensitive string

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (s *Sensitive) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) error {
	if dbValue == nil {
		return nil
	}

	cipher, ok := dbcontext.GetCipherFromContext(ctx)
	if !ok {
		return errors.New("db context does not contain cipher")
	}

	switch value := dbValue.(type) {
	case string:
		if plaintext, err := cipher.DecryptBase64(value); err == nil {
			*s = Sensitive(plaintext)
		} else {
			*s = Sensitive(value)
		}
	case []byte:
		if plaindata, err := cipher.Decrypt(value); err == nil {
			*s = Sensitive(plaindata)
		} else {
			*s = Sensitive(value)
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
func (s Sensitive) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error) {
	cipher, ok := dbcontext.GetCipherFromContext(ctx)
	if !ok {
		return nil, errors.New("db context does not contain cipher")
	}
	return cipher.EncryptBase64(string(s)), nil
}
