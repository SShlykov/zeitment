package app

import "github.com/SShlykov/zeitment/metrics/localmetrics"

func (app *App) initMetrics(_ string) error {
	logger := app.logger
	logger.Info("initializing metrics as local metrics")

	app.metrics = localmetrics.NewLocalMetrics(logger)
	return nil
}
