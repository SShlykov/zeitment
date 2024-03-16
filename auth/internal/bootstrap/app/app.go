package app

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/grpc/server"
	"github.com/SShlykov/zeitment/auth/pkg/config"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/SShlykov/zeitment/metrics"
	"github.com/SShlykov/zeitment/postgres"
	"os"
	"os/signal"
	"sync"
)

type App struct {
	logger  loggerPkg.Logger
	config  *config.Config
	db      postgres.Client
	metrics metrics.Metrics

	ctx      context.Context
	closeCtx func()
}

func NewApp(configPath string) (*App, error) {
	ctx, closeCtx := context.WithCancel(context.Background())
	app := &App{ctx: ctx, closeCtx: closeCtx}

	inits := []func(cfg string) error{
		app.initLogger,
		app.initMetrics,
		app.initDB,
	}

	for _, init := range inits {
		if err := init(configPath); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (app *App) Run() error {
	ctx, stop := signal.NotifyContext(app.ctx, os.Interrupt)
	defer stop()

	var wg sync.WaitGroup

	logg := app.logger
	logg.Info("starting auth app")
	logg.Debug("debug messages enabled")

	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = server.NewServer(logg, app.db)
	}()

	<-ctx.Done()
	return nil
}
