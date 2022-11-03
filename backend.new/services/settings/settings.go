package settings

import (
	"my-app/backend.new/app"
	"my-app/backend.new/app/types"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetOptions() (opts map[types.ConfigName]string) {
	return app.App().Cfg().OptionPairs()
}
