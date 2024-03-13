package app

import (
	"context"
	"errors"
)

func (app *App) closer(ctx context.Context) error {
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), app.config.ShutdownTimeout)
	app.logger.Info("Shutting down controller")
	defer cancel()
	if err := app.web.Shutdown(shutdownCtx); err != nil {
		return errors.New("failed to shutdown controller: " + err.Error())
	}

	return nil
}
