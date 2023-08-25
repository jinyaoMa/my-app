package configs

import (
	"my-app/backend/pkg/id"
)

type Database struct {
	LogFile     string
	CipherKey   string
	IdGenerator *id.Config
}
