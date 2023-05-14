package main

import (
	"fmt"
	"my-app/backend/pkg/database/crud"
	"my-app/backend/pkg/database/engine"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/options"
	"my-app/backend/pkg/snowflake"

	"xorm.io/builder"
)

func main() {
	idGen, err := snowflake.Default()
	if err != nil {
		panic(err)
	}

	engine, err := engine.NewEngine(&options.OEngine{
		DataSource: "./test.db?_pragma=foreign_keys(1)",
		Snowflake:  idGen,
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

	var users []interface{}
	for i := 0; i < 20; i++ {
		test := "test"
		if i%2 == 0 {
			test += "_"
		}
		users = append(users, engine.NewEntity(&entity.User{
			Account:  fmt.Sprint(i) + test,
			Password: fmt.Sprint(i) + test,
		}))
	}
	count, err := engine.Insert(users...)
	if err != nil {
		panic(err)
	}

	println("Inserted", count, "users")

	crud := crud.NewCrud(engine, new(entity.User))
	queryUsers, err := crud.Query(options.NewOCriteria(&options.OCriteria{
		Page: 1,
		Size: 3,
		Sorts: []*options.OCriteriaSort{
			{
				Column: "modified_at",
				Order:  options.OrdDescending,
			},
			{
				Column: "account",
				Order:  options.OrdDescending,
			},
		},
	}), func() (query interface{}, args []interface{}) {
		return builder.Like{"account", "test_"}, nil
	})
	if err != nil {
		panic(err)
	}

	for _, u := range queryUsers {
		println(u.Account)
	}
}
