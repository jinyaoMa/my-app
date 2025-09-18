package db

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var DefaultCopierOption = MakeCopiterOption()

func MakeCopiterOption(converters ...copier.TypeConverter) copier.Option {
	return copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters:  append(TypeConverters(), converters...),
	}
}

func TypeConverters() []copier.TypeConverter {
	return []copier.TypeConverter{
		{
			SrcType: int64(0),
			DstType: copier.String,
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(int64)
				if !ok {
					return dst, errors.New("copier src type [int64] not matched dst type [string]")
				}

				return ConvertIdToString(s), nil
			},
		},
		{
			SrcType: copier.String,
			DstType: int64(0),
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(string)
				if !ok {
					return nil, errors.New("copier src type [string] not matched dst type [int64]")
				}

				dst, err = ConvertStringToId(s)
				if err != nil {
					return nil, err
				}
				return dst, nil
			},
		},
		{
			SrcType: new(time.Time),
			DstType: gorm.DeletedAt{},
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(*time.Time)
				if !ok {
					return nil, errors.New("copier src type [*time.Time] not matched dst type [gorm.DeletedAt]")
				}

				if s == nil {
					return gorm.DeletedAt{}, nil
				}
				return gorm.DeletedAt{
					Time:  *s,
					Valid: true,
				}, nil
			},
		},
		{
			SrcType: gorm.DeletedAt{},
			DstType: new(time.Time),
			Fn: func(src any) (dst any, err error) {
				s, ok := src.(gorm.DeletedAt)
				if !ok {
					return nil, errors.New("copier src type [gorm.DeletedAt] not matched dst type [*time.Time]")
				}

				if s.Valid {
					return s.Time, nil
				}
				return nil, nil
			},
		},
	}
}
