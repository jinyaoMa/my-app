package service

import (
	"my-app/backend/model"
	"my-app/backend/pkg/utils"
)

const (
	SuperUserAccount = "root"
)

func (s *service) GetSuperUserAccount() string {
	return SuperUserAccount
}

func (s *service) InitializeSuperUser() bool {
	admin := &model.User{
		Account: SuperUserAccount,
	}
	if admin.Find() {
		return true
	}
	return admin.Create()
}

func (s *service) CheckSuperUserPassword(password string) bool {
	admin := &model.User{
		Account: SuperUserAccount,
	}
	return admin.Find() && admin.Password == utils.SHA1(password)
}

func (s *service) UpdateSuperUserPassword(oldPassword string, newPassword string) bool {
	admin := &model.User{
		Account: SuperUserAccount,
	}
	if admin.Find() && admin.Password == utils.SHA1(oldPassword) {
		admin.Password = newPassword
		admin.PasswordHashed = false
		return admin.Save()
	}
	return false
}
