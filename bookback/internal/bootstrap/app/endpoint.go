package app

import (
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/endpoint"
	"github.com/SShlykov/zeitment/bookback/pkg/config"
	loggerPkg "github.com/SShlykov/zeitment/logger"
)

func (app *App) initWebServer() error {
	cfg, err := getConfig(app.configPath)
	if err != nil {
		return err
	}
	app.web, err = endpoint.NewHandler(app.db, app.minio, app.metrics, app.logger, app.ctx, cfg)

	if err != nil {
		return err
	}

	return nil
}

func (app *App) RunWebServer() {
	err := endpoint.RunServer(app.web, app.logger)
	if err != nil {
		app.logger.Error("HTTP server stopped", loggerPkg.Err(err))
	}
}

func getConfig(configPath string) (*endpoint.HTTPServerConfig, error) {
	cfg, err := config.LoadServerConfig(configPath)

	if err != nil {
		return nil, err
	}

	return FileConfigToServerConfig(cfg), nil
}

func FileConfigToServerConfig(cfg *config.HTTPServer) *endpoint.HTTPServerConfig {
	return &endpoint.HTTPServerConfig{
		RequestLimit:             cfg.RequestLimit,
		MinRequests:              cfg.MinRequests,
		ErrorThresholdPercentage: cfg.ErrorThresholdPercentage,
		IntervalDuration:         cfg.IntervalDuration,
		OpenStateTimeout:         cfg.OpenStateTimeout,
		CorsEnabled:              cfg.CorsEnabled,
		SwaggerEnabled:           cfg.SwaggerEnabled,
		Timeout:                  cfg.Timeout,
		IddleTimeout:             cfg.IddleTimeout,
		Address:                  cfg.Address,
	}
}
