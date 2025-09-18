package app

import (
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db"
)

func SHUTDOWN_H3S(ctx context.Context) error {
	return H3S.Shutdown(ctx)
}

func RUN_H3S(ctx context.Context, isAuto bool) error {
	tx, cancel := db.SectionUnderContextWithCancel(ctx, DB)
	defer cancel()
	return tx.Transaction(func(tx *gorm.DB) error {
		var optionServerAutoRun,
			optionServerPort,
			optionServerSecurePort,
			optionServerCertFile,
			optionServerKeyFile entity.Option

		if isAuto {
			if res := tx.First(&optionServerAutoRun, entity.Option{
				Key: entity.OptionKeyServerAutoRun,
			}); res.Error != nil {
				return res.Error
			}
			if !optionServerAutoRun.GetBool() {
				return errors.New("server auto run is false")
			}
		}

		if res := tx.First(&optionServerPort, entity.Option{
			Key: entity.OptionKeyServerPort,
		}); res.Error != nil {
			return res.Error
		}
		if res := tx.First(&optionServerSecurePort, entity.Option{
			Key: entity.OptionKeyServerSecurePort,
		}); res.Error != nil {
			return res.Error
		}
		if res := tx.First(&optionServerCertFile, entity.Option{
			Key: entity.OptionKeyServerCertFile,
		}); res.Error != nil {
			return res.Error
		}
		if res := tx.First(&optionServerKeyFile, entity.Option{
			Key: entity.OptionKeyServerKeyFile,
		}); res.Error != nil {
			return res.Error
		}

		waitStart := make(chan any)
		go func() {
			err := H3S.Run(
				optionServerPort.GetUint16(),
				optionServerSecurePort.GetUint16(),
				optionServerCertFile.GetString(),
				optionServerKeyFile.GetString(),
				func() {
					close(waitStart)
				})
			if !errors.Is(err, http.ErrServerClosed) {
				LOG.Println("h3s server error:", err)
			}
		}()
		<-waitStart
		return nil
	})
}
