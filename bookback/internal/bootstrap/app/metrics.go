package app

import (
	"github.com/SShlykov/zeitment/metrics/localmetrics"
)

func (app *App) initMetrics() error {
	logger := app.logger
	logger.Info("initializing metrics as local metrics")

	app.metrics = localmetrics.NewLocalMetrics(logger)
	return nil
}
