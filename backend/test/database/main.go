package main

import (
	"fmt"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/crud"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/options"
	optionsLogger "my-app/backend/pkg/logger/options"
	"my-app/backend/pkg/snowflake"
	"my-app/backend/pkg/utility"

	"gorm.io/driver/sqlite"
)

func main() {
	util, err := utility.NewUtility()
	if err != nil {
		panic(err)
	}

	entity.SetAes(utility.NewAesWithSalt(util.GetExecutableFileName("option.key")))

	idGen, err := snowflake.Default()
	if err != nil {
		panic(err)
	}

	entity.SetSnowflake(idGen)

	db, err := database.NewDatabase(&options.ODatabase{
		Dialector: sqlite.Open("test.db?_pragma=foreign_keys(1)"),
		Logger: options.ODatabaseLogger{
			OLogger: optionsLogger.OLogger{
				Tag: "TST",
			},
		},
	})
	if err != nil {
		panic(err)
	}

	option := &entity.Option{
		Key:       "test",
		Value:     "test",
		Encrypted: true,
	}
	tx := db.Create(option)
	if tx.Error != nil {
		panic(tx.Error)
	}

	log := &entity.Log{
		Tag:     "TEST",
		Code:    1,
		Message: "test test test ...",
	}
	tx = db.Create(log)
	if tx.Error != nil {
		panic(tx.Error)
	}

	var users []*entity.User
	for i := 0; i < 20; i++ {
		test := "test"
		if i%2 == 0 {
			test += "_"
		}
		users = append(users, &entity.User{
			Account:  fmt.Sprint(i) + test,
			Password: fmt.Sprint(i) + test,
		})
	}
	tx = db.CreateInBatches(users, len(users))
	if tx.Error != nil {
		panic(tx.Error)
	}

	println("Inserted", tx.RowsAffected, "users")

	crudUser := crud.NewCrud[*entity.User](db)
	queryUsers, err := crudUser.Query(options.NewOCriteria(&options.OCriteria{
		Page: 1,
		Size: 3,
		Sorts: []options.OCriteriaSort{
			{
				Column: "updated_at",
				Order:  options.OrdDescending,
			},
			{
				Column: "account",
				Order:  options.OrdDescending,
			},
		},
	}), func(where func(query any, args ...any)) {
		where("account LIKE ?", "%test_%")
	}, "OldPasswords")
	if err != nil {
		panic(err)
	}

	for _, u := range queryUsers {
		println(u.Account)
		println(u.OldPasswords[0].PasswordHash)
	}

	user1, err := crudUser.GetById(users[0].ID)
	if err != nil {
		panic(err)
	}
	println(user1.ID)

	users1, err := crudUser.All()
	if err != nil {
		panic(err)
	}
	println(len(users1))

	user2, err := crudUser.FindOne(func(where func(query any, args ...any)) {
		where("id = ?", user1.ID)
	})
	if err != nil {
		panic(err)
	}
	println(user2.Account)

	tmpHash := user2.PasswordHash
	user2.Account = "deleted"
	user2.Password = "abc123"
	affected, err := crudUser.Save(user2)
	if err != nil {
		panic(err)
	}
	println(user2.PasswordHash)
	println(tmpHash)
	println(user2.PasswordHash != tmpHash)
	println(affected)

	affected, err = crudUser.Delete(user2.ID)
	if err != nil {
		panic(err)
	}
	println(affected)

	crudOption := crud.NewCrudOption(db)
	option1, err := crudOption.GetById(option.ID)
	if err != nil {
		panic(err)
	}
	println(option.Value)
	println(option1.Value)
	println(option.Value == option1.Value)

	affected1, err := crudOption.Save(&entity.Option{
		Key:   "Not_Encrypted",
		Value: "Plain",
	})
	if err != nil {
		panic(err)
	}
	println(affected1)
}
