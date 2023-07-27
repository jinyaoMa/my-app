package assetio

type Interface interface {
	GetBytes(paths ...string) (data []byte)
}
