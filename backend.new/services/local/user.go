package local

import (
	"my-app/backend.new/app"
	"my-app/backend.new/model"
	"my-app/backend.new/utils"
)

const (
	SuperUserAccount = "root"
)

func (s *service) GetSuperUserAccount() string {
	return SuperUserAccount
}

func (s *service) InitializeSuperUser() bool {
	db := app.App().DB()
	admin := &model.User{
		Account: SuperUserAccount,
	}
	if admin.FindByAccount(db) {
		return true
	}
	return admin.Create(db)
}

func (s *service) CheckSuperUserPassword(password string) bool {
	db := app.App().DB()
	admin := &model.User{
		Account: SuperUserAccount,
	}
	return admin.FindByAccount(db) && admin.Password == utils.Utils().SHA1(password)
}

func (s *service) UpdateSuperUserPassword(oldPassword string, newPassword string) bool {
	db := app.App().DB()
	admin := &model.User{
		Account: SuperUserAccount,
	}
	if admin.FindByAccount(db) && admin.Password == utils.Utils().SHA1(oldPassword) {
		admin.Password = newPassword
		admin.PasswordHashed = false
		return admin.Save(db)
	}
	return false
}
