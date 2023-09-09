package interfaces

import (
	"my-app/backend/internal/entity"
	"my-app/backend/pkg/aio"
	"my-app/backend/pkg/db"

	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

type ICRUDFile interface {
	db.ICRUD[*entity.File]
}

type ICRUDFileCategory interface {
	db.ICRUD[*entity.FileCategory]
}

type ICRUDFileExtension interface {
	db.ICRUD[*entity.FileExtension]
}

type ICRUDLog interface {
	db.ICRUD[*entity.Log]
}

type ICRUDNode interface {
	db.ICRUD[*entity.Node]
}

type ICRUDOption interface {
	db.ICRUD[*entity.Option]

	GetByOptionName(name string) (value string, opt *entity.Option, err error)
	GetOrCreateByOptionName(name string, def string, encrypted ...bool) (value string, opt *entity.Option, err error)

	GetBoolByOptionName(name string) (value bool, opt *entity.Option, err error)
	GetOrCreateBoolByOptionName(name string, def bool, encrypted ...bool) (value bool, opt *entity.Option, err error)

	GetUint16ByOptionName(name string) (value uint16, opt *entity.Option, err error)
	GetOrCreateUint16ByOptionName(name string, def uint16, encrypted ...bool) (value uint16, opt *entity.Option, err error)

	GetStringsByOptionName(name string) (value []string, opt *entity.Option, err error)
	GetOrCreateStringsByOptionName(name string, def []string, encrypted ...bool) (value []string, opt *entity.Option, err error)

	GetColorThemeByOptionName(name string) (value windows.Theme, opt *entity.Option, err error)
	SaveColorThemeByOptionName(name string, def windows.Theme, encrypted ...bool) (value windows.Theme, opt *entity.Option, err error)
	GetOrCreateColorThemeByOptionName(name string, def windows.Theme, encrypted ...bool) (value windows.Theme, opt *entity.Option, err error)

	GetDisplayLanguageByOptionName(name string, availLangs []aio.Lang) (value *aio.Lang, opt *entity.Option, err error)
	SaveDisplayLanguageByOptionName(name string, availLangs []aio.Lang, lang string, encrypted ...bool) (value *aio.Lang, opt *entity.Option, err error)
	GetOrCreateDisplayLanguageByOptionName(name string, availLangs []aio.Lang, def string, encrypted ...bool) (value *aio.Lang, opt *entity.Option, err error)
}

type ICRUDUserPassword interface {
	db.ICRUD[*entity.UserPassword]
}

type ICRUDUser interface {
	db.ICRUD[*entity.User]
}
