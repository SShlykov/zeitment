package app

import (
	"errors"
	"github.com/SShlykov/zeitment/bookback/pkg/logger"
)

func (app *App) initLogger() error {
	if app.config == nil {
		return errors.New("config is nil")
	}
	app.logger = logger.SetupLogger(app.config.Level)
	return nil
}
