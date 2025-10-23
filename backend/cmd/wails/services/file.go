package services

import (
	"context"

	"github.com/wailsapp/wails/v3/pkg/application"
	"majinyao.cn/my-app/backend/internal/app"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type FileService struct {
	ctx         context.Context
	fileService service.IFileService
}

func (s *FileService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	s.ctx = ctx
	s.fileService, _ = service.NewFileService(s.ctx, app.CFG.Get().Crud, app.DB, app.STORE)
	return nil
}

func (s *FileService) ServiceShutdown() error {
	return nil
}

func (s *FileService) UploadAvatar(userId string) (avatar string) {
	uid, err := datatype.ParseIdFromB36(userId)
	if err != nil {
		app.LOG.Println(err)
		return
	}

	a := application.Get()
	t := app.I18N.GetTranslation()

	path, err := a.Dialog.OpenFileWithOptions(&application.OpenFileDialogOptions{
		Filters: []application.FileFilter{
			{
				DisplayName: t.Get("dialog.common.image.files"),
				Pattern:     "*.jpg;*.jpeg;*.png",
			},
		},
		Title:     t.Get("dialog.upload.avatar.title"),
		Directory: app.EXE.GetDir(),
	}).PromptForSingleSelection()
	if err != nil {
		app.LOG.Println(err)
		return
	}

	avatar, err = s.fileService.UploadAvatarFromLocalPath(uid, path)
	if err != nil {
		app.LOG.Println(err)
		return
	}
	return
}
