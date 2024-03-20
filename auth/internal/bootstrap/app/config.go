package app

import (
	"errors"
	"github.com/SShlykov/zeitment/auth/pkg/config"
)

func (app *App) initConfig(configPath string) error {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return errors.New("failed to load config: " + err.Error())
	}
	app.config = cfg

	return nil
}
