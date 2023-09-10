package crud

import (
	"errors"
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/aio"
	"my-app/backend/pkg/db"
	"my-app/backend/pkg/funcs"
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"gorm.io/gorm"
)

const (
	// option names
	OptionNameDisplayLanguage = "DisplayLanguage"
	OptionNameColorTheme      = "ColorTheme"

	OptionNameWebAutoStart     = "Web.AutoStart"
	OptionNameWebPortHttp      = "Web.PortHttp"
	OptionNameWebPortHttps     = "Web.PortHttps"
	OptionNameWebDirCerts      = "Web.DirCerts"
	OptionNameWebHostWhitelist = "Web.HostWhitelist"
	OptionNameWebSwagger       = "Web.Swagger"
	OptionNameWebVitePress     = "Web.VitePress"

	// option values
	OptionValueStringsSeparater = "\n"
)

var (
	ErrOptionValueInvalid = errors.New("option.value invalid")
)

type Option struct {
	*db.CRUD[*entity.Option]
}

// GetDisplayLanguageUsingAvailLangs implements interfaces.ICRUDOption.
func (*Option) GetDisplayLanguageUsingAvailLangs(availLangs []aio.Lang, lang string) (value aio.Lang) {
	var found bool
	if value, found = funcs.First(availLangs, func(e aio.Lang) bool {
		return e.Code == lang
	}); !found && len(availLangs) > 0 {
		value = availLangs[0]
	}
	return
}

// GetOrSaveDisplayLanguageByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrSaveDisplayLanguageByOptionName(name string, availLangs []aio.Lang, def string, encrypted ...bool) (value aio.Lang, opt *entity.Option, err error) {
	value, opt, err = o.GetDisplayLanguageByOptionName(name, availLangs)
	v := o.GetDisplayLanguageUsingAvailLangs(availLangs, def)
	if o.trySave(err, opt, name, func() string {
		return v.Code
	}, encrypted...) {
		value = v
		err = nil
	}
	return
}

// GetDisplayLanguageByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetDisplayLanguageByOptionName(name string, availLangs []aio.Lang) (value aio.Lang, opt *entity.Option, err error) {
	_, opt, err = o.GetByOptionName(name)
	if err != nil {
		return
	}
	var found bool
	value, found = funcs.First(availLangs, func(e aio.Lang) bool {
		return e.Code == opt.Value
	})
	if !found {
		err = ErrOptionValueInvalid
	}
	return
}

// GetOrSaveColorThemeByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrSaveColorThemeByOptionName(name string, def windows.Theme, encrypted ...bool) (value windows.Theme, opt *entity.Option, err error) {
	value, opt, err = o.GetColorThemeByOptionName(name)
	if o.trySave(err, opt, name, func() string {
		return strconv.FormatInt(int64(def), 10)
	}, encrypted...) {
		value = def
		err = nil
	}
	return
}

// GetColorThemeByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetColorThemeByOptionName(name string) (value windows.Theme, opt *entity.Option, err error) {
	_, opt, err = o.GetByOptionName(name)
	if err != nil {
		return
	}
	var tmp int64
	tmp, err = strconv.ParseInt(opt.Value, 10, 32)
	if err != nil {
		err = ErrOptionValueInvalid
		return
	}
	value = windows.Theme(tmp)
	switch value {
	case windows.Light:
	case windows.Dark:
	case windows.SystemDefault:
	default:
		err = ErrOptionValueInvalid
	}
	return
}

// GetOrSaveStringsByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrSaveStringsByOptionName(name string, def []string, encrypted ...bool) (value []string, opt *entity.Option, err error) {
	value, opt, err = o.GetStringsByOptionName(name)
	if o.trySave(err, opt, name, func() string {
		return strings.Join(def, OptionValueStringsSeparater)
	}, encrypted...) {
		value = def
		err = nil
	}
	return
}

// GetStringsByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetStringsByOptionName(name string) (value []string, opt *entity.Option, err error) {
	_, opt, err = o.GetByOptionName(name)
	if err != nil {
		return
	}
	value = strings.Split(opt.Value, OptionValueStringsSeparater)
	return
}

// GetOrSaveUint16ByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrSaveUint16ByOptionName(name string, def uint16, encrypted ...bool) (value uint16, opt *entity.Option, err error) {
	value, opt, err = o.GetUint16ByOptionName(name)
	if o.trySave(err, opt, name, func() string {
		return strconv.FormatUint(uint64(def), 10)
	}, encrypted...) {
		value = def
		err = nil
	}
	return
}

// GetUint16ByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetUint16ByOptionName(name string) (value uint16, opt *entity.Option, err error) {
	_, opt, err = o.GetByOptionName(name)
	if err != nil {
		return
	}
	var tmp uint64
	tmp, err = strconv.ParseUint(opt.Value, 10, 16)
	if err != nil {
		err = ErrOptionValueInvalid
		return
	}
	return uint16(tmp), opt, nil
}

// GetOrSaveBoolByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrSaveBoolByOptionName(name string, def bool, encrypted ...bool) (value bool, opt *entity.Option, err error) {
	value, opt, err = o.GetBoolByOptionName(name)
	if o.trySave(err, opt, name, func() string {
		return strconv.FormatBool(def)
	}, encrypted...) {
		value = def
		err = nil
	}
	return
}

// GetBoolByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetBoolByOptionName(name string) (value bool, opt *entity.Option, err error) {
	_, opt, err = o.GetByOptionName(name)
	if err != nil {
		return
	}
	value, err = strconv.ParseBool(opt.Value)
	if err != nil {
		err = ErrOptionValueInvalid
	}
	return
}

// GetOrSaveByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrSaveByOptionName(name string, def string, encrypted ...bool) (value string, opt *entity.Option, err error) {
	value, opt, err = o.GetByOptionName(name)
	if o.trySave(err, opt, name, func() string {
		return def
	}, encrypted...) {
		value = def
		err = nil
	}
	return
}

// GetByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetByOptionName(name string) (value string, opt *entity.Option, err error) {
	opt, err = o.FindOne(func(where func(query any, args ...any)) {
		where(&entity.Option{
			Name: name,
		})
	})
	value = opt.Value
	return
}

// try to save when option not found or its value invalid
func (o *Option) trySave(err error, opt *entity.Option, name string, makeValue func() string, encrypted ...bool) (saved bool) {
	if err == gorm.ErrRecordNotFound {
		*opt = entity.Option{
			Name:  name,
			Value: makeValue(),
			Encrypted: funcs.Any[bool](encrypted, func(e bool) bool {
				return e
			}),
		}
		_, err = o.Save(opt)
		return err == nil
	} else if err == ErrOptionValueInvalid {
		opt.Value = makeValue()
		opt.Encrypted = funcs.Any[bool](encrypted, func(e bool) bool {
			return e
		})
		_, err = o.Save(opt)
		return err == nil
	}
	return
}

func NewOption(dbs *db.DB) *Option {
	return &Option{
		CRUD: db.NewCRUD[*entity.Option](dbs),
	}
}

func NewIOption(dbs *db.DB) interfaces.ICRUDOption {
	return NewOption(dbs)
}
