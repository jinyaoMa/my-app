package crud_test

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"majinyao.cn/my-app/backend/pkg/cflog"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/sm4"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sm3"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen/argon2"
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
	"majinyao.cn/my-app/backend/pkg/snowflake"
	"majinyao.cn/my-app/backend/pkg/test"
)

type User struct {
	model.Model
	Account  datatype.Encrypted
	Password datatype.Password

	UserRoles []UserRole
	Roles     []Role `gorm:"many2many:user_roles;"`
}

type UserScan struct {
	Id       int64
	Password string
}

type Role struct {
	model.Model
	Name   string
	Desc   string
	Hmac   datatype.Hashed
	HmacOk bool `gorm:"-:all"`

	UserRoles []UserRole
	Users     []User `gorm:"many2many:user_roles;"`
}

func (r *Role) BeforeSave(tx *gorm.DB) (err error) {
	r.Hmac = datatype.Hashed(fmt.Sprint(r.Name, r.Desc))
	return nil
}

func (r *Role) AfterSave(tx *gorm.DB) (err error) {
	r.HmacOk, err = r.Hmac.VerifyBase64(tx, fmt.Sprint(r.Name, r.Desc))
	return
}

func (r *Role) AfterFind(tx *gorm.DB) (err error) {
	r.HmacOk, err = r.Hmac.VerifyBase64(tx, fmt.Sprint(r.Name, r.Desc))
	return
}

type UserRole struct {
	model.Model

	UserId datatype.Id
	User   User

	RoleId datatype.Id
	Role   Role
}

func (r *UserRole) AfterFind(tx *gorm.DB) (err error) {
	if err = r.Role.AfterFind(tx); err != nil {
		return
	}
	return
}

func (r *UserRole) GetM2MSetups() []model.M2MSetup {
	return []model.M2MSetup{
		{
			Model:     &User{},
			Field:     "Roles",
			JoinTable: &UserRole{},
		},
		{
			Model:     &Role{},
			Field:     "Users",
			JoinTable: &UserRole{},
		},
	}
}

