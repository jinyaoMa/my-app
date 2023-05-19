package entity

type UserPassword struct {
	Entity

	/* internal fields */
	PasswordHash string `gorm:"size:64; notnull; <-:create"`

	/* relational fields */
	UserID int64 `gorm:""`
}
