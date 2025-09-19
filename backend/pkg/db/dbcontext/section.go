package dbcontext

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// session with new timer context which will be deadline exceeded after the given timeout
func SectionWithTimeout(tx *gorm.DB, timeout time.Duration) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(tx.Statement.Context, timeout)
	return tx.Session(&gorm.Session{
		NewDB:   true,
		Context: ctx,
	}), cancel
}

// session with new cancel context which can be canceled after calling cancel function
func SectionWithCancelCause(tx *gorm.DB) (*gorm.DB, context.CancelCauseFunc) {
	ctx, cancel := context.WithCancelCause(tx.Statement.Context)
	return tx.Session(&gorm.Session{
		NewDB:   true,
		Context: ctx,
	}), cancel
}

// session with new timer context which will be deadline exceeded after the given timeout,
// but the timeout of the new timer context won't affect the given context
func SectionUnderContextWithTimeout(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*gorm.DB, context.CancelFunc) {
	tx, cancel := SectionWithTimeout(tx, timeout)
	stop := context.AfterFunc(ctx, cancel)
	return tx, func() {
		stop()
		cancel()
	}
}

// session with new cancel context which can be canceled by the given context,
// but the cancellation of the new cancel context won't affect the given context
func SectionUnderContextWithCancel(ctx context.Context, tx *gorm.DB) (*gorm.DB, context.CancelFunc) {
	tx, cancel := SectionWithCancelCause(tx)
	stop := context.AfterFunc(ctx, func() {
		cancel(context.Cause(ctx))
	})
	return tx, func() {
		stop()
		cancel(context.Canceled)
	}
}
