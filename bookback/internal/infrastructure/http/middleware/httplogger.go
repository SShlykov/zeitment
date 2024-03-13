package middleware

import (
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerConfiguration(logger loggerPkg.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.Info("REQUEST",
					loggerPkg.String("uri", v.URI),
					loggerPkg.Int("status", v.Status),
				)
			} else {
				logger.Error("REQUEST_ERROR",
					loggerPkg.String("uri", v.URI),
					loggerPkg.Int("status", v.Status),
					loggerPkg.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	})
}
