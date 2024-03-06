package book

import (
	"context"
	bookrepo2 "github.com/SShlykov/zeitment/bookback/internal/adapters/db/postgres/bookrepo"
	"github.com/SShlykov/zeitment/bookback/internal/controller/http/httpmiddlewares"
	bookrepo "github.com/SShlykov/zeitment/bookback/internal/domain/services/book"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

// SetBookController регистрирует контроллер книг в маршрутизаторе.
func SetBookController(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	service := bookrepo.NewService(bookrepo2.NewRepository(database))
	controller := NewController(service, metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

// RegisterRoutes регистрирует маршруты для обработки запросов к книгам.
func (bc *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(PathPrefix)
	group.Use(httpmiddlewares.MetricsLogger(bc.Metrics))

	group.POST("", bc.ListBooks)
	group.POST("", bc.CreateBook)
	group.GET("/:id", bc.GetBookByID)
	group.PUT("/:id", bc.UpdateBook)
	group.DELETE("/:id", bc.DeleteBook)
}
