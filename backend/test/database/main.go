package main

import (
	"my-app/backend/pkg/database/engine"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/options"
	"my-app/backend/pkg/snowflake"
)

func main() {
	idGen, err := snowflake.Default()
	if err != nil {
		panic(err)
	}

	engine, err := engine.NewEngine(&options.OEngine{
		Snowflake: idGen,
		Logger: &options.OEngineLogger{
			ShowSQL: true,
		},
	})
	if err != nil {
		panic(err)
	}

	option := engine.NewEntity(&entity.Option{
		Key:   "test",
		Value: "test",
	})
	_, err = engine.Insert(option)
	if err != nil {
		panic(err)
	}

	log := engine.NewEntity(&entity.Log{
		Tag:     "TEST",
		Code:    1,
		Message: "test test test ...",
	})
	_, err = engine.Insert(log)
	if err != nil {
		panic(err)
	}

	user := engine.NewEntity(&entity.User{
		Account:  "test",
		Password: "test",
	})
	_, err = engine.Insert(user)
	if err != nil {
		panic(err)
	}

	println("Inserted")
}
