package crud

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
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
			SrcType: datatype.Id(0),
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
			DstType: datatype.Id(0),
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
