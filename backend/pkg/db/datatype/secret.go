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

type Secret struct {
	Data   string
	Hashed bool
}

func (s Secret) Verify(db *gorm.DB, other Secret) (ok bool, err error) {
	if s.Hashed == other.Hashed {
		return s.Data == other.Data, nil
	}

	hasher, ok := dbcontext.GetHasher(db)
	if !ok {
		return false, errors.New("db context does not contain hasher")
	}
	if s.Hashed {
		ok = hasher.VerifyBase64(s.Data, other.Data)
	} else {
		ok = hasher.VerifyBase64(other.Data, s.Data)
	}
	return
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (s *Secret) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) error {
	if dbValue == nil {
		return nil
	}

	switch value := dbValue.(type) {
	case string:
		*s = Secret{
			Data:   value,
			Hashed: true,
		}
	default:
		return fmt.Errorf("db datatype hashed: unsupported data %#v", dbValue)
	}
	return nil
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// fieldValue: current field's value of the dst
func (s *Secret) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error) {
	if s.Hashed {
		return s.Data, nil
	}

	hasher, ok := dbcontext.GetHasherFromContext(ctx)
	if !ok {
		return nil, errors.New("db context does not contain hasher")
	}
	return hasher.HashBase64(s.Data), nil
}