func TestCrud(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	epoch := time.Now()
	text := "mjy"
	threads := runtime.NumCPU()
	keyGen, _ := keygen.New(keygen.Options{
		Alg:       argon2.Alg,
		Salt:      text,
		Threads:   threads,
		KeyLength: 16,
	})
	ivGen, _ := keygen.New(keygen.Options{
		Alg:       argon2.Alg,
		Salt:      text,
		Threads:   threads,
		KeyLength: 12,
	})
	key := keyGen.DeriveBase64(text, false)
	iv := ivGen.DeriveBase64(text, false)

	tx, err := db.Open([]any{
		&User{},
		&Role{},
		&UserRole{},
	}, db.Options{
		Cflog: cflog.Options{
			EnableConsole: true,
		},
		LogLevel: logger.Info,
		Driver:   db.DrvSqlite,
		Dsn:      ":memory:?_pragma=foreign_keys(1)",
		Snowflake: snowflake.Options{
			Epoch:    epoch,
			NodeBits: 7,
			StepBits: 14,
			NodeId:   1,
		},
		Keygen: keygen.Options{
			Alg:       argon2.Alg,
			Salt:      text,
			Threads:   threads,
			KeyLength: 32,
		},
		Hasher: hasher.Options{
			Alg:  sm3.Alg,
			Salt: text,
			Key:  text,
		},
		Cipher: cipher.Options{
			Alg: sm4.Alg,
			AAD: text,
			Key: key,
			Iv:  iv,
		},
		AutoMigrate: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	userService, cancel := crud.NewtWithCancelUnderContext[User](ctx, tx)
	defer cancel()

	_, _, err = userService.All()
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	userService, cancel = crud.NewtWithCancelUnderContext[User](ctx, tx)
	cancel()

	_, _, err = userService.All()
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}

	userService, cancel = crud.NewWithTimeoutUnderContext[User](ctx, tx, 0)
	defer cancel()

	_, _, err = userService.All()
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}

	userService, cancel = crud.NewtWithCancelUnderContext[User](ctx, tx)
	defer cancel()

	newUser := User{
		Account:  datatype.Encrypted("1"),
		Password: datatype.Password("1"),
	}
	affected, err := userService.Create(&newUser)
	if err != nil || affected != 1 {
		t.Fatal(affected, err)
	}

	newUsers := []User{
		{
			Account:  datatype.Encrypted("2"),
			Password: datatype.Password("2"),
		},
		{
			Account:  datatype.Encrypted("3"),
			Password: datatype.Password("3"),
		},
	}
	affected, err = userService.BatchCreate(&newUsers)
	if err != nil || affected != 2 {
		t.Fatal(affected, err)
	}

	affected, err = userService.Delete(newUser.Id)
	if err != nil || affected != 1 {
		t.Fatal(affected, err)
	}

	record, notFound, err := userService.GetById(newUsers[1].Id)
	if err != nil || notFound || record.Account != newUsers[1].Account {
		t.Fatal(notFound, record.Account, err)
	}

	var scanUser UserScan
	notFound, err = userService.ScanById(&scanUser, record.Id)
	if err != nil || notFound || scanUser.Id != record.Id.Int64() || scanUser.Password != string(record.Password) {
		t.Fatal(notFound, scanUser.Id, scanUser.Password, err)
	}

	var total int64

	records, total, err := userService.All()
	if err != nil || len(records) != 2 || total != 2 || records[0].Account != newUsers[0].Account || records[1].Account != newUsers[1].Account {
		t.Fatal(len(records), total, records[0].Account, records[1].Account, err)
	}

	var scanUsers []UserScan
	total, err = userService.ScanAll(&scanUsers)
	if err != nil || len(scanUsers) != 2 || total != 2 || scanUsers[0].Password != string(records[0].Password) || scanUsers[1].Password != string(records[1].Password) {
		t.Fatal(len(scanUsers), total, scanUsers[0].Password, scanUsers[1].Password, err)
	}

	records, total, err = userService.Query(crud.Criteria{
		Page: 2,
		Size: 1,
	})
	if err != nil || len(records) != 1 || total != 2 || records[0].Account != newUsers[1].Account {
		t.Fatal(len(records), total, records[0].Account, err)
	}

	total, err = userService.QueryScan(&scanUsers, crud.Criteria{
		Page: 2,
		Size: 1,
	})
	if err != nil || len(scanUsers) != 1 || total != 2 || scanUsers[0].Password != string(records[0].Password) {
		t.Fatal(len(scanUsers), total, scanUsers[0].Password, err)
	}

	records, total, err = userService.QueryWithCondition(crud.Criteria{
		Page: 1,
		Size: 2,
		Filters: []crud.Filter{
			{
				Type:    crud.FilterTypeEqual,
				Field:   "users.password",
				Params:  []any{string(newUsers[0].Password)},
				Special: crud.FilterSpecialPassword,
			},
		},
	}, func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(&User{Account: datatype.Encrypted("2")}), nil
	})
	if err != nil || len(records) != 1 || total != 1 || records[0].Account != newUsers[0].Account {
		t.Fatal(len(records), total, records[0].Account, err)
	}

	total, err = userService.QueryScanWithCondition(&scanUsers, crud.Criteria{
		Page: 1,
		Size: 2,
		Filters: []crud.Filter{
			{
				Type:    crud.FilterTypeEqual,
				Field:   "users.password",
				Params:  []any{string(newUsers[0].Password)},
				Special: crud.FilterSpecialPassword,
			},
		},
	}, func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(&User{Account: datatype.Encrypted("2")}), nil
	})
	if err != nil || len(scanUsers) != 1 || total != 1 || scanUsers[0].Password != string(records[0].Password) {
		t.Fatal(len(scanUsers), total, scanUsers[0].Password, err)
	}

	roleService, cancel := crud.NewtWithCancelUnderContext[Role](ctx, tx)
	defer cancel()

	newRole := Role{
		Name: "admin",
		Desc: "admin",
	}
	affected, err = roleService.Create(&newRole)
	if err != nil || affected != 1 {
		t.Fatal(affected, err)
	}

	queryRoles, total, err := roleService.Query(crud.Criteria{
		Page: 1,
		Size: 10,
		Filters: []crud.Filter{
			{
				Type:   crud.FilterTypeEqual,
				Field:  "roles.name",
				Params: []any{"admin"},
			},
		},
	})
	if err != nil || len(queryRoles) != 1 || total != 1 || !queryRoles[0].HmacOk {
		t.Fatal(len(queryRoles), total, queryRoles[0], err)
	}

	userRoleService, cancel := crud.NewtWithCancelUnderContext[UserRole](ctx, tx)
	defer cancel()

	newUserRoles := []UserRole{
		{
			UserId: newUsers[0].Id,
			RoleId: newRole.Id,
		},
		{
			UserId: newUsers[1].Id,
			RoleId: newRole.Id,
		},
	}
	affected, err = userRoleService.BatchCreate(&newUserRoles)
	if err != nil || affected != 2 {
		t.Fatal(affected, err)
	}

	recordUserRole, notFound, err := userRoleService.GetById(newUserRoles[1].Id, "User", "Role")
	if err != nil || notFound || recordUserRole.User.Account != newUsers[1].Account || recordUserRole.Role.Name != newRole.Name || !recordUserRole.Role.HmacOk {
		t.Fatal(notFound, recordUserRole, err)
	}

	queryUserRoles, total, err := userRoleService.Query(crud.Criteria{
		Page:     1,
		Size:     10,
		Includes: []string{"User", "Role"},
		Joins:    []string{"User", "Role"},
		Filters: []crud.Filter{
			{
				Type:    crud.FilterTypeEqual,
				Field:   "user_roles.user_id",
				Params:  []any{newUsers[0].Id.B36String()},
				Special: crud.FilterSpecialIdString,
			},
			{
				Type:    crud.FilterTypeEqual,
				Field:   "User.account",
				Params:  []any{string(newUsers[0].Account)},
				Special: crud.FilterSpecialEncrypted,
			},
			{
				Type:   crud.FilterTypeEqual,
				Field:  "Role.desc",
				Params: []any{newRole.Desc},
			},
		},
	})
	if err != nil || len(queryUserRoles) != 1 || total != 2 || queryUserRoles[0].User.Account != newUsers[0].Account || queryUserRoles[0].Role.Name != newRole.Name || !queryUserRoles[0].Role.HmacOk {
		t.Fatal(len(queryUserRoles), total, queryUserRoles[0], err)
	}
}
