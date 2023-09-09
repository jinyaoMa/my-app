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

	OptionValueColorThemeSystem = "system"
	OptionValueColorThemeLight  = "light"
	OptionValueColorThemeDark   = "dark"
)

var (
	ErrOptionValueColorThemeInvalid = errors.New("option.value[color theme] invalid")
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

// GetOrCreateDisplayLanguageByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrCreateDisplayLanguageByOptionName(name string, availLangs []aio.Lang, def string, encrypted ...bool) (value aio.Lang, opt *entity.Option, err error) {
	value, opt, err = o.GetDisplayLanguageByOptionName(name, availLangs)
	if err == gorm.ErrRecordNotFound {
		value = o.GetDisplayLanguageUsingAvailLangs(availLangs, def)
		opt = &entity.Option{
			Name:      name,
			Value:     value.Code,
			Encrypted: len(encrypted) > 0 && encrypted[0],
		}
		_, err = o.Save(opt)
		return
	}
	return
}

// GetDisplayLanguageByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetDisplayLanguageByOptionName(name string, availLangs []aio.Lang) (value aio.Lang, opt *entity.Option, err error) {
	_, opt, err = o.GetByOptionName(name)
	if err != nil {
		return
	}
	value, _ = funcs.First(availLangs, func(e aio.Lang) bool {
		return e.Code == opt.Value
	})
	return
}

// GetColorThemeUsingWindowsTheme implements interfaces.ICRUDOption.
func (*Option) GetColorThemeUsingWindowsTheme(theme windows.Theme) (value string) {
	switch theme {
	case windows.Light:
		value = OptionValueColorThemeLight
	case windows.Dark:
		value = OptionValueColorThemeDark
	// case windows.SystemDefault:
	default:
		value = OptionValueColorThemeSystem
	}
	return
}

// GetOrCreateColorThemeByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrCreateColorThemeByOptionName(name string, def windows.Theme, encrypted ...bool) (value windows.Theme, opt *entity.Option, err error) {
	value, opt, err = o.GetColorThemeByOptionName(name)
	if err == gorm.ErrRecordNotFound || err == ErrOptionValueColorThemeInvalid {
		opt = &entity.Option{
			Name:  name,
			Value: o.GetColorThemeUsingWindowsTheme(def),
		}
		if len(encrypted) > 0 && encrypted[0] {
			opt.Encrypted = true
		}
		_, err = o.Save(opt)
		if err == nil {
			value = def
		}
		return
	}
	return
}

// GetColorThemeByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetColorThemeByOptionName(name string) (value windows.Theme, opt *entity.Option, err error) {
	_, opt, err = o.GetByOptionName(name)
	if err != nil {
		return
	}
	switch opt.Value {
	case OptionValueColorThemeLight:
		value = windows.Light
	case OptionValueColorThemeDark:
		value = windows.Dark
	case OptionValueColorThemeSystem:
		value = windows.SystemDefault
	default:
		err = ErrOptionValueColorThemeInvalid
	}
	return
}

// GetOrCreateStringsByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrCreateStringsByOptionName(name string, def []string, encrypted ...bool) (value []string, opt *entity.Option, err error) {
	value, opt, err = o.GetStringsByOptionName(name)
	if err == gorm.ErrRecordNotFound {
		opt = &entity.Option{
			Name:  name,
			Value: strings.Join(def, OptionValueStringsSeparater),
		}
		if len(encrypted) > 0 && encrypted[0] {
			opt.Encrypted = true
		}
		_, err = o.Save(opt)
		if err == nil {
			value = def
		}
		return
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

// GetOrCreateUint16ByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrCreateUint16ByOptionName(name string, def uint16, encrypted ...bool) (value uint16, opt *entity.Option, err error) {
	value, opt, err = o.GetUint16ByOptionName(name)
	if err == gorm.ErrRecordNotFound {
		opt = &entity.Option{
			Name:  name,
			Value: strconv.FormatUint(uint64(def), 10),
		}
		if len(encrypted) > 0 && encrypted[0] {
			opt.Encrypted = true
		}
		_, err = o.Save(opt)
		if err == nil {
			value = def
		}
		return
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
		return
	}
	return uint16(tmp), opt, nil
}

// GetOrCreateBoolByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrCreateBoolByOptionName(name string, def bool, encrypted ...bool) (value bool, opt *entity.Option, err error) {
	value, opt, err = o.GetBoolByOptionName(name)
	if err == gorm.ErrRecordNotFound {
		opt = &entity.Option{
			Name:  name,
			Value: strconv.FormatBool(def),
		}
		if len(encrypted) > 0 && encrypted[0] {
			opt.Encrypted = true
		}
		_, err = o.Save(opt)
		if err == nil {
			value = def
		}
		return
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
	return
}

// GetOrCreateByOptionName implements interfaces.ICRUDOption.
func (o *Option) GetOrCreateByOptionName(name string, def string, encrypted ...bool) (value string, opt *entity.Option, err error) {
	value, opt, err = o.GetByOptionName(name)
	if err == gorm.ErrRecordNotFound {
		opt = &entity.Option{
			Name:  name,
			Value: def,
		}
		if len(encrypted) > 0 && encrypted[0] {
			opt.Encrypted = true
		}
		_, err = o.Save(opt)
		return
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

func NewOption(dbs *db.DB) *Option {
	return &Option{
		CRUD: db.NewCRUD[*entity.Option](dbs),
	}
}

func NewIOption(dbs *db.DB) interfaces.ICRUDOption {
	return NewOption(dbs)
}
