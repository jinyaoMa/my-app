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
	GetOrSaveByOptionName(name string, def string, encrypted ...bool) (value string, opt *entity.Option, err error)

	GetBoolByOptionName(name string) (value bool, opt *entity.Option, err error)
	GetOrSaveBoolByOptionName(name string, def bool, encrypted ...bool) (value bool, opt *entity.Option, err error)

	GetUint16ByOptionName(name string) (value uint16, opt *entity.Option, err error)
	GetOrSaveUint16ByOptionName(name string, def uint16, encrypted ...bool) (value uint16, opt *entity.Option, err error)

	GetStringsByOptionName(name string) (value []string, opt *entity.Option, err error)
	GetOrSaveStringsByOptionName(name string, def []string, encrypted ...bool) (value []string, opt *entity.Option, err error)

	GetColorThemeByOptionName(name string) (value windows.Theme, opt *entity.Option, err error)
	GetOrSaveColorThemeByOptionName(name string, def windows.Theme, encrypted ...bool) (value windows.Theme, opt *entity.Option, err error)

	GetDisplayLanguageByOptionName(name string, availLangs []aio.Lang) (value aio.Lang, opt *entity.Option, err error)
	GetOrSaveDisplayLanguageByOptionName(name string, availLangs []aio.Lang, def string, encrypted ...bool) (value aio.Lang, opt *entity.Option, err error)
	GetDisplayLanguageUsingAvailLangs(availLangs []aio.Lang, lang string) (value aio.Lang)
}

type ICRUDUserPassword interface {
	db.ICRUD[*entity.UserPassword]
}

type ICRUDUser interface {
	db.ICRUD[*entity.User]
}
