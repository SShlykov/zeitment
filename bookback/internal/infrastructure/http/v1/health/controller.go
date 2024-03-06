package health

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Controller struct {
	Metrics metrics.Metrics
	Logger  *slog.Logger
	Ctx     context.Context
}

// NewController создает новый экземпляр Controller.
func NewController(metric metrics.Metrics, logger *slog.Logger, ctx context.Context) *Controller {
	return &Controller{Metrics: metric, Logger: logger, Ctx: ctx}
}

// GetHealthCheck возвращает статус приложения.
// @router /health [get]
// @summary Получить статус приложения
// @description Возвращает статус приложения
// @tags Статус приложения
// @produce  application/json
// @success 200 {string} string "healthy"
// @failure 500 {object} config.HTTPError
func (hc *Controller) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}
