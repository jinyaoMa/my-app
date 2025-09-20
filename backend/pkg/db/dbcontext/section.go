package dbcontext

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// session with new timer context which will be deadline exceeded after the given timeout
func SectionWithTimeout(db *gorm.DB, timeout time.Duration) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(db.Statement.Context, timeout)
	return db.Session(&gorm.Session{
		NewDB:   true,
		Context: ctx,
	}), cancel
}

// session with new cancel context which can be canceled after calling cancel function
func SectionWithCancelCause(db *gorm.DB) (*gorm.DB, context.CancelCauseFunc) {
	ctx, cancel := context.WithCancelCause(db.Statement.Context)
	return db.Session(&gorm.Session{
		NewDB:   true,
		Context: ctx,
	}), cancel
}

// session with new timer context which will be deadline exceeded after the given timeout,
// but the timeout of the new timer context won't affect the given context
func SectionUnderContextWithTimeout(ctx context.Context, db *gorm.DB, timeout time.Duration) (*gorm.DB, context.CancelFunc) {
	db, cancel := SectionWithTimeout(db, timeout)
	stop := context.AfterFunc(ctx, cancel)
	return db, func() {
		stop()
		cancel()
	}
}

// session with new cancel context which can be canceled by the given context,
// but the cancellation of the new cancel context won't affect the given context
func SectionUnderContextWithCancel(ctx context.Context, db *gorm.DB) (*gorm.DB, context.CancelFunc) {
	db, cancel := SectionWithCancelCause(db)
	stop := context.AfterFunc(ctx, func() {
		cancel(context.Cause(ctx))
	})
	return db, func() {
		stop()
		cancel(context.Canceled)
	}
}
