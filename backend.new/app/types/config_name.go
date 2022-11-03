package types

// config names
const (
	ConfigNameDisplayLanguage = ConfigName("DisplayLanguage")
	ConfigNameColorTheme      = ConfigName("ColorTheme")

	ConfigNameFileLog      = ConfigName("FileLog")
	ConfigNameDirLanguages = ConfigName("DirLanguages")
	ConfigNameDirAssets    = ConfigName("DirAssets")
	ConfigNameDirUserData  = ConfigName("DirUserData")
	ConfigNameDirDocs      = ConfigName("DirDocs")

	ConfigNameWebAutoStart = ConfigName("Web.AutoStart")
	ConfigNameWebPortHttp  = ConfigName("Web.PortHttp")
	ConfigNameWebPortHttps = ConfigName("Web.PortHttps")
	ConfigNameWebDirCerts  = ConfigName("Web.DirCerts")
)

type ConfigName string

func (cn ConfigName) ToString() string {
	return string(cn)
}
