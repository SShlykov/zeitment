package app

import (
	"context"
	cfg "github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/labstack/echo/v4"
	"log/slog"
	"os"
	"os/signal"
	"sync"
)

type App struct {
	configPath string
	logger     *slog.Logger
	config     *cfg.Config
	db         db.Client
	Echo       *echo.Echo
	metrics    metrics.Metrics

	ctx      context.Context
	closeCtx func()
}

func NewApp(configPath string) (*App, error) {
	ctx, closeCtx := context.WithCancel(context.Background())
	app := &App{ctx: ctx, closeCtx: closeCtx, configPath: configPath}

	inits := []func(ctx context.Context) error{
		app.initConfig,
		app.initLogger,
		app.initMetrics,
		app.initDB,
		app.initEndpoint,
		app.initRouter,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (app *App) Run() error {
	ctx, stop := signal.NotifyContext(app.ctx, os.Interrupt)
	defer stop()

	logger := app.logger
	logger.Info("starting book app", slog.String("at", app.config.Address))
	logger.Debug("debug messages enabled")

	var wg sync.WaitGroup

	runs := []func(*sync.WaitGroup, context.Context){
		app.runWebServer,
	}

	for _, run := range runs {
		run(&wg, ctx)
	}

	return app.closer(ctx)
}
