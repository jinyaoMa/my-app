package entity

import (
	"fmt"
	"strconv"

	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

const (
	OptionKeySystemLocale     = "system.locale"
	OptionKeySystemColorTheme = "system.color.theme"
	OptionKeyServerAutoRun    = "server.auto.run"
	OptionKeyServerPort       = "server.port"
	OptionKeyServerSecurePort = "server.secure.port"
	OptionKeyServerCertFile   = "server.cert.file"
	OptionKeyServerKeyFile    = "server.key.file"
)

const (
	OptionColorThemeLight  = "light"
	OptionColorThemeDark   = "dark"
	OptionColorThemeSystem = "system"
)

func MustNewOption(key string, value any) *Option {
	o, err := NewOption(key, value)
	if err != nil {
		panic(err)
	}
	return o
}

func NewOption(key string, value any) (*Option, error) {
	o := &Option{
		Key: key,
	}
	switch t := value.(type) {
	case string:
		o.SetString(value.(string))
	case bool:
		o.SetBool(value.(bool))
	case int64:
		o.SetInt64(value.(int64))
	case uint64:
		o.SetUint64(value.(uint64))
	case float64:
		o.SetFloat64(value.(float64))
	default:
		return nil, fmt.Errorf("unsupported type %T", t)
	}
	return o, nil
}

type Option struct {
	db.Entity
	db.EntityReserved
	Key   string             `gorm:"<-:create;index;not null;size:254;comment:Option Key;"`
	Value datatype.Encrypted `gorm:"size:254;comment:Option Value;"`
}

func (o *Option) GetString() (v string) {
	return string(o.Value)
}

func (o *Option) SetString(v string) {
	o.Value = datatype.Encrypted(v)
}

func (o *Option) GetColorTheme() (v string) {
	switch o.Value {
	case OptionColorThemeLight:
		return OptionColorThemeLight
	case OptionColorThemeDark:
		return OptionColorThemeDark
	}
	return OptionColorThemeSystem
}

func (o *Option) SetColorTheme(v string) {
	switch v {
	case OptionColorThemeLight:
		o.SetString(OptionColorThemeLight)
	case OptionColorThemeDark:
		o.SetString(OptionColorThemeDark)
	}
	o.SetString(OptionColorThemeSystem)
}

func (o *Option) GetBool() (v bool) {
	return o.GetString() == "true"
}

func (o *Option) SetBool(v bool) {
	switch v {
	case true:
		o.SetString("true")
	default:
		o.SetString("false")
	}
}

func (o *Option) GetInt64() (v int64) {
	v, err := strconv.ParseInt(o.GetString(), 36, 64)
	if err != nil {
		return v
	}
	return v
}

func (o *Option) SetInt64(v int64) {
	o.SetString(strconv.FormatInt(v, 36))
}

func (o *Option) GetUint16() (v uint16) {
	v64, err := strconv.ParseUint(o.GetString(), 36, 16)
	if err != nil {
		return v
	}
	return uint16(v64)
}

func (o *Option) SetUint16(v uint16) {
	o.SetString(strconv.FormatUint(uint64(v), 36))
}

func (o *Option) GetUint64() (v uint64) {
	v, err := strconv.ParseUint(o.GetString(), 36, 64)
	if err != nil {
		return v
	}
	return v
}

func (o *Option) SetUint64(v uint64) {
	o.SetString(strconv.FormatUint(v, 36))
}

func (o *Option) GetFloat64() (v float64) {
	v, err := strconv.ParseFloat(o.GetString(), 64)
	if err != nil {
		return v
	}
	return v
}

func (o *Option) SetFloat64(v float64) {
	o.SetString(strconv.FormatFloat(v, 'f', -1, 64))
}
