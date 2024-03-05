package app

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"time"
)

func (app *App) initDB(ctx context.Context) error {
	pgConf, err := config.NewPGConfig()
	if err != nil {
		return errors.New("failed to init pg config: " + err.Error())
	}
	db, err := postgres.NewClient(ctx, app.logger, pgConf.DSN())
	if err != nil {
		return errors.New("failed to init pg client: " + err.Error())
	}
	app.db = db

	go func() {
		var broken int
		ticker := time.NewTicker(pgConf.PingInterval())
		for range ticker.C {
			err = db.DB().Ping(ctx)
			if err != nil {
				broken++
				if broken > pgConf.MaxPingAttempts() {
					app.logger.Error("postgres is down")
					app.closeCtx()
					break
				}
				app.logger.Error("failed to ping postgres")
			} else {
				broken = 0
			}
		}
	}()

	return nil
}
