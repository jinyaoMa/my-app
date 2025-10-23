package crud_test

import (
	"context"
	"errors"
	"runtime"
	"strings"
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
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
	"majinyao.cn/my-app/backend/pkg/db/model"
	"majinyao.cn/my-app/backend/pkg/snowflake"
	"majinyao.cn/my-app/backend/pkg/test"
)

type User struct {
	model.Model
	Account  datatype.Sensitive
	Password datatype.Password

	UserRoles []UserRole
	Roles     []Role `gorm:"many2many:user_roles;"`
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Account") {
		return errors.New("account not allowed to change")
	}
	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) Many2ManySetups() (setups []model.Many2ManySetup) {
	setups = append(setups, model.Many2ManySetup{
		Field:     "Roles",
		JoinTable: new(UserRole),
	})
	return
}

type UserScan struct {
	Id       int64
	Password string
}

type Role struct {
	model.Model
	model.ModelReserved
	model.ModelHmac
	Name string
	Desc string

	UserRoles []UserRole
	Users     []User `gorm:"many2many:user_roles;"`
}

func (r *Role) buildHmacData() string {
	return strings.Join([]string{r.Name, r.Desc}, ":")
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if err = r.ModelReserved.BeforeCreate(tx); err != nil {
		return err
	}
	if err = r.Model.BeforeCreate(tx); err != nil {
		return err
	}
	return
}

func (r *Role) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = r.ModelReserved.BeforeUpdate(tx); err != nil {
		return err
	}
	return nil
}

func (r *Role) AfterDelete(tx *gorm.DB) (err error) {
	return nil
}

func (r *Role) AfterUpdate(tx *gorm.DB) (err error) {
	if err = r.ModelHmac.Update(tx, r, r.buildHmacData); err != nil {
		return err
	}
	return nil
}

func (r *Role) AfterCreate(tx *gorm.DB) (err error) {
	if err = r.ModelHmac.Update(tx, r, r.buildHmacData); err != nil {
		return err
	}
	return nil
}

func (r *Role) AfterFind(tx *gorm.DB) (err error) {
	if err = r.ModelHmac.Verify(tx, r.buildHmacData); err != nil {
		return err
	}
	return nil
}

func (r *Role) Many2ManySetups() (setups []model.Many2ManySetup) {
	setups = append(setups, model.Many2ManySetup{
		Field:     "Users",
		JoinTable: new(UserRole),
	})
	return
}

type UserRole struct {
	model.Model

	UserId datatype.Id
	User   User

	RoleId datatype.Id
	Role   Role
}

func (r *UserRole) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (r *UserRole) BeforeDelete(tx *gorm.DB) (err error) {
	return nil
}

func (r *UserRole) AfterDelete(tx *gorm.DB) (err error) {
	return nil
}

func (r *UserRole) AfterUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (r *UserRole) AfterCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r *UserRole) AfterFind(tx *gorm.DB) (err error) {
	if err = r.Role.AfterFind(tx); err != nil {
		return
	}
	return
}

