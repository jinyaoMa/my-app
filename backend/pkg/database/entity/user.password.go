package entity

import iSnowflake "my-app/backend/pkg/snowflake/interfaces"

type UserPassword struct {
	Entity
	PasswordHash string `gorm:"size:64"`
	UserID       int64
}

func NewUserPassword(snowflake iSnowflake.ISnowflake, userPassword *UserPassword) *UserPassword {
	userPassword.snowflake = snowflake
	return userPassword
}
