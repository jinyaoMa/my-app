package storage

type Storage struct {
	paths []string
}

// AddPath implements Interface
func (*Storage) AddPath() {
	panic("unimplemented")
}

func New(paths ...string) Interface {
	return &Storage{
		paths: paths,
	}
}
