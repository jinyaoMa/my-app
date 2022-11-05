package local

var _service = &service{}

type service struct{}

func Service() *service {
	return _service
}
