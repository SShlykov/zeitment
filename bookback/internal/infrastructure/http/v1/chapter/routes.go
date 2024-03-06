package chapter

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services/chapter"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func SetChapterController(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	chapterRepo := pgrepo.NewChapterRepository(database)
	service := chapter.NewService(chapterRepo)
	controller := NewController(service, metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

func (ch *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(pathPrefix)
	group.Use(middleware.MetricsLogger(ch.Metrics))

	group.GET("", ch.ListChapters)
	group.POST("", ch.CreateChapter)
	group.GET("/:id", ch.GetChapterByID)
	group.PUT("/:id", ch.UpdateChapter)
	group.DELETE("/:id", ch.DeleteChapter)
	group.GET("/book/:id", ch.GetChapterByBookID)
}
