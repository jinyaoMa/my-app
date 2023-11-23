package fstore

type Options struct {
	CacheFolderName string
}

var (
	DefaultOptions = &Options{
		CacheFolderName: ".cache",
	}
)
