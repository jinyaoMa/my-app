package types

// config names
const (
	ConfigNameDisplayLanguage = ConfigName("DisplayLanguage")
	ConfigNameColorTheme      = ConfigName("ColorTheme")

	ConfigNameLogFile      = ConfigName("LogFile")
	ConfigNameDirLanguages = ConfigName("DirLanguages")
	ConfigNameDirAssets    = ConfigName("DirAssets")
	ConfigNameDirUserData  = ConfigName("DirUserData")
	ConfigNameDirDocs      = ConfigName("DirDocs")

	ConfigNameWebAutoStart = ConfigName("WebAutoStart")
	ConfigNameWebPortHttp  = ConfigName("WebPortHttp")
	ConfigNameWebPortHttps = ConfigName("WebPortHttps")
	ConfigNameWebDirCerts  = ConfigName("WebDirCerts")
)

type ConfigName string

func (cn ConfigName) ToString() string {
	return string(cn)
}
