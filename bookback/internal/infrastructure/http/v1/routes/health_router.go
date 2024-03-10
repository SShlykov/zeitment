package routes

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/controllers"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func Health(e *echo.Echo, _ postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	cntr := controllers.NewHealthController(metrics, logger, ctx)

	group := e.Group(v1.HealthPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.GET("/", cntr.GetHealthCheck)
}
