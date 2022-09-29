package service

import (
	"my-app/backend/app"
	"my-app/backend/model"
)

type settings struct{}

func (s *settings) SaveOption(name string, value string) error {
	option := model.MyOption{
		Name: name,
	}
	result := option.Update(value)
	if result.Error == nil {
		switch name {
		case app.CfgDisplayLanguage:
			app.App().Config().DisplayLanguage = value
		case app.CfgColorTheme:
			app.App().Config().ColorTheme = value
		case app.CfgLogPath:
			app.App().Config().LogPath = value
		case app.CfgWebPortHttp:
			app.App().Config().Web.PortHttp = value
		case app.CfgWebPortHttps:
			app.App().Config().Web.PortHttps = value
		case app.CfgWebDirCerts:
			app.App().Config().Web.DirCerts = value
		}
	}
	return result.Error
}

func (s *settings) GetOption(name string) string {
	switch name {
	case app.CfgDisplayLanguage:
		return app.App().Config().DisplayLanguage
	case app.CfgColorTheme:
		return app.App().Config().ColorTheme
	case app.CfgLogPath:
		return app.App().Config().LogPath
	case app.CfgWebPortHttp:
		return app.App().Config().Web.PortHttp
	case app.CfgWebPortHttps:
		return app.App().Config().Web.PortHttps
	case app.CfgWebDirCerts:
		return app.App().Config().Web.DirCerts
	}
	return ""
}

func (s *settings) GetOptions() *app.Config {
	return app.App().Config()
}
