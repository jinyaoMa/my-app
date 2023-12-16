package app

import (
	"my-app/internal/entity"
	"my-app/pkg/base"
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
	dirCerts, err := base.GetPathStartedFromExecutable("Certs")
	if err != nil {
		panic(err)
	}

	options = []*entity.Option{
		{
			Name:      "",
			Value:     "",
			Encrypted: false,
		},
	}
	return options
}
