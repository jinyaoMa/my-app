package local

import (
	"my-app/backend/app"
	"my-app/backend/app/types"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (s *service) GetOptions() map[string]string {
	opts := make(map[string]string)
	for k, v := range app.App().Cfg().OptionPairs() {
		opts[k.ToString()] = v
	}
	return opts
}

func (s *service) ChooseLogFile(path string, title string) string {
	chosenPath, err := runtime.OpenFileDialog(app.App().Ctx(), runtime.OpenDialogOptions{
		DefaultDirectory:           filepath.Dir(path),
		DefaultFilename:            filepath.Base(path),
		Title:                      title,
		ShowHiddenFiles:            true,
		CanCreateDirectories:       true,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		app.App().Log().Services().Fatalf("fail to open file dialog for ChooseLogFile: %+v\n", err)
		return ""
	}

	if chosenPath != "" && app.App().Cfg().Set(types.ConfigNameLogFile, chosenPath) {
		return chosenPath
	}
	return ""
}

func (s *service) ChooseDirectory(which string, path string, title string) string {
	var cfgName types.ConfigName
	switch which {
	case "languages":
		cfgName = types.ConfigNameDirLanguages
	case "assets":
		cfgName = types.ConfigNameDirAssets
	case "userdata":
		cfgName = types.ConfigNameDirUserData
	case "docs":
		cfgName = types.ConfigNameDirDocs
	case "certs":
		cfgName = types.ConfigNameWebDirCerts
	}
	if cfgName == "" {
		app.App().Log().Services().Fatalf(
			"invalid arguments for ChooseDirectory: which=%s, path=%s, title=%s\n",
			which, path, title,
		)
		return ""
	}
	chosenPath, err := runtime.OpenDirectoryDialog(app.App().Ctx(), runtime.OpenDialogOptions{
		DefaultDirectory:           filepath.Dir(path),
		Title:                      title,
		ShowHiddenFiles:            true,
		CanCreateDirectories:       true,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		app.App().Log().Services().Fatalf("fail to open directory dialog for ChooseDirectory (%s): %+v\n", cfgName, err)
		return ""
	}

	if chosenPath != "" && app.App().Cfg().Set(cfgName, chosenPath) {
		return chosenPath
	}
	return ""
}

func (s *service) SavePortHttp(port uint) (ok bool) {
	return app.App().Cfg().Set(types.ConfigNameWebPortHttp, types.Port(port).ToString())
}

func (s *service) SavePortHttps(port uint) (ok bool) {
	return app.App().Cfg().Set(types.ConfigNameWebPortHttps, types.Port(port).ToString())
}

func (s *service) SaveAutoStart(isStart bool) (ok bool) {
	return app.App().Cfg().Set(types.ConfigNameWebAutoStart, types.Boolean(isStart).ToString())
}
