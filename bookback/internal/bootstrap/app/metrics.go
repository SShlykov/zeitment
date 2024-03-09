package app

import (
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics/localmetrics"
)

func (app *App) initMetrics() error {
	logger := app.logger
	logger.Info("initializing metrics as local metrics")

	app.metrics = localmetrics.NewLocalMetrics(logger)
	return nil
}
