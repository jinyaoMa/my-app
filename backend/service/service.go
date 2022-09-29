package service

var (
	instance *service
)

type service struct {
	Settings *settings
}

func init() {
	instance = &service{
		Settings: &settings{},
	}
}

func Settings() *settings {
	return instance.Settings
}
