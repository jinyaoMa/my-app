package entity

type Log struct {
	Entity
	Tag     string `gorm:"size:3"`
	Code    int64  `gorm:""`
	Message string `gorm:"size:2048"`
}
