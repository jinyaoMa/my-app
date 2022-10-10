package config

// web option names
const (
	CfgWebAutoStart = "Web.AutoStart"
	CfgWebPortHttp  = "Web.PortHttp"
	CfgWebPortHttps = "Web.PortHttps"
	CfgWebDirCerts  = "Web.DirCerts"
)

type WebConfig struct {
	AutoStart string
	PortHttp  string
	PortHttps string
	DirCerts  string
}

func (w *WebConfig) IsAutoStart() bool {
	return w.AutoStart == "true"
}
