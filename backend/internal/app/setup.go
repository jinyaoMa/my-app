package app

import (
	"context"
	"mime"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/cflog"
	"majinyao.cn/my-app/backend/pkg/config"
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/executable"
	"majinyao.cn/my-app/backend/pkg/flag"
	"majinyao.cn/my-app/backend/pkg/h3server"
	"majinyao.cn/my-app/backend/pkg/i18n"
	"majinyao.cn/my-app/backend/pkg/reactive"
	"majinyao.cn/my-app/backend/pkg/router"
	"majinyao.cn/my-app/backend/pkg/storage"
)

var (
	HasReservedUsersOnStartup bool
)

func setup(
	ctx context.Context,
	exe executable.IExecutable,
	cfg config.IConfig[Config],
	log *cflog.Cflog,
	store storage.IStorage,
	i19 i18n.II18n,
	tx *gorm.DB,
	api router.IRouter,
	h3s h3server.IH3Server,
) {
	// setup data
	tx, cancel := db.SectionUnderContextWithCancel(ctx, tx)
	defer cancel()
	err := tx.Transaction(func(tx *gorm.DB) error {
		optionSystemLocale := entity.Option{
			Key: entity.OptionKeySystemLocale,
		}
		optionSystemLocale.SysOp(true)
		optionSystemLocale.SetString("zh-CN")
		if res := tx.FirstOrCreate(&optionSystemLocale, entity.Option{
			Key: entity.OptionKeySystemLocale,
		}); res.Error != nil {
			return res.Error
		}

		i19.SetLocale(optionSystemLocale.GetString(), true)
		i19.Watch(func(t i18n.Translation) (err error) {
			return DB.Transaction(func(tx *gorm.DB) error {
				var optionSystemLocale entity.Option
				if res := tx.First(&optionSystemLocale, entity.Option{
					Key: entity.OptionKeySystemLocale,
				}); res.Error != nil {
					return res.Error
				}
				optionSystemLocale.SysOp(true)
				optionSystemLocale.SetString(t.Code)
				return tx.Save(&optionSystemLocale).Error
			})
		})
		translation := i19.GetTranslation()

		optionSystemColorTheme := entity.Option{
			Key: entity.OptionKeySystemColorTheme,
		}
		optionSystemColorTheme.SysOp(true)
		optionSystemColorTheme.SetColorTheme(entity.OptionColorThemeSystem)
		if res := tx.FirstOrCreate(&optionSystemColorTheme, entity.Option{
			Key: entity.OptionKeySystemColorTheme,
		}); res.Error != nil {
			return res.Error
		}

		THEME, _ = reactive.New(optionSystemColorTheme.GetColorTheme())
		THEME.Watch(func(value string) (err error) {
			return DB.Transaction(func(tx *gorm.DB) error {
				var optionSystemColorTheme entity.Option
				if res := tx.First(&optionSystemColorTheme, entity.Option{
					Key: entity.OptionKeySystemColorTheme,
				}); res.Error != nil {
					return res.Error
				}
				optionSystemColorTheme.SysOp(true)
				optionSystemColorTheme.SetColorTheme(value)
				return tx.Save(&optionSystemColorTheme).Error
			})
		})

		lastOperationIdEnumPair := entity.OperationIdEnumPair{}
		if res := tx.First(&lastOperationIdEnumPair); res.Error != nil {
			return res.Error
		}

		codePerrmissionFullControl := "full.control"
		permissionFullControl := entity.Permission{
			Code:        codePerrmissionFullControl,
			Name:        translation.Get("permission." + codePerrmissionFullControl),
			Description: translation.Get("permission." + codePerrmissionFullControl),
		}
		permissionFullControl.SysOp(true)
		permissionFullControl.SetFlag(flag.Make(lastOperationIdEnumPair.Enum+1, true))
		if res := tx.FirstOrCreate(&permissionFullControl, entity.Permission{
			Code: codePerrmissionFullControl,
		}); res.Error != nil {
			return res.Error
		}

		codeRoleSuperUser := "super.user"
		roleSuperUser := entity.Role{
			Code:        codeRoleSuperUser,
			Name:        translation.Get("role." + codeRoleSuperUser),
			Description: translation.Get("role." + codeRoleSuperUser),
		}
		roleSuperUser.SysOp(true)
		if res := tx.FirstOrCreate(&roleSuperUser, entity.Role{
			Code: codeRoleSuperUser,
		}); res.Error != nil {
			return res.Error
		}

		rolePermissionSuperUser := entity.RolePermission{
			RoleId:       roleSuperUser.Id,
			PermissionId: permissionFullControl.Id,
		}
		rolePermissionSuperUser.SysOp(true)
		if res := tx.FirstOrCreate(&rolePermissionSuperUser, entity.RolePermission{
			RoleId:       roleSuperUser.Id,
			PermissionId: permissionFullControl.Id,
		}); res.Error != nil {
			return res.Error
		}

		codeGroupEveryone := "everyone"
		groupEveryone := entity.Group{
			Code:        codeGroupEveryone,
			Name:        translation.Get("group." + codeGroupEveryone),
			Description: translation.Get("group." + codeGroupEveryone),
		}
		groupEveryone.SysOp(true)
		if res := tx.FirstOrCreate(&groupEveryone, entity.Group{
			Code: codeGroupEveryone,
		}); res.Error != nil {
			return res.Error
		}

		groupRoleEveryone := entity.GroupRole{
			GroupId: groupEveryone.Id,
			RoleId:  roleSuperUser.Id,
		}
		groupRoleEveryone.SysOp(true)
		if res := tx.FirstOrCreate(&groupRoleEveryone, entity.GroupRole{
			GroupId: groupEveryone.Id,
			RoleId:  roleSuperUser.Id,
		}); res.Error != nil {
			return res.Error
		}

		optionServerAutoRun := entity.Option{
			Key: entity.OptionKeyServerAutoRun,
		}
		optionServerAutoRun.SysOp(true)
		optionServerAutoRun.SetBool(true)
		if res := tx.FirstOrCreate(&optionServerAutoRun, entity.Option{
			Key: entity.OptionKeyServerAutoRun,
		}); res.Error != nil {
			return res.Error
		}

		optionServerPort := entity.Option{
			Key: entity.OptionKeyServerPort,
		}
		optionServerPort.SysOp(true)
		optionServerPort.SetUint16(18080)
		if res := tx.FirstOrCreate(&optionServerPort, entity.Option{
			Key: entity.OptionKeyServerPort,
		}); res.Error != nil {
			return res.Error
		}

		optionServerSecurePort := entity.Option{
			Key: entity.OptionKeyServerSecurePort,
		}
		optionServerSecurePort.SysOp(true)
		optionServerSecurePort.SetUint16(18443)
		if res := tx.FirstOrCreate(&optionServerSecurePort, entity.Option{
			Key: entity.OptionKeyServerSecurePort,
		}); res.Error != nil {
			return res.Error
		}

		optionServerCertFile := entity.Option{
			Key: entity.OptionKeyServerCertFile,
		}
		optionServerCertFile.SysOp(true)
		optionServerCertFile.SetString(exe.GetPathWithExt(".cert"))
		if res := tx.FirstOrCreate(&optionServerCertFile, entity.Option{
			Key: entity.OptionKeyServerCertFile,
		}); res.Error != nil {
			return res.Error
		}

		optionServerKeyFile := entity.Option{
			Key: entity.OptionKeyServerKeyFile,
		}
		optionServerKeyFile.SysOp(true)
		optionServerKeyFile.SetString(exe.GetPathWithExt(".key"))
		if res := tx.FirstOrCreate(&optionServerKeyFile, entity.Option{
			Key: entity.OptionKeyServerKeyFile,
		}); res.Error != nil {
			return res.Error
		}

		for codeFileCategory, fileExtensions := range map[string][]string{
			"document": {".doc", ".docx", ".pdf", ".txt", ".rtf", ".xls", ".xlsx", ".csv", ".ppt", ".pptx", ".md", ".html", ".xml", ".log"},
			"image":    {".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".svg", ".webp", ".ico", ".heic"},
			"video":    {".mp4", ".avi", ".mov", ".mkv", ".wmv", ".flv", ".webm", ".3gp"},
			"audio":    {".mp3", ".wav", ".aac", ".flac", ".ogg", ".m4a", ".wma", ".midi"},
		} {
			fileCategory := entity.FileCategory{
				Code: codeFileCategory,
				Name: translation.Get("file.category." + codeFileCategory),
			}
			fileCategory.SysOp(true)
			if res := tx.FirstOrCreate(&fileCategory, entity.FileCategory{
				Code: codeFileCategory,
			}); res.Error != nil {
				return res.Error
			}

			for _, ext := range fileExtensions {
				fileExtension := entity.FileExtension{
					Ext:            ext,
					Name:           translation.Get("file.extension" + ext),
					Mime:           mime.TypeByExtension(ext),
					FileCategoryId: &fileCategory.Id,
				}
				fileExtension.SysOp(true)
				if res := tx.FirstOrCreate(&fileExtension, entity.FileExtension{
					Ext: ext,
				}); res.Error != nil {
					return res.Error
				}
			}
		}

		var count int64
		tx.Table("users").Where("reserved = ?", true).Count(&count)
		HasReservedUsersOnStartup = count > 0
		return nil
	})
	if err != nil {
		log.Panicln("setup data failed", err)
	}

	// setup server
	RUN_H3S(ctx, true)
}
