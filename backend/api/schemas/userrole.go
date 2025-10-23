package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type UserRoleItem struct {
	schema.ModelId
	schema.ModelTime

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`

	IsDefault bool `json:"isDefault" doc:"Is Default Role"`
}

type UserRoleSave struct {
	schema.ModelId

	UserId string `json:"userId" doc:"User Id"`

	RoleId string `json:"roleId" doc:"Role Id"`

	IsDefault bool `json:"isDefault" doc:"Is Default Role"`
}

type UserRoleDetail struct {
	UserRoleSave
	schema.ModelTime

	User UserItem `json:"user" doc:"User"`

	Role RoleItem `json:"role" doc:"Role"`
}

type UserRoleSearchItem struct {
	schema.ModelId

	UserId string         `json:"userId" doc:"User Id"`
	User   UserSearchItem `json:"user" doc:"User"`

	RoleId string         `json:"roleId" doc:"Role Id"`
	Role   RoleSearchItem `json:"role" doc:"Role"`

	IsDefault bool `json:"isDefault" doc:"Is Default Role"`
}
