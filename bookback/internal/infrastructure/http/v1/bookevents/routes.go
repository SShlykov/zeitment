package bookevents

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services/bookevents"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func SetBookEventController(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewBookEventsRepository(database)
	service := bookevents.NewService(repo)
	controller := NewController(service, metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

func (bec *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(PathPrefix)
	group.Use(middleware.MetricsLogger(bec.Metrics))

	group.POST("", bec.CreateBookEvent)
	group.GET("/:id", bec.GetBookEventByID)
	group.PUT("/:id", bec.UpdateBookEvent)
	group.DELETE("/:id", bec.DeleteBookEvent)
	group.GET("/book/:id", bec.GetBookEventsByBookID)
	group.GET("/chapter/:id", bec.GetBookEventsByChapterID)
	group.GET("/page/:id", bec.GetBookEventsByPageID)
	group.GET("/paragraph/:id", bec.GetBookEventsByParagraphID)
}
