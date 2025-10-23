package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/view"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/authfwt"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
	"majinyao.cn/my-app/backend/pkg/storage"
)

type IFileService interface {
	crud.ICrud[entity.File]
	OpenFileById(id datatype.Id) (file *os.File, mime string, err error)
	UploadAvatarFromLocalPath(userId datatype.Id, path string) (avatar string, err error)
}

func NewFileService(ctx context.Context, options crud.Options, db *gorm.DB, store storage.IStorage) (IFileService, context.CancelFunc) {
	s, cancel := new(FileService).InitUnderContextWithCancel(ctx, options, db, store)
	return s, cancel
}

func UseFileService(ctx context.Context, options crud.Options, db *gorm.DB, store storage.IStorage) IFileService {
	return new(FileService).Init(ctx, options, db, store)
}

type FileService struct {
	crud.Crud[entity.File]
	ctx   context.Context
	store storage.IStorage
}

func (s *FileService) OpenFileById(id datatype.Id) (file *os.File, mime string, err error) {
	claims := authfwt.GetClaimsFromContext[view.UserData](s.ctx)
	if claims.Data.UserId == 0 {
		err = fmt.Errorf("user not found")
		return
	}

	targetFile, notFound, err := s.GetById(id, crud.Includes{
		{
			Association: "FileExtension",
			IgnoreSize:  true,
		},
		{
			Association: "FileUsers",
			IgnoreSize:  true,
		},
		{
			Association: "FileGroups.Group",
			IgnoreSize:  true,
			Includes: []crud.Include{
				{
					Association: "GroupUsers",
					IgnoreSize:  true,
				},
				{
					Association: "GroupRoles",
					IgnoreSize:  true,
				},
			},
		},
	})
	if err != nil || notFound {
		err = fmt.Errorf("file not found")
		return
	}
	if targetFile.IsDir.Valid && targetFile.IsDir.Bool {
		err = fmt.Errorf("file is directory")
		return
	}

	userId := claims.Data.GetUserId()
	roleId := claims.Data.GetRoleId()
	if slices.ContainsFunc(targetFile.FileUsers, func(fu entity.FileUser) bool {
		return (!fu.NoRead.Valid || !fu.NoRead.Bool) && fu.UserId == userId
	}) || slices.ContainsFunc(targetFile.FileGroups, func(fg entity.FileGroup) bool {
		return (!fg.NoRead.Valid || !fg.NoRead.Bool) && (slices.ContainsFunc(fg.Group.GroupUsers, func(fu entity.GroupUser) bool {
			return fu.UserId == userId
		}) || slices.ContainsFunc(fg.Group.GroupRoles, func(fr entity.GroupRole) bool {
			return fr.RoleId == roleId
		}))
	}) {
		mime = targetFile.FileExtension.Mime.String
		file, _, _, err = s.store.SearchDirectory(targetFile.FileExtension.Ext, targetFile.Size.Int64, targetFile.Checksums...)
		return
	}
	return nil, "", fmt.Errorf("file access denied")
}

func (s *FileService) UploadAvatarFromLocalPath(userId datatype.Id, path string) (avatar string, err error) {
	err = s.Db.Transaction(func(tx *gorm.DB) error {
		fileService := UseFileService(s.ctx, s.Options, tx, s.store)
		fileExtensionService := UseFileExtensionService(s.ctx, s.Options, tx)
		fileUserService := UseFileUserService(s.ctx, s.Options, tx)

		checksums, size, ext, err := s.store.PersistLocalFile(s.ctx, path)
		if err != nil {
			return err
		}

		base := filepath.Base(path)
		newFile := entity.File{
			Name:      strings.TrimSuffix(base, ext),
			Checksums: checksums,
			Size:      sql.NullInt64{Int64: size, Valid: true},
			Status:    entity.FileStatusPersisted,
			Readonly:  sql.NullBool{Bool: true, Valid: true},
			Hidden:    sql.NullBool{Bool: true, Valid: true},
		}

		targetFileExtension, notFound, _ := fileExtensionService.FindOne(func(tx *gorm.DB) (*gorm.DB, error) {
			return tx.Where(entity.FileExtension{Ext: ext}), nil
		}, nil)
		if !notFound {
			newFile.FileExtensionId = &targetFileExtension.Id
		}

		affected, err := fileService.Create(&newFile)
		if err != nil {
			return err
		}
		if affected == 0 {
			return fmt.Errorf("file service upload avatar [%d] from local path create file failed with affected 0", userId)
		}

		affected, err = fileUserService.Create(&entity.FileUser{
			FileId:   newFile.Id,
			UserId:   userId,
			IsAvatar: sql.NullBool{Bool: true, Valid: true},
		})
		if err != nil {
			return err
		}
		if affected == 0 {
			return fmt.Errorf("file service upload avatar [%d] from local path create file user failed with affected 0", userId)
		}

		var newOpenedFile *os.File
		newOpenedFile, _, _, err = s.store.SearchDirectory(targetFileExtension.Ext, newFile.Size.Int64, newFile.Checksums...)
		if err != nil {
			return err
		}

		var imgByte []byte
		imgByte, err = io.ReadAll(newOpenedFile)
		if err != nil {
			return err
		}

		avatar = "data:" + targetFileExtension.Mime.String + ";base64," + base64.StdEncoding.EncodeToString(imgByte)
		return nil
	})
	return
}

func (s *FileService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB, store storage.IStorage) (*FileService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db, store), cancel
}

func (s *FileService) Init(ctx context.Context, options crud.Options, db *gorm.DB, store storage.IStorage) *FileService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	s.store = store
	return s
}
