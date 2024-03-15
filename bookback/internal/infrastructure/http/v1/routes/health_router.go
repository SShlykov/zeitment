package routes

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/controllers"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/SShlykov/zeitment/metrics"
	"github.com/SShlykov/zeitment/postgres"
	"github.com/labstack/echo/v4"
)

func Health(e *echo.Echo, _ postgres.Client, metrics metrics.Metrics, logger loggerPkg.Logger, ctx context.Context) {
	cntr := controllers.NewHealthController(metrics, logger, ctx)

	group := e.Group(v1.HealthPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.GET("", cntr.GetHealthCheck)
}
