package logger

import (
	"context"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"log/slog"
	"strconv"
	"time"
)

// SlogAdapter is an adapter to allow slog.Logger to be used as a GORM logger.
type SlogAdapter struct {
	slogLogger *slog.Logger
}

// NewSlogAdapter creates a new instance of SlogAdapter.
func NewSlogAdapter(logger *slog.Logger) *SlogAdapter {
	return &SlogAdapter{slogLogger: logger}
}

// LogMode sets log level.
func (s *SlogAdapter) LogMode(_ logger.LogLevel) logger.Interface {
	return s
}

// Info prints info messages.
func (s *SlogAdapter) Info(_ context.Context, msg string, _ ...interface{}) {
	s.slogLogger.Info(msg)
}

// Warn prints warning messages.
func (s *SlogAdapter) Warn(_ context.Context, msg string, _ ...interface{}) {
	s.slogLogger.Warn(msg)
}

// Error prints error messages.
func (s *SlogAdapter) Error(_ context.Context, msg string, _ ...interface{}) {
	s.slogLogger.Error(msg)
}

// Trace prints sql and elapsed time.
// elapsed := time.Since(begin).Milliseconds()
func (s *SlogAdapter) Trace(_ context.Context, _ time.Time, fc func() (string, int64), err error) {
	sql, execTime := fc()
	if err != nil {
		s.slogLogger.Error(utils.FileWithLineNum() + ": " + err.Error() + " " + sql + ", Time elapsed: " + strconv.FormatInt(execTime, 10) + "ms")
	} else {
		s.slogLogger.Debug(utils.FileWithLineNum() + ": " + sql + ", Time elapsed: " + strconv.FormatInt(execTime, 10) + "ms")
	}
}

// RowsAffected prints rows affected.
func (s *SlogAdapter) RowsAffected(_ context.Context, _ int64) {
	// do nothing
}

// RowsNotFound prints rows not found.
func (s *SlogAdapter) RowsNotFound(_ context.Context) {
	// do nothing
}
