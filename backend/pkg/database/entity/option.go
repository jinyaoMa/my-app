package entity

type Option struct {
	Entity `xorm:"extends"`
	Key    string `xorm:"varchar(255) notnull"`
	Value  string `xorm:"varchar(255) notnull"`
}
