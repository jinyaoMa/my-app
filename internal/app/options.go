package app

import (
	"my-app/internal/entity"
	"my-app/pkg/db"
)

var (
	options []*entity.Option
)

const (
	OptionNameLocale = "Locale"
	OptionNameTheme  = "Theme"

	OptionNameWebAutoStart = "Web.AutoStart"
)

func initOptions() []*entity.Option {
	options = []*entity.Option{
		{
			Entity: db.Entity[*entity.Option]{
				ID: 1,
			},
			Name:      OptionNameLocale,
			Value:     "",
			Encrypted: false,
		},
		{
			Entity: db.Entity[*entity.Option]{
				ID: 2,
			},
			Name:      OptionNameTheme,
			Value:     "",
			Encrypted: false,
		},
	}
	return options
}
