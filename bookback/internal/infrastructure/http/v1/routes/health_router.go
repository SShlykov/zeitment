package routes

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/controllers"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func SetHealthController(e *echo.Echo, _ postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	cntr := controllers.NewHealthController(metrics, logger, ctx)

	group := e.Group(HealthPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.GET("/", cntr.GetHealthCheck)
}
