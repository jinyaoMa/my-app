package view

import "majinyao.cn/my-app/backend/pkg/db/datatype"

type UserData struct {
	VisitorId string
	UserId    int64
	RoleId    int64
}

func (u UserData) GetIdentity() string {
	return u.GetUserId().B36String() + "_" + u.VisitorId
}

func (u UserData) GetUserId() datatype.Id {
	return datatype.Id(u.UserId)
}

func (u UserData) GetRoleId() datatype.Id {
	return datatype.Id(u.RoleId)
}

type CurrentUser struct {
	Id     string            `json:"id" doc:"User Id (Base36)"`
	Name   string            `json:"name" doc:"Name"`
	Roles  []CurrentUserRole `json:"roles" doc:"Available Roles"`
	Avatar string            `json:"avatar" doc:"Avatar (DataURL)"`
}

type CurrentUserRole struct {
	Id        string `json:"id" doc:"Role Id (Base36)"`
	Code      string `json:"code" doc:"Code"`
	Name      string `json:"name" doc:"Name"`
	IsCurrent bool   `json:"isCurrent" doc:"Is Current Role"`
}

type CreateReservedUserForm struct {
	Name     string
	Account  string
	Password string
}

type ReservedUserInfo struct {
	Id      string
	Name    string
	Account string
	Avatar  string
}
