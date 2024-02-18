package app

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
)

func (a *App) initEndpoint(ctx context.Context) error {
	e := echo.New()
	e.Use(middleware.Recover())
	addLogger(e, a.logger, ctx)
	a.Echo = e

	return nil
}

func (a *App) initRouter(_ context.Context) error {
	router.SetCORSConfig(a.Echo, a.config.CorsEnabled)
	router.SetHealthController(a.Echo, a.ctx)
	router.SetBookController(a.Echo, a.db, a.ctx)
	router.SetSwagger(a.Echo, a.config.SwaggerEnabled)

	return nil
}

func addLogger(e *echo.Echo, logger *slog.Logger, ctx context.Context) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(ctx, slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(ctx, slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
}
