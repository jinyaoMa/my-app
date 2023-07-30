package app

import (
	"my-app/backend/configs"
	"my-app/backend/pkg/crypto"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/snowflake"
)

func initDB(cfg *configs.Configs) (db *database.Database, err error) {
	entity.Cipher(crypto.NewAesWithSalt(cfg.Database.CipherKey))

	idGen, err := snowflake.New(cfg.Database.Snowflake)
	if err != nil {
		return
	}
	entity.IdGenerator(idGen)

	return
}
