package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/circuitbreaker"
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

func (app *App) initEndpoint(_ context.Context) error {
	e := echo.New()
	cb := circuitbreaker.NewCircuitBreaker(app.config.RequestLimit, app.config.MinRequests, app.config.ErrorThresholdPercentage,
		app.config.IntervalDuration, app.config.OpenStateTimeout)

	e.Use(middleware.Recover())

	middlewares := []echo.MiddlewareFunc{
		loggerConfiguration(app.logger),
		middleware.Recover(),
		createCircuitBreakerMiddleware(cb),
	}

	e.Use(middlewares...)

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

func createCircuitBreakerMiddleware(cb *circuitbreaker.CircuitBreaker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := cb.Execute(func() error {
				return next(c)
			})

			if err != nil {
				if errors.Is(err, circuitbreaker.ErrorCb) {
					return c.JSON(http.StatusUnavailableForLegalReasons,
						map[string]string{"error": "Server is overloaded, please try again later.", "status": "error"})
				}
				return err
			}

			return nil
		}
	}
}

func loggerConfiguration(logger *slog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	})
}
