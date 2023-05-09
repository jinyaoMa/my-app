package main

import (
	"my-app/backend/pkg/database/engine"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/snowflake"
)

func main() {
	idGen, err := snowflake.Default()
	if err != nil {
		panic(err)
	}

	engine, err := engine.New(&engine.Options{
		Driver:     engine.DrvSQLite3,
		DataSource: "./test.db",
		Snowflake:  idGen,
	})
	if err != nil {
		panic(err)
	}

	user := engine.NewEntity(&entity.User{
		Account:  "test",
		Password: "test",
	})
	count, err := engine.Insert(user)
	if err != nil {
		panic(err)
	}

	println("Insert:", count)
}
