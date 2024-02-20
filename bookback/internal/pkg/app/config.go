package app

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/config"
)

func (app *App) initConfig(_ context.Context) error {
	cfg, err := config.LoadConfig(app.configPath)
	if err != nil {
		return errors.New("failed to load config: " + err.Error())
	}
	app.config = cfg

	return nil
}
