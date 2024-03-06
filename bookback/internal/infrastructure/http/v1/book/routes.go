package book

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services/book"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

// SetBookController регистрирует контроллер книг в маршрутизаторе.
func SetBookController(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewBookRepository(database)
	service := book.NewService(repo)
	controller := NewController(service, metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

// RegisterRoutes регистрирует маршруты для обработки запросов к книгам.
func (bc *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(PathPrefix)
	group.Use(middleware.MetricsLogger(bc.Metrics))

	group.POST("", bc.ListBooks)
	group.POST("", bc.CreateBook)
	group.GET("/:id", bc.GetBookByID)
	group.PUT("/:id", bc.UpdateBook)
	group.DELETE("/:id", bc.DeleteBook)
}
