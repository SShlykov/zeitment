package app

import (
	"errors"
	"github.com/SShlykov/zeitment/auth/pkg/config"
	"github.com/SShlykov/zeitment/logger"
)

func (app *App) initLogger(configPath string) error {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return errors.New("cant setup config")
	}
	app.logger = logger.SetupLogger(cfg.Level)
	return nil
}
