package service

import (
	"my-app/backend/app"
	"my-app/backend/model"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
		case app.CfgWebAutoStart:
			app.App().Config().Web.AutoStart = value
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
	case app.CfgWebAutoStart:
		return app.App().Config().Web.AutoStart
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

func (s *settings) ChooseLogPath(path string, title string) string {
	chosenPath, err := runtime.OpenFileDialog(app.App().WailsContext(), runtime.OpenDialogOptions{
		DefaultDirectory:           filepath.Dir(path),
		DefaultFilename:            filepath.Base(path),
		Title:                      title,
		ShowHiddenFiles:            true,
		CanCreateDirectories:       true,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		app.App().ServiceLog().Fatalf("fail to open file dialog for ChooseLogPath: %+v\n", err)
		return ""
	}

	if chosenPath != "" {
		err = s.SaveOption(app.CfgLogPath, chosenPath)
		if err != nil {
			app.App().ServiceLog().Fatalf("fail to update CfgLogPath: %+v\n", err)
			return ""
		}
	}
	return chosenPath
}

func (s *settings) ChooseDirCerts(path string, title string) string {
	chosenPath, err := runtime.OpenDirectoryDialog(app.App().WailsContext(), runtime.OpenDialogOptions{
		DefaultDirectory:           filepath.Dir(path),
		Title:                      title,
		ShowHiddenFiles:            true,
		CanCreateDirectories:       true,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		app.App().ServiceLog().Fatalf("fail to open directory dialog for ChooseDirCerts: %+v\n", err)
		return ""
	}

	if chosenPath != "" {
		err = s.SaveOption(app.CfgWebDirCerts, chosenPath)
		if err != nil {
			app.App().ServiceLog().Fatalf("fail to update CfgWebDirCerts: %+v\n", err)
			return ""
		}
	}
	return chosenPath
}

func (s *settings) SavePortHttp(port string) bool {
	err := s.SaveOption(app.CfgWebPortHttp, port)
	if err != nil {
		app.App().ServiceLog().Fatalf("fail to update CfgWebPortHttp: %+v\n", err)
		return false
	}
	return true
}

func (s *settings) SavePortHttps(port string) bool {
	err := s.SaveOption(app.CfgWebPortHttps, port)
	if err != nil {
		app.App().ServiceLog().Fatalf("fail to update CfgWebPortHttps: %+v\n", err)
		return false
	}
	return true
}

func (s *settings) SaveAutoStart(isStart string) bool {
	err := s.SaveOption(app.CfgWebAutoStart, isStart)
	if err != nil {
		app.App().ServiceLog().Fatalf("fail to update CfgWebAutoStart: %+v\n", err)
		return false
	}
	return true
}
