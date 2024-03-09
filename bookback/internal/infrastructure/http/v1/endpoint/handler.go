package endpoint

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/routes"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/circuitbreaker"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
	"time"
)

type Handler struct {
	Timeout      time.Duration
	IddleTimeout time.Duration
	Address      string

	e *echo.Echo
}

type HTTPServerConfig struct {
	RequestLimit             int
	MinRequests              int
	ErrorThresholdPercentage float64
	IntervalDuration         time.Duration
	OpenStateTimeout         time.Duration
	CorsEnabled              bool
	SwaggerEnabled           bool
	Timeout                  time.Duration
	IddleTimeout             time.Duration
	Address                  string
}

func (h *Handler) Shutdown(ctx context.Context) error {
	return h.e.Shutdown(ctx)
}

func NewHandler(database postgres.Client, metric metrics.Metrics,
	logger *slog.Logger, ctx context.Context, cfg *HTTPServerConfig) (*Handler, error) {
	e := echo.New()

	setMiddlewares(e, logger, cfg)

	setRouter(e, database, metric, logger, ctx)

	if cfg.SwaggerEnabled {
		routes.SetSwagger(e)
	}

	return &Handler{e: e, Timeout: cfg.Timeout, IddleTimeout: cfg.IddleTimeout, Address: cfg.Address}, nil
}

func setMiddlewares(e *echo.Echo, logger *slog.Logger, config *HTTPServerConfig) {
	cb := circuitbreaker.NewCircuitBreaker(
		config.RequestLimit,
		config.MinRequests,
		config.ErrorThresholdPercentage,
		config.IntervalDuration,
		config.OpenStateTimeout,
	)

	middlewares := []echo.MiddlewareFunc{
		middleware.LoggerConfiguration(logger),
		middleware.Recover(),
		middleware.CreateCircuitBreakerMiddleware(cb),
	}

	if config.CorsEnabled {
		middlewares = append(middlewares, middleware.CORS())
	}

	e.Use(middlewares...)
}

func setRouter(e *echo.Echo, database postgres.Client, metric metrics.Metrics,
	logger *slog.Logger, ctx context.Context) {
	controllers :=
		[]func(e *echo.Echo, database postgres.Client, metric metrics.Metrics,
			logger *slog.Logger, ctx context.Context){
			routes.Health,
			routes.Book,
			routes.Chapter,
			routes.Page,
			routes.Paragraph,
			routes.BookEvent,
			routes.MapVariables,
		}

	for _, controller := range controllers {
		controller(e, database, metric, logger, ctx)
	}

	return
}
