package crud

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

var GlobalCopierOption = MakeCopiterOption()

func SetGlobalCopierOption(o copier.Option) {
	GlobalCopierOption = o
}

func MakeCopiterOption(converters ...copier.TypeConverter) copier.Option {
	return copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters:  CopierSqlTypeConverters(CopierTypeConverters(converters...)...),
	}
}

func CopierSqlTypeConverters(converters ...copier.TypeConverter) []copier.TypeConverter {
	var sqlNullBool sql.NullBool
	var sqlNullByte sql.NullByte
	var sqlNullFloat64 sql.NullFloat64
	var sqlNullInt16 sql.NullInt16
	var sqlNullInt32 sql.NullInt32
	var sqlNullInt64 sql.NullInt64
	var sqlNullString sql.NullString
	var sqlNullTime sql.NullTime
	var deletedAt gorm.DeletedAt

	var zeroBool bool
	var zeroByte byte
	var zeroFloat64 float64
	var zeroInt16 int16
	var zeroInt32 int32
	var zeroInt64 int64
	var zeroString string
	var zeroTime time.Time

	return append([]copier.TypeConverter{
		{
			SrcType: sqlNullBool,
			DstType: zeroBool,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(sql.NullBool)
				if !ok {
					return dst, errors.New("copier src type [sql.NullBool] not matched dst type [bool]")
				}
				if s.Valid {
					return s.Bool, nil
				}
				return zeroBool, nil
			},
		},
		{
			SrcType: sqlNullByte,
			DstType: zeroByte,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(sql.NullByte)
				if !ok {
					return dst, errors.New("copier src type [sql.NullByte] not matched dst type [byte]")
				}
				if s.Valid {
					return s.Byte, nil
				}
				return zeroByte, nil
			},
		},
		{
			SrcType: sqlNullFloat64,
			DstType: zeroFloat64,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(sql.NullFloat64)
				if !ok {
					return dst, errors.New("copier src type [sql.NullFloat64] not matched dst type [float64]")
				}
				if s.Valid {
					return s.Float64, nil
				}
				return zeroFloat64, nil
			},
		},
		{
			SrcType: sqlNullInt16,
			DstType: zeroInt16,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(sql.NullInt16)
				if !ok {
					return dst, errors.New("copier src type [sql.NullInt16] not matched dst type [int16]")
				}
				if s.Valid {
					return s.Int16, nil
				}
				return zeroInt16, nil
			},
		},
		{
			SrcType: sqlNullInt32,
			DstType: zeroInt32,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(sql.NullInt32)
				if !ok {
					return dst, errors.New("copier src type [sql.NullInt32] not matched dst type [int32]")
				}
				if s.Valid {
					return s.Int32, nil
				}
				return zeroInt32, nil
			},
		},
		{
			SrcType: sqlNullInt64,
			DstType: zeroInt64,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(sql.NullInt64)
				if !ok {
					return dst, errors.New("copier src type [sql.NullInt64] not matched dst type [int64]")
				}
				if s.Valid {
					return s.Int64, nil
				}
				return zeroInt64, nil
			},
		},
		{
			SrcType: sqlNullString,
			DstType: zeroString,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(sql.NullString)
				if !ok {
					return dst, errors.New("copier src type [sql.NullString] not matched dst type [string]")
				}
				if s.Valid {
					return s.String, nil
				}
				return zeroString, nil
			},
		},
		{
			SrcType: sqlNullTime,
			DstType: zeroTime,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(sql.NullTime)
				if !ok {
					return dst, errors.New("copier src type [sql.NullTime] not matched dst type [time.Time]")
				}
				if s.Valid {
					return s.Time, nil
				}
				return zeroTime, nil
			},
		},
		{
			SrcType: deletedAt,
			DstType: zeroTime,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(gorm.DeletedAt)
				if !ok {
					return dst, errors.New("copier src type [gorm.DeletedAt] not matched dst type [time.Time]")
				}
				if s.Valid {
					return s.Time, nil
				}
				return zeroTime, nil
			},
		},
	}, converters...)
}

func CopierTypeConverters(converters ...copier.TypeConverter) []copier.TypeConverter {
	id := datatype.Id(0)
	oid := make(datatype.Oid, 0)
	password := datatype.Password{}
	secret := datatype.Secret{}

	return append([]copier.TypeConverter{
		{
			SrcType: id,
			DstType: copier.String,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(datatype.Id)
				if !ok {
					return dst, errors.New("copier src type [datatype.Id] not matched dst type [string]")
				}
				return s.B36String(), nil
			},
		},
		{
			SrcType: copier.String,
			DstType: id,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(string)
				if !ok {
					return nil, errors.New("copier src type [string] not matched dst type [datatype.Id]")
				}
				dst, err = datatype.ParseIdFromB36(s)
				if err != nil {
					return nil, err
				}
				return dst, nil
			},
		},
		{
			SrcType: oid,
			DstType: copier.String,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(datatype.Oid)
				if !ok {
					return nil, errors.New("copier src type [datatype.Oid] not matched dst type [string]")
				}
				return s.String(), nil
			},
		},
		{
			SrcType: copier.String,
			DstType: oid,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(string)
				if !ok {
					return nil, errors.New("copier src type [string] not matched dst type [datatype.Oid]")
				}
				oid, err := datatype.ParseOid(s)
				if err != nil {
					return nil, err
				}
				return oid, nil
			},
		},
		{
			SrcType: password,
			DstType: copier.String,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(datatype.Password)
				if !ok {
					return nil, errors.New("copier src type [datatype.Password] not matched dst type [string]")
				}
				return s.Data, nil
			},
		},
		{
			SrcType: copier.String,
			DstType: password,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(string)
				if !ok {
					return nil, errors.New("copier src type [string] not matched dst type [datatype.Password]")
				}
				return datatype.Password{
					Data: s,
				}, nil
			},
		},
		{
			SrcType: secret,
			DstType: copier.String,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(datatype.Secret)
				if !ok {
					return nil, errors.New("copier src type [datatype.Secret] not matched dst type [string]")
				}
				return s.Data, nil
			},
		},
		{
			SrcType: copier.String,
			DstType: secret,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(string)
				if !ok {
					return nil, errors.New("copier src type [string] not matched dst type [datatype.Secret]")
				}
				return datatype.Secret{
					Data: s,
				}, nil
			},
		},
	}, converters...)
}
