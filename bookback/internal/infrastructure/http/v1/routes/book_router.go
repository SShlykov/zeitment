package routes

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/controllers"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

// SetBookRoutes регистрирует контроллер книг в маршрутизаторе.
func SetBookRoutes(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewBookRepository(database)
	service := services.NewBookService(repo)
	cntr := controllers.NewBookController(service, metrics, logger, ctx)

	group := e.Group(BooksPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.POST("", cntr.ListBooks)
	group.POST("", cntr.CreateBook)
	group.GET("/:id", cntr.GetBookByID)
	group.PUT("/:id", cntr.UpdateBook)
	group.DELETE("/:id", cntr.DeleteBook)
}
