package logger

import (
	"log/slog"
	"os"
)

type Logger interface {
	Warn(msg string, attrs ...any)
	Info(msg string, attrs ...any)
	Debug(msg string, attrs ...any)
	Error(msg string, attrs ...any)
}

type loggerImp struct {
	logger *slog.Logger
}

func SetupLogger(level string) Logger {
	levels := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}
	opts := PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: levels[level],
		},
	}

	logger := slog.New(NewPrettyHandler(os.Stdout, opts))

	return &loggerImp{logger: logger}
}

func (l *loggerImp) Warn(msg string, attrs ...any) {
	l.logger.Warn(msg, attrs...)
}

func (l *loggerImp) Info(msg string, attrs ...any) {
	l.logger.Info(msg, attrs...)
}

func (l *loggerImp) Debug(msg string, attrs ...any) {
	l.logger.Debug(msg, attrs...)
}

func (l *loggerImp) Error(msg string, attrs ...any) {
	l.logger.Error(msg, attrs...)
}
