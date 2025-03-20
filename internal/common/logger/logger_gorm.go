package logger

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// MysqlLogger : Mysql custom Logger
type GormLogger struct {
}

func NewGormLogger() *GormLogger {
	return &GormLogger{}
}

// LogMode log mode
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

// Info print info
func (l GormLogger) Info(ctx context.Context, msg string, data ...any) {
	log.Info().Msgf(msg, data...)
}

// Warn print warn messages
func (l GormLogger) Warn(ctx context.Context, msg string, data ...any) {
	log.Warn().Msgf(msg, data...)
}

// Error print error messages
func (l GormLogger) Error(ctx context.Context, msg string, data ...any) {
	log.Error().Msgf(msg, data...)
}

// Trace print sql message
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	traceStr := "%s\n" + "[%.3fms] " + "[rows:%v]" + " %s"

	sql, rows := fc()
	if err != nil {
		if rows == -1 {
			log.Debug().Msgf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.Debug().Msgf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	} else {
		if rows == -1 {
			log.Trace().Msgf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.Trace().Msgf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
