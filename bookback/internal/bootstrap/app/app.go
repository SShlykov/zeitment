package app

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/endpoint"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/config"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"

	"log/slog"
	"os"
	"os/signal"
	"sync"
)

type App struct {
	configPath string
	logger     *slog.Logger
	config     *config.Config
	db         postgres.Client
	web        *endpoint.Handler
	metrics    metrics.Metrics

	ctx      context.Context
	closeCtx func()
}

func NewApp(configPath string) (*App, error) {
	ctx, closeCtx := context.WithCancel(context.Background())
	app := &App{ctx: ctx, closeCtx: closeCtx, configPath: configPath}

	inits := []func() error{
		app.initConfig,
		app.initLogger,
		app.initMetrics,
		app.initDB,
		app.initWebServer,
	}

	for _, init := range inits {
		if err := init(); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (app *App) Run() error {
	ctx, stop := signal.NotifyContext(app.ctx, os.Interrupt)
	defer stop()

	logger := app.logger
	logger.Info("starting book app", slog.String("at", app.web.Address))
	logger.Debug("debug messages enabled")

	var wg sync.WaitGroup

	app.RunWebServer(&wg)

	return app.closer(ctx)
}
