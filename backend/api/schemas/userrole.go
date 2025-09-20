package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type UserRoleItem struct {
	schema.ModelId
	schema.ModelTime

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`
}

type UserRoleDetail struct {
	schema.ModelId
	schema.ModelTime

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`
}

type UserRoleSave struct {
	schema.ModelId

	UserId string `json:"userId" doc:"User Id"`

	RoleId string `json:"roleId" doc:"Role Id"`
}
