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
	Logger  *slog.Logger
	Ctx     context.Context
}

// NewHealthController создает новый экземпляр Controller.
func NewHealthController(metric metrics.Metrics, logger *slog.Logger, ctx context.Context) *HealthController {
	return &HealthController{Metrics: metric, Logger: logger, Ctx: ctx}
}

// GetHealthCheck возвращает статус приложения.
// @router /health [get]
// @summary Получить статус приложения
// @description Возвращает статус приложения
// @tags Статус приложения
// @produce  application/json
// @success 200 {string} string "healthy"
// @failure 500 {object} config.HTTPError
func (hc *HealthController) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}
