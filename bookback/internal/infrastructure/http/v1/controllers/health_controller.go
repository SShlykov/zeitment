package controllers

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type HealthController struct {
	Metrics metrics.Metrics
	Logger  loggerPkg.Logger
	Ctx     context.Context
}

// NewHealthController создает новый экземпляр Controller.
func NewHealthController(metric metrics.Metrics, logger loggerPkg.Logger, ctx context.Context) *HealthController {
	return &HealthController{Metrics: metric, Logger: logger, Ctx: ctx}
}

func (hc *HealthController) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "healthy", "status": "ok"})
}
