package health

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/httpmiddlewares"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func SetHealthController(e *echo.Echo, _ db.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	controller := NewController(metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

func (hc *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(PathPrefix)
	group.Use(httpmiddlewares.MetricsLogger(hc.Metrics))

	group.GET("/", hc.GetHealthCheck)
}
