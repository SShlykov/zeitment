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

// GetHealthCheck returns whether this application is alive or not.
// @Summary Get the status of this application
// @Description Get the status of this application
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {string} message "healthy: This application is started."
// @Failure 404 {string} message "None: This application is stopped."
// @Router /health [get]
func (hc *healthController) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}
