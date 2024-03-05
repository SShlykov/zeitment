package swagger

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetSwagger(e *echo.Echo, swaggerEnabled bool) {
	if swaggerEnabled {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
}
