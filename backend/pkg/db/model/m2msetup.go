package model

type M2MSetup struct {
	Model     any
	Field     string
	JoinTable any
}

type M2MSetupsGetter interface {
	GetM2MSetups() []M2MSetup
}
