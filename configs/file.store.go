package configs

type FileStore struct {
	CacheFolderName string
	FragmentSize    uint64 // cache file fragment size
}
