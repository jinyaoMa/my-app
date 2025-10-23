package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/view"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/authbase"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/authfwt"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/memcache"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
	"majinyao.cn/my-app/backend/pkg/flag"
	"majinyao.cn/my-app/backend/pkg/storage"
)

type IUserService interface {
	crud.ICrud[entity.User]
	authbase.Verifier[view.UserData]
	VerifyLoginWithTemporaryPassword(input *authbase.LoginInput) (userdata view.UserData, err error)
	VerifyUserDataRoleId(userdata view.UserData, visitorId string, newRoleId datatype.Id) (newUserdata view.UserData, err error)
	GetCurrentUser() (user view.CurrentUser, err error)
	GetByAccountTemporaryPassword(account string, password string, includes crud.Includes) (user entity.User, notFound bool, err error)
	GetByAccountPassword(account string, password string, includes crud.Includes) (user entity.User, notFound bool, err error)
	CreateReservedUser(form view.CreateReservedUserForm) (info view.ReservedUserInfo, err error)
	GetReservedUserInfo() (info view.ReservedUserInfo, err error)
}

func NewUserService(ctx context.Context, options crud.Options, db *gorm.DB, store storage.IStorage) (IUserService, context.CancelFunc) {
	s, cancel := new(UserService).InitUnderContextWithCancel(ctx, options, db, store)
	return s, cancel
}

func UseUserService(ctx context.Context, options crud.Options, db *gorm.DB, store storage.IStorage) IUserService {
	return new(UserService).Init(ctx, options, db, store)
}

type UserService struct {
	crud.Crud[entity.User]
	ctx   context.Context
	store storage.IStorage
}

func (s *UserService) VerifyLoginWithTemporaryPassword(input *authbase.LoginInput) (userdata view.UserData, err error) {
	now := time.Now()
	user, notFound, err := s.GetByAccountTemporaryPassword(input.Body.Account, input.Body.Password, []crud.Include{
		{
			Association: "UserRoles.Role.Permissions",
			IgnoreSize:  false,
		},
	})
	if err != nil {
		return
	}
	if notFound {
		err = errors.New("account or temporary password wrong")
		return
	}
	if !user.TemporaryPasswordExpiredAt.Valid || user.TemporaryPasswordExpiredAt.Time.Before(now) {
		err = errors.New("temporary password expired")
		return
	}

	result := s.Db.Model(&user).Updates(entity.User{
		TemporaryPasswordExpiredAt: sql.NullTime{Time: now, Valid: true},
	})
	if result.Error != nil {
		return
	}
	if result.RowsAffected == 0 {
		err = errors.New("expire temporary password failed")
		return
	}

	var roleId int64
	var flag flag.IFlag
	for _, userRole := range user.UserRoles {
		if !userRole.IsDefault.Valid || !userRole.IsDefault.Bool {
			continue
		}
		for _, perm := range userRole.Role.Permissions {
			flag = perm.GetFlag().Or(flag)
		}
		roleId = userRole.RoleId.Int64()
		break
	}

	cache := memcache.GetFromContext(s.ctx)
	_, _, err = cache.Set(fmt.Sprintf("auth_perm_%d_%d", user.Id.Int64(), roleId), flag)
	if err != nil {
		return
	}

	return view.UserData{
		VisitorId: input.VisitorId,
		UserId:    user.Id.Int64(),
		RoleId:    roleId,
	}, nil
}

func (s *UserService) VerifyUserDataRoleId(userdata view.UserData, visitorId string, newRoleId datatype.Id) (newUserdata view.UserData, err error) {
	newUserdata, err = s.VerifyUserData(userdata, visitorId)
	if err != nil {
		return
	}

	userRoleService := UseUserRoleService(s.ctx, s.Options, s.Db)
	_, notFound, err := userRoleService.FindOne(func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(entity.UserRole{
			UserId: datatype.Id(newUserdata.UserId),
			RoleId: newRoleId,
		}), nil
	}, nil)
	if err != nil || notFound {
		err = fmt.Errorf("new role not found")
		return
	}

	newUserdata.RoleId = newRoleId.Int64()
	return
}

