package app

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/book"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/bookevents"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/chapter"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/health"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/mapvariables"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/page"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/paragraph"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/swagger"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/circuitbreaker"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	echomv "github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"sync"
)

func (app *App) runWebServer(wg *sync.WaitGroup, _ context.Context) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpServer := &http.Server{
			ReadHeaderTimeout: app.config.Timeout,
			ReadTimeout:       app.config.Timeout,
			WriteTimeout:      app.config.Timeout,
			IdleTimeout:       app.config.IddleTimeout,
			Addr:              app.config.Address,
			Handler:           app.Echo,
		}

		app.logger.Info("HTTP server started")
		err := httpServer.ListenAndServe()
		if err != nil {
			app.logger.Error("HTTP server stopped", slog.Group("err", err))
		}
	}()
}

func (app *App) initEndpoint(_ context.Context) error {
	e := echo.New()
	cb := circuitbreaker.NewCircuitBreaker(
		app.config.RequestLimit,
		app.config.MinRequests,
		app.config.ErrorThresholdPercentage,
		app.config.IntervalDuration,
		app.config.OpenStateTimeout,
	)

	middlewares := []echo.MiddlewareFunc{
		middleware.LoggerConfiguration(app.logger),
		echomv.Recover(),
		middleware.CreateCircuitBreakerMiddleware(cb),
	}

	if app.config.CorsEnabled {
		middlewares = append(middlewares, middleware.CORS())
	}

	e.Use(middlewares...)

	app.Echo = e
	return nil
}

func (app *App) initRouter(_ context.Context) error {
	controllers := []func(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context){
		health.SetHealthController,
		book.SetBookController,
		chapter.SetChapterController,
		page.SetPageController,
		paragraph.SetParagraphController,
		bookevents.SetBookEventController,
		mapvariables.SetMapVariablesController,
	}

	for _, controller := range controllers {
		controller(app.Echo, app.db, app.metrics, app.logger, app.ctx)
	}

	swagger.SetSwagger(app.Echo, app.config.SwaggerEnabled)

	return nil
}
