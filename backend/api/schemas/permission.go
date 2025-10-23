package schemas

import (
	"database/sql"

	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/flag"
)

type PermissionItem struct {
	schema.ModelId
	schema.ModelTime
	schema.ModelReserved
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}

type PermissionSave struct {
	schema.ModelId
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
	FlagBase64  string `json:"flagBase64" doc:"Flag (Base64)"`
}

func (s *PermissionSave) Flag(bs sql.RawBytes) {
	s.FlagBase64 = flag.FromBytes(bs).ToBase64()
}

type PermissionDetail struct {
  PermissionSave
	schema.ModelTime
	schema.ModelReserved

	RolePermissions []RolePermissionItem `json:"rolePermissions" doc:"Role Permissions"`
	Roles           []RoleItem           `json:"roles" doc:"Roles"`
}

type PermissionSearchItem struct {
	schema.ModelId
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`
}
