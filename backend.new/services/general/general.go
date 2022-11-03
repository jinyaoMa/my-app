package general

import (
	"my-app/backend.new/app"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// get AppName from i18n translation
func (s *Service) GetAppName() (appname string) {
	return app.App().T().AppName
}
