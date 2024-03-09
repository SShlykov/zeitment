package app

import (
	"errors"
	"github.com/SShlykov/zeitment/bookback/pkg/config"
)

func (app *App) initConfig() error {
	cfg, err := config.LoadConfig(app.configPath)
	if err != nil {
		return errors.New("failed to load config: " + err.Error())
	}
	app.config = cfg

	return nil
}
