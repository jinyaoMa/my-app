package entity

type Option struct {
	Entity
	Key   string `xorm:"size:256"`
	Value string `xorm:"size:256"`
}
