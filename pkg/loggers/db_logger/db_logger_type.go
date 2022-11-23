package dbLogger

import (
	"context"
	"gorm.io/gorm/logger"
	"time"
)

type IDbLogger interface {
	Print(v ...interface{})
	LogMode(logLevel logger.LogLevel) logger.Interface
	Info(ctx context.Context, s string, args ...interface{})
	Warn(ctx context.Context, s string, args ...interface{})
	Error(ctx context.Context, s string, args ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}
type DbLogger struct {
	slowThreshold         time.Duration
	sourceField           string
	skipErrRecordNotFound bool
	debug                 bool
}

func NewDbLogger() *DbLogger {
	return &DbLogger{}
}
