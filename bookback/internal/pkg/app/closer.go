package app

import (
	"context"
	"errors"
	"log/slog"
)

func (a *App) closer(ctx context.Context) error {
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), a.config.ShutdownTimeout)
	a.logger.Log(context.Background(), slog.LevelInfo, "Shutting down servers")
	defer cancel()
	if err := a.Echo.Shutdown(ctx); err != nil {
		return errors.New("failed to shutdown servers: " + err.Error())
	}

	return nil
}
