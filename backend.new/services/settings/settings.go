package settings

import (
	"my-app/backend.new/app"
	"my-app/backend.new/model"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetOptions() (opts map[model.OptionName]string) {
	app.App().UseCfg(func(cfg *app.Config) {
		opts = cfg.Map()
	})
	return
}
