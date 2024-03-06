package app

import (
	"context"
	"errors"
	"log/slog"
)

func (app *App) closer(ctx context.Context) error {
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), app.config.ShutdownTimeout)
	app.logger.Log(context.Background(), slog.LevelInfo, "Shutting down controller")
	defer cancel()
	if err := app.Echo.Shutdown(shutdownCtx); err != nil {
		return errors.New("failed to shutdown controller: " + err.Error())
	}

	return nil
}
