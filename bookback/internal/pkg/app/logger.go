package app

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/pkg/logger"
)

func (a *App) initLogger(ctx context.Context) error {
	if a.config == nil {
		return errors.New("config is nil")
	}
	a.logger = logger.SetupLogger(a.config.Level)
	return nil
}
