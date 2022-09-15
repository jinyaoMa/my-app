package web

const (
	CfgPortHttp  = Package + ".PortHttp"
	CfgPortHttps = Package + ".PortHttps"
	CfgDirCerts  = Package + ".DirCerts"
)

type Config struct {
	PortHttp  string
	PortHttps string
	DirCerts  string
}

func DefaultConfig() Config {
	return Config{
		PortHttp:  ":10080",
		PortHttps: ":10443",
		DirCerts:  "",
	}
}