func (r *UserRole) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
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

	tx, err := db.Open([]model.Setup{
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

	tx1, cancel := dbcontext.SectionUnderContextWithCancel(ctx, tx)
	defer cancel()

	userService := crud.New[User](crud.DefaultOptions(), tx1)

	time.Sleep(time.Millisecond)
	_, _, err = userService.All(nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}

	// ctx, cancel = context.WithTimeout(context.Background(), time.Hour)
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx2, cancel := dbcontext.SectionUnderContextWithCancel(ctx, tx)
	cancel()

	userService = crud.New[User](crud.DefaultOptions(), tx2)

	_, _, err = userService.All(nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}

	tx3, cancel := dbcontext.SectionUnderContextWithTimeout(ctx, tx, 0)
	defer cancel()

	userService = crud.New[User](crud.DefaultOptions(), tx3)

	_, _, err = userService.All(nil)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}

	tx4, cancel := dbcontext.SectionUnderContextWithCancel(ctx, tx)
	defer cancel()

	userService = crud.New[User](crud.DefaultOptions(), tx4)

	newUser := User{
		Account: "1",
		Password: datatype.Password{
			Data: "1",
		},
	}
	affected, err := userService.Create(&newUser)
	if err != nil || affected != 1 {
		t.Fatal(affected, err)
	}

	newUsers := []User{
		{
			Account: "2",
			Password: datatype.Password{
				Data: "2",
			},
		},
		{
			Account: "3",
			Password: datatype.Password{
				Data: "3",
			},
		},
	}
	affected, err = userService.BatchCreate(&newUsers)
	if err != nil || affected != 2 {
		t.Fatal(affected, err)
	}

	newUsers[0].Account = "99"
	_, affected, err = userService.Update(newUsers[0])
	if !strings.Contains(err.Error(), "account not allowed to change") {
		t.Fatal(affected, err)
	}
	newUsers[0].Account = "2"

	_, affected, err = userService.Delete(newUser.Id)
	if err != nil || affected != 1 {
		t.Fatal(affected, err)
	}

	record, notFound, err := userService.GetById(newUsers[1].Id, nil)
	if err != nil || notFound || record.Account != newUsers[1].Account {
		t.Fatal(notFound, record.Account, err)
	}

	var scanUser UserScan
	notFound, err = userService.ScanById(&scanUser, record.Id, nil)
	if err != nil || notFound || scanUser.Id != record.Id.Int64() || scanUser.Password != record.Password.Data {
		t.Fatal(notFound, scanUser.Id, scanUser.Password, err)
	}

	var total int64

	records, total, err := userService.All(nil)
	if err != nil || len(records) != 2 || total != 2 || records[0].Account != newUsers[0].Account || records[1].Account != newUsers[1].Account {
		t.Fatal(len(records), total, records[0].Account, records[1].Account, err)
	}

	var scanUsers []UserScan
	total, err = userService.ScanAll(&scanUsers, nil)
	if err != nil || len(scanUsers) != 2 || total != 2 || scanUsers[0].Password != records[0].Password.Data || scanUsers[1].Password != records[1].Password.Data {
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
	if err != nil || len(scanUsers) != 1 || total != 2 || scanUsers[0].Password != records[0].Password.Data {
		t.Fatal(len(scanUsers), total, scanUsers[0].Password, err)
	}

	records, total, err = userService.Query(crud.Criteria{
		Page: 1,
		Size: 2,
		Filters: []crud.Filter{
			{
				Type:    crud.FilterTypeEqual,
				Field:   "users.password",
				Params:  []any{"2"},
				Special: crud.FilterSpecialPassword,
			},
		},
	}, func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(&User{Account: "2"}), nil
	})
	if err != nil || len(records) != 1 || total != 1 || records[0].Account != newUsers[0].Account {
		t.Fatal(len(records), total, records[0].Account, err)
	}

	total, err = userService.QueryScan(&scanUsers, crud.Criteria{
		Page: 1,
		Size: 2,
		Filters: []crud.Filter{
			{
				Type:    crud.FilterTypeEqual,
				Field:   "users.password",
				Params:  []any{"2"},
				Special: crud.FilterSpecialPassword,
			},
		},
	}, func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(&User{Account: "2"}), nil
	})
	if err != nil || len(scanUsers) != 1 || total != 1 || scanUsers[0].Password != records[0].Password.Data {
		t.Fatal(len(scanUsers), total, scanUsers[0].Password, err)
	}

	tx5, cancel := dbcontext.SectionUnderContextWithCancel(ctx, tx)
	defer cancel()

	roleService := crud.New[Role](crud.DefaultOptions(), tx5)

	newRole := Role{
		Name: "admin",
		Desc: "admin\\\"'%admin",
	}
	newRole.SysOp(true)
	affected, err = roleService.Create(&newRole)
	if err != nil || affected != 1 {
		t.Fatal(affected, err)
	}

	editRole := Role{
		Model: model.Model{
			Id: newRole.Id,
		},
	}
	newRole, affected, err = roleService.Update(editRole)
	if !strings.Contains(err.Error(), "can not be modified") {
		t.Fatal(affected, err)
	}
	_, affected, err = roleService.Delete(newRole.Id)
	if !strings.Contains(err.Error(), "can not be deleted") {
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

	tx6, cancel := dbcontext.SectionUnderContextWithCancel(ctx, tx)
	defer cancel()

	userRoleService := crud.New[UserRole](crud.DefaultOptions(), tx6)

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

	recordUserRole, notFound, err := userRoleService.GetById(newUserRoles[1].Id, []crud.Include{
		{
			Association: "User",
		},
		{
			Association: "Role",
		},
	})
	if err != nil || notFound || recordUserRole.User.Account != newUsers[1].Account || recordUserRole.Role.Name != newRole.Name || !recordUserRole.Role.HmacOk {
		t.Fatal(notFound, recordUserRole, err)
	}

	queryUserRoles, total, err := userRoleService.Query(crud.Criteria{
		Page: 1,
		Size: 10,
		Includes: []crud.Include{
			{
				Association: "User",
			},
			{
				Association: "Role",
			},
		},
		Joins: []string{"User", "Role"},
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
				Special: crud.FilterSpecialSensitive,
			},
			{
				Type:   crud.FilterTypeContains,
				Field:  "Role.desc",
				Params: []any{"n\\\"'%a"},
			},
		},
	})
	if err != nil || len(queryUserRoles) != 1 || total != 1 || queryUserRoles[0].User.Account != newUsers[0].Account || queryUserRoles[0].Role.Name != newRole.Name || !queryUserRoles[0].Role.HmacOk {
		t.Fatal(len(queryUserRoles), total, queryUserRoles[0], err)
	}
}
