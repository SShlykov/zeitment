package middleware

import (
	"errors"
	"github.com/SShlykov/zeitment/bookback/pkg/circuitbreaker"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateCircuitBreakerMiddleware(cb *circuitbreaker.CircuitBreaker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := cb.Execute(func() error {
				return next(c)
			})

			if err != nil {
				if errors.Is(err, circuitbreaker.ErrorCb) {
					return c.JSON(http.StatusUnavailableForLegalReasons,
						map[string]string{"error": "Server is overloaded, please try again later.", "status": "error"})
				}
				return err
			}

			return nil
		}
	}
}
