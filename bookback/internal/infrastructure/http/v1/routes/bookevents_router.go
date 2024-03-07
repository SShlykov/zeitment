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

func SetBookEventController(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewBookEventsRepository(database)
	service := services.NewBookEventsService(repo)
	cntr := controllers.NewBookEventController(service, metrics, logger, ctx)

	group := e.Group(BookEventsPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.POST("", cntr.CreateBookEvent)
	group.GET("/:id", cntr.GetBookEventByID)
	group.PUT("/:id", cntr.UpdateBookEvent)
	group.DELETE("/:id", cntr.DeleteBookEvent)
	group.GET("/book/:id", cntr.GetBookEventsByBookID)
	group.GET("/chapter/:id", cntr.GetBookEventsByChapterID)
	group.GET("/page/:id", cntr.GetBookEventsByPageID)
	group.GET("/paragraph/:id", cntr.GetBookEventsByParagraphID)
}
