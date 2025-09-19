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

type Password string

func (p Password) VerifyBase64(tx *gorm.DB, password string) (ok bool, err error) {
	keygen, ok := dbcontext.GetKeygen(tx)
	if !ok {
		return false, errors.New("db context does not contain keygen")
	}
	return keygen.VerifyBase64(string(p), password, true), nil
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (p *Password) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) (err error) {
	switch value := dbValue.(type) {
	case string:
		*p = Password(value)
	default:
		return fmt.Errorf("db datatype password: unsupported data %#v", dbValue)
	}
	return nil
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// fieldValue: current field's value of the dst
func (p Password) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error) {
	keygen, ok := dbcontext.GetKeygenFromContext(ctx)
	if !ok {
		return nil, errors.New("db context does not contain keygen")
	}
	return keygen.DeriveBase64(string(p), true), nil
}
