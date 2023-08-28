package main

import "my-app/backend/internal/entity"

func (a *App) GetOptions() (opts []*entity.Option) {
	opts, _ = a.crudOption.All()
	return
}

func (a *App) SaveOption(opt *entity.Option) (ok bool) {
	_, err := a.crudOption.Save(opt)
	return err == nil
}