func (s *UserService) VerifyUserData(userdata view.UserData, visitorId string) (newUserdata view.UserData, err error) {
	if userdata.VisitorId != visitorId {
		return userdata, fmt.Errorf("visitor id not matched")
	}
	userRoleService := UseUserRoleService(s.ctx, s.Options, s.Db)
	_, notFound, err := userRoleService.FindOne(func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(entity.UserRole{
			UserId: datatype.Id(userdata.UserId),
			RoleId: datatype.Id(userdata.RoleId),
		}), nil
	}, nil)
	if err != nil || notFound {
		err = fmt.Errorf("role not found")
		return
	}
	return userdata, nil
}

func (s *UserService) VerifyLogin(input *authbase.LoginInput) (userdata view.UserData, err error) {
	user, notFound, err := s.GetByAccountPassword(input.Body.Account, input.Body.Password, []crud.Include{
		{
			Association: "UserRoles.Role.Permissions",
			IgnoreSize:  false,
		},
	})
	if err != nil {
		return
	}
	if notFound {
		err = errors.New("account or password wrong")
		return
	}

	var roleId int64
	var flag flag.IFlag
	for _, userRole := range user.UserRoles {
		if !userRole.IsDefault.Valid || !userRole.IsDefault.Bool {
			continue
		}
		for _, perm := range userRole.Role.Permissions {
			flag = perm.GetFlag().Or(flag)
		}
		roleId = userRole.RoleId.Int64()
		break
	}

	cache := memcache.GetFromContext(s.ctx)
	_, _, err = cache.Set(fmt.Sprintf("auth_perm_%d_%d", user.Id.Int64(), roleId), flag)
	if err != nil {
		return
	}

	return view.UserData{
		VisitorId: input.VisitorId,
		UserId:    user.Id.Int64(),
		RoleId:    roleId,
	}, nil
}

func (s *UserService) GetCurrentUser() (user view.CurrentUser, err error) {
	claims := authfwt.GetClaimsFromContext[view.UserData](s.ctx)
	if claims.Data.UserId == 0 {
		err = fmt.Errorf("user not found")
		return
	}

	var userEntity entity.User
	userEntity, notFound, err := s.GetCopyById(&user, claims.Data.GetUserId(), []crud.Include{
		{
			Association: "Roles",
			IgnoreSize:  false,
		},
		{
			Association: "FileUsers",
			IgnoreSize:  false,
			Condition: func(tx *gorm.DB) *gorm.DB {
				return tx.Where("is_avatar = ?", true)
			},
			Includes: []crud.Include{
				{
					Association: "File.FileExtension",
					IgnoreSize:  false,
				},
			},
		},
	})
	if err != nil || notFound {
		err = fmt.Errorf("user not found")
		return
	}

	roleIdB36 := claims.Data.GetRoleId().B36String()
	for i := range user.Roles {
		if user.Roles[i].Id == roleIdB36 {
			user.Roles[i].IsCurrent = true
			break
		}
	}

	if idx := slices.IndexFunc(userEntity.FileUsers, func(fu entity.FileUser) bool {
		return fu.IsAvatar.Valid && fu.IsAvatar.Bool
	}); idx != -1 {
		targetFile := userEntity.FileUsers[idx].File
		if targetFile.FileExtension == nil {
			return
		}

		var targetOpenedFile *os.File
		targetOpenedFile, _, _, err = s.store.SearchDirectory(targetFile.FileExtension.Ext, targetFile.Size.Int64, targetFile.Checksums...)
		if err != nil {
			return
		}
		defer targetOpenedFile.Close()

		var imgByte []byte
		imgByte, err = io.ReadAll(targetOpenedFile)
		if err != nil {
			return
		}
		user.Avatar = "data:" + targetFile.FileExtension.Mime.String + ";base64," + base64.StdEncoding.EncodeToString(imgByte)
	}
	return
}

