package app

import (
	"context"
	"errors"
	"log/slog"
)

func (app *App) closer(ctx context.Context) error {
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), app.config.ShutdownTimeout)
	app.logger.Log(context.Background(), slog.LevelInfo, "Shutting down servers")
	defer cancel()
	if err := app.Echo.Shutdown(ctx); err != nil {
		return errors.New("failed to shutdown servers: " + err.Error())
	}

	return nil
}
