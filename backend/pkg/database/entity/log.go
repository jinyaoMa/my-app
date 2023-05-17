package entity

type Log struct {
	Entity
	Tag     string `gorm:"size:3; default:''"`
	Code    int64  `gorm:"default:-1"`
	Message string `gorm:"size:2048; default:''"`
}