func (s *UserService) GetByAccountTemporaryPassword(account string, password string, includes crud.Includes) (user entity.User, notFound bool, err error) {
	var tPwd = datatype.Sensitive(password)
	return s.FindOne(func(tx *gorm.DB) (*gorm.DB, error) {
		tx = tx.Where(entity.User{
			Account:           datatype.Sensitive(account),
			TemporaryPassword: &tPwd,
		})
		return tx, nil
	}, includes)
}

func (s *UserService) GetByAccountPassword(account string, password string, includes crud.Includes) (user entity.User, notFound bool, err error) {
	return s.FindOne(func(tx *gorm.DB) (*gorm.DB, error) {
		tx = tx.Where(entity.User{
			Account: datatype.Sensitive(account),
			Password: datatype.Password{
				Data: password,
			},
		})
		return tx, nil
	}, includes)
}

func (s *UserService) CreateReservedUser(form view.CreateReservedUserForm) (info view.ReservedUserInfo, err error) {
	err = s.Db.Transaction(func(tx *gorm.DB) error {
		roleService := UseRoleService(s.ctx, s.Options, tx)
		userService := UseUserService(s.ctx, s.Options, tx, s.store)
		userRoleService := UseUserRoleService(s.ctx, s.Options, tx)

		reservedRole, notFound, err := roleService.FindOne(func(tx *gorm.DB) (*gorm.DB, error) {
			return tx.Where("reserved = ?", true), nil
		}, nil)
		if err != nil {
			return err
		}
		if notFound {
			return errors.New("reserved role not found")
		}

		var user entity.User
		err = copier.CopyWithOption(&user, &form, crud.GlobalCopierOption)
		if err != nil {
			return err
		}

		user.SysOp(true)
		affected, err := userService.Create(&user)
		if err != nil {
			return err
		}
		if affected < 1 {
			return errors.New("create reserved user failed")
		}

		affected, err = userRoleService.Create(&entity.UserRole{
			UserId:    user.Id,
			RoleId:    reservedRole.Id,
			IsDefault: sql.NullBool{Bool: true, Valid: true},
		})
		if err != nil {
			return err
		}
		if affected < 1 {
			return errors.New("create reserved user role failed")
		}

		err = copier.CopyWithOption(&info, &user, crud.GlobalCopierOption)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

func (s *UserService) GetReservedUserInfo() (info view.ReservedUserInfo, err error) {
	var user entity.User
	err = s.Db.Preload("FileUsers", "is_avatar = ?", true).
		Preload("FileUsers.File.FileExtension").
		Where("reserved = ?", true).
		Take(&user).Error
	if err != nil {
		return
	}

	err = copier.CopyWithOption(&info, &user, crud.GlobalCopierOption)
	if err != nil {
		return
	}

	if idx := slices.IndexFunc(user.FileUsers, func(fu entity.FileUser) bool {
		return fu.IsAvatar.Valid && fu.IsAvatar.Bool
	}); idx != -1 {
		targetFile := user.FileUsers[idx].File
		if targetFile.FileExtension == nil {
			return
		}

		var targetOpenedFile *os.File
		targetOpenedFile, _, _, err = s.store.SearchDirectory(targetFile.FileExtension.Ext, targetFile.Size.Int64, targetFile.Checksums...)
		if err != nil {
			return
		}
		defer targetOpenedFile.Close()

		var imgByte []byte
		imgByte, err = io.ReadAll(targetOpenedFile)
		if err != nil {
			return
		}
		info.Avatar = "data:" + targetFile.FileExtension.Mime.String + ";base64," + base64.StdEncoding.EncodeToString(imgByte)
	}
	return
}

func (s *UserService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB, store storage.IStorage) (*UserService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db, store), cancel
}

func (s *UserService) Init(ctx context.Context, options crud.Options, db *gorm.DB, store storage.IStorage) *UserService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	s.store = store
	return s
}
