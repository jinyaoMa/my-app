package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type UserPasswordItem struct {
	schema.ModelId
	schema.ModelTime
	Password string `json:"password" doc:"Password"`

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`
}

type UserPasswordDetail struct {
	schema.ModelId
	schema.ModelTime
	Password string `json:"password" doc:"Password"`

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`
}

type UserPasswordSave struct {
	schema.ModelId
	Password string `json:"password" doc:"Password"`

	UserId string `json:"userId" doc:"User Id"`
}
