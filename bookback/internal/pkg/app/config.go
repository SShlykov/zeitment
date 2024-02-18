package app

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/config"
)

func (a *App) initConfig(_ context.Context) error {
	cfg, err := config.LoadConfig(a.configPath)
	if err != nil {
		return errors.New("failed to load config: " + err.Error())
	}
	a.config = cfg

	return nil
}
