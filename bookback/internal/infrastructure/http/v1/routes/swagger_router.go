package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetSwagger(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
