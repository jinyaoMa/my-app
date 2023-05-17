package entity

type UserPassword struct {
	Entity
	PasswordHash string `gorm:"size:64; notnull; <-:create"`
	UserID       int64
}
