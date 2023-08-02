package main

import "my-app/backend/pkg/database/entity"

func (a *App) GetOptions() (opts []*entity.Option) {
	opts, _ = a.optionService.All()
	return
}

func (a *App) SaveOption(opt *entity.Option) (ok bool) {
	_, err := a.optionService.Save(opt)
	return err == nil
}
