package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type UserPasswordItem struct {
	schema.EntityId
	schema.EntityTime
	Password string `json:"password" doc:"Password"`

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`
}

type UserPasswordDetail struct {
	schema.EntityId
	schema.EntityTime
	Password string `json:"password" doc:"Password"`

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`
}

type UserPasswordSave struct {
	schema.EntityId
	Password string `json:"password" doc:"Password"`

	UserId string `json:"userId" doc:"User Id"`
}
