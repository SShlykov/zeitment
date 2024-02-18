package app

import (
	"context"
	"fmt"
	cfg "github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
)

type App struct {
	configPath string
	logger     *slog.Logger
	config     *cfg.Config
	db         db.Client
	Echo       *echo.Echo

	ctx      context.Context
	closeCtx func()
}

func NewApp(configPath string) (*App, error) {
	ctx, closeCtx := context.WithCancel(context.Background())
	app := &App{ctx: ctx, closeCtx: closeCtx, configPath: configPath}

	inits := []func(ctx context.Context) error{
		app.initConfig,
		app.initLogger,
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

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(a.ctx, os.Interrupt)
	defer stop()

	go func() {
		httpServer := &http.Server{
			ReadHeaderTimeout: a.config.Timeout,
			ReadTimeout:       a.config.Timeout,
			WriteTimeout:      a.config.Timeout,
			IdleTimeout:       a.config.IddleTimeout,
			Addr:              fmt.Sprintf(a.config.Address),
			Handler:           a.Echo,
		}

		a.logger.Info("Started servers", slog.String("address", a.config.Address))
		err := httpServer.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	return a.closer(ctx)
}
