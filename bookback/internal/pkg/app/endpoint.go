package app

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/circuitbreaker"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/book"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/bookevents"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/chapter"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/health"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/mapvariables"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/page"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/paragraph"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/swagger"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/httpmiddlewares"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"sync"
)

func (app *App) runWebServer(wg *sync.WaitGroup, _ context.Context) {
	go func() {
		wg.Add(1)
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
		httpmiddlewares.LoggerConfiguration(app.logger),
		middleware.Recover(),
		httpmiddlewares.CreateCircuitBreakerMiddleware(cb),
	}

	if app.config.CorsEnabled {
		middlewares = append(middlewares, httpmiddlewares.CORS())
	}

	e.Use(middlewares...)

	app.Echo = e
	return nil
}

func (app *App) initRouter(_ context.Context) error {
	controllers := []func(e *echo.Echo, database db.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context){
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
