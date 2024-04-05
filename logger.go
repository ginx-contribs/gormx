package gormx

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"log/slog"
	"time"
)

func NewLogger(logger *slog.Logger) *Logger {
	return &Logger{logger: logger}
}

// Logger customizes gorm sql logger with slog.Logger
type Logger struct {
	logger *slog.Logger
	level  logger.LogLevel
}

func (g *Logger) LogMode(level logger.LogLevel) logger.Interface {
	g.level = level
	return g
}

func (g *Logger) Info(ctx context.Context, s string, i ...interface{}) {
	if g.level >= logger.Info {
		g.logger.InfoContext(ctx, fmt.Sprintf(s, i...))
	}
}

func (g *Logger) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.level >= logger.Warn {
		g.logger.WarnContext(ctx, fmt.Sprintf(s, i...))
	}
}

func (g *Logger) Error(ctx context.Context, s string, i ...interface{}) {
	if g.level >= logger.Error {
		g.logger.ErrorContext(ctx, fmt.Sprintf(s, i...))
	}
}

func (g *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.level <= logger.Silent {
		return
	}
	cost := time.Now().Sub(begin)
	sql, affected := fc()
	slog.DebugContext(ctx, "[Gorm Trace]", slog.String("cost", fmt.Sprintf("%dms", cost/time.Millisecond)), slog.Int64("row affected", affected), slog.String("sql", sql), slog.Any("error", err))
}
