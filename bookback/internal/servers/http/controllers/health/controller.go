package health

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller interface {
	GetHealthCheck(c echo.Context) error
	RegisterRoutes(e *echo.Echo, _ context.Context)
}

type healthController struct {
}

// NewController is constructor.
func NewController() Controller {
	return &healthController{}
}

func (hc *healthController) RegisterRoutes(e *echo.Echo, _ context.Context) {
	e.GET("/api/v1/health", hc.GetHealthCheck)
}

// GetHealthCheck возвращает статус приложения.
// @router /health [get]
// @summary Получить статус приложения
// @description Возвращает статус приложения
// @tags Статус приложения
// @produce  application/json
// @success 200 {string} string "healthy"
// @failure 500 {object} config.HTTPError
func (hc *healthController) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}
