package wails

import "my-app/backend.new/web"

type binding struct{}

func NewBinding() *binding {
	return &binding{}
}

func (*binding) IsWebServiceRunning() bool {
	return web.Web().IsRunning()
}
