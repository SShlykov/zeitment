package logger

import (
	"log/slog"
	"os"
)

func SetupLogger(level string) *slog.Logger {
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

	return logger
}
