package middleware

import (
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/labstack/echo/v4"
	"time"
)

func MetricsLogger(metrics metrics.Metrics) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			duration := time.Since(start)

			req, res := c.Request(), c.Response()

			// Метрики
			status := fmt.Sprintf("%d", res.Status)
			metrics.IncCounter("http_requests_total", "method", req.Method, "path", req.URL.Path, "status", status)
			metrics.ObserveHistogram("http_request_duration_seconds", duration.Seconds(), "method", req.Method, "path", req.URL.Path)

			return err
		}
	}
}
