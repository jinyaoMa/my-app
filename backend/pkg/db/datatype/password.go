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

type Password struct {
	Data    string
	Derived bool
}

func (p Password) Verify(db *gorm.DB, other Password) (ok bool, err error) {
	if p.Derived == other.Derived {
		return p.Data == other.Data, nil
	}

	keygen, ok := dbcontext.GetKeygen(db)
	if !ok {
		return false, errors.New("db context does not contain keygen")
	}
	if p.Derived {
		ok = keygen.VerifyBase64(p.Data, other.Data, true)
	} else {
		ok = keygen.VerifyBase64(other.Data, p.Data, true)
	}
	return
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (p *Password) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) error {
	if dbValue == nil {
		return nil
	}

	switch value := dbValue.(type) {
	case string:
		*p = Password{
			Data:    value,
			Derived: true,
		}
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
	if p.Derived {
		return p.Data, nil
	}

	keygen, ok := dbcontext.GetKeygenFromContext(ctx)
	if !ok {
		return nil, errors.New("db context does not contain keygen")
	}
	return keygen.DeriveBase64(p.Data, true), nil
}
