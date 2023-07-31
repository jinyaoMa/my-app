package main

import (
	"fmt"
	ientity "my-app/backend/internal/entity"
	"my-app/backend/internal/service"
	"my-app/backend/pkg/crypto"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/vmodel"
	"my-app/backend/pkg/helper"
	"my-app/backend/pkg/logger"
	"my-app/backend/pkg/snowflake"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func main() {
	key, err := helper.GetFilenameSameAsExecutable("option.key")
	if err != nil {
		panic(err)
	}

	entity.Cipher(crypto.NewAesWithSalt(key))

	idGen, err := snowflake.Default()
	if err != nil {
		panic(err)
	}

	entity.IdGenerator(idGen)

	db, err := database.New(&database.Option{
		Dialector: sqlite.Open("test.db?_pragma=foreign_keys(1)"),
		OnInitialized: func(db *gorm.DB) {
			logs := new(entity.Log)
			options := new(entity.Option)
			db.Use(dbresolver.Register(dbresolver.Config{
				Sources: []gorm.Dialector{sqlite.Open("test.logs.db")},
			}, logs).Register(dbresolver.Config{
				Sources: []gorm.Dialector{sqlite.Open("test.options.db")},
			}, options))
			db.Clauses(dbresolver.Use("logs")).AutoMigrate(logs)
			db.Clauses(dbresolver.Use("options")).AutoMigrate(options)
		},
		Migrate: []interface{}{
			new(ientity.Node),
		},
		Logger: database.OptionLogger{
			Option: logger.Option{
				Tag: "DBS",
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
		Message: "[TEST] test test test ...",
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

	crudUser := database.NewCrudService[*entity.User](db)
	queryUsers, err := crudUser.Query(vmodel.NewCriteria(&vmodel.Criteria{
		Page: 1,
		Size: 3,
		Sorts: []vmodel.CriteriaSort{
			{
				Column: "updated_at",
				Order:  vmodel.OrdDescending,
			},
			{
				Column: "account",
				Order:  vmodel.OrdDescending,
			},
		},
	}), func(where func(query any, args ...any)) {
		where("account LIKE ?", "%test_%")
	}, "UserPasswords")
	if err != nil {
		panic(err)
	}

	for _, u := range queryUsers {
		println(u.Account)
		println(len(u.UserPasswords))
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

	crudOption := service.NewOptionService(db)
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

	ext := &entity.FileExtension{
		Name:    "Ext",
		DotName: "ext",
	}
	tx = db.Create(ext)
	if tx.Error != nil {
		panic(tx.Error)
	}

	file := &entity.File{
		Path:            "/",
		Name:            "file",
		Size:            0,
		FileExtensionID: ext.ID,
		Checksum:        "00000000000000000000000000000000:00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000:0",
		UserID:          user1.ID,
		AccessableUsers: []*entity.User{
			user2,
		},
	}
	tx = db.Create(file)
	if tx.Error != nil {
		panic(tx.Error)
	}
}
