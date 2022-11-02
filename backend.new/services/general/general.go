package general

import (
	"my-app/backend.new/app"
	"my-app/backend.new/app/i18n"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// get AppName from i18n translation
func (s *Service) GetAppName() (appname string) {
	app.App().UseI18n(func(T func() *i18n.Translation, i18n *i18n.I18n) {
		appname = T().AppName
	})
	return
}
