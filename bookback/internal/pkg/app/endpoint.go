package app

import (
	"context"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"sync"
)

func (app *App) runWebServer(wg *sync.WaitGroup, ctx context.Context) {

	wg.Add(1)

	go func() {
		defer wg.Done()
		httpServer := &http.Server{
			ReadHeaderTimeout: app.config.Timeout,
			ReadTimeout:       app.config.Timeout,
			WriteTimeout:      app.config.Timeout,
			IdleTimeout:       app.config.IddleTimeout,
			Addr:              fmt.Sprintf(app.config.Address),
			Handler:           app.Echo,
		}

		app.logger.Info("Started servers", slog.String("address", app.config.Address))
		err := httpServer.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

}

func (app *App) initEndpoint(ctx context.Context) error {
	e := echo.New()
	e.Use(middleware.Recover())
	addLogger(e, app.logger, ctx)
	app.Echo = e

	return nil
}

func (app *App) initRouter(_ context.Context) error {
	router.SetCORSConfig(app.Echo, app.config.CorsEnabled)
	router.SetHealthController(app.Echo, app.ctx)
	router.SetBookController(app.Echo, app.db, app.ctx)
	router.SetSwagger(app.Echo, app.config.SwaggerEnabled)

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
