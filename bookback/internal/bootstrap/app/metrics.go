package app

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/metrics/localmetrics"
)

func (app *App) initMetrics(_ context.Context) error {
	logger := app.logger
	logger.Info("initializing metrics as local metrics")

	app.metrics = localmetrics.NewLocalMetrics(logger)
	return nil
}
