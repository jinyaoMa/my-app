package service

import (
	"fmt"
	"my-app/backend/app"
	"my-app/backend/app/config"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (s *service) GetOptions() *config.Config {
	return app.App().Config()
}

func (s *service) ChooseLogPath(path string, title string) string {
	chosenPath, err := runtime.OpenFileDialog(s.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:           filepath.Dir(path),
		DefaultFilename:            filepath.Base(path),
		Title:                      title,
		ShowHiddenFiles:            true,
		CanCreateDirectories:       true,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		app.App().Log().Service().Fatalf("fail to open file dialog for ChooseLogPath: %+v\n", err)
		return ""
	}

	if chosenPath != "" {
		app.App().Config().Update(config.CfgLogPath, chosenPath)
	}
	return chosenPath
}

func (s *service) ChooseDirCerts(path string, title string) string {
	chosenPath, err := runtime.OpenDirectoryDialog(s.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:           filepath.Dir(path),
		Title:                      title,
		ShowHiddenFiles:            true,
		CanCreateDirectories:       true,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		app.App().Log().Service().Fatalf("fail to open directory dialog for ChooseDirCerts: %+v\n", err)
		return ""
	}

	if chosenPath != "" {
		app.App().Config().Update(config.CfgWebDirCerts, chosenPath)
	}
	return chosenPath
}

func (s *service) SavePortHttp(port uint) (ok bool) {
	return app.App().Config().Update(config.CfgWebPortHttp, fmt.Sprintf(":%d", port))
}

func (s *service) SavePortHttps(port uint) (ok bool) {
	return app.App().Config().Update(config.CfgWebPortHttps, fmt.Sprintf(":%d", port))
}

func (s *service) SaveAutoStart(isStart bool) (ok bool) {
	newValue := "false"
	if isStart {
		newValue = "true"
	}
	return app.App().Config().Update(config.CfgWebAutoStart, newValue)
}
