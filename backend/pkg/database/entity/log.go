package entity

type Log struct {
	Entity  `xorm:"extends"`
	Tag     string `xorm:"varchar(3) notnull"`
	Code    int64  `xorm:"notnull"`
	Message string `xorm:"varchar(2048) notnull"`
}
