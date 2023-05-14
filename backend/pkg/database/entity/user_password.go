package entity

type UserPassword struct {
	Entity       `xorm:"extends"`
	UserId       int64
	PasswordHash string `xorm:"varchar(64) notnull"`
}
