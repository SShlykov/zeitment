package app

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/pkg/db/pg"
)

func (a *App) initDB(ctx context.Context) error {
	pgConf, err := config.NewPGConfig()
	if err != nil {
		return errors.New("failed to init pg config: " + err.Error())
	}
	db, err := pg.NewClient(ctx, pgConf.DSN())
	if err != nil {
		return errors.New("failed to init pg client: " + err.Error())
	}
	a.db = db
	return nil
}
