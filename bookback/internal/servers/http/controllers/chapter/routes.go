package chapter

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/httpmiddlewares"
	bookrepo "github.com/SShlykov/zeitment/bookback/internal/services/book"
	chapterrepo "github.com/SShlykov/zeitment/bookback/internal/services/chapter"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func SetChapterController(e *echo.Echo, database db.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	chapterRepo := chapterrepo.NewRepository(database)
	bookRepo := bookrepo.NewRepository(database)
	service := chapterrepo.NewService(chapterRepo, bookRepo)
	controller := NewController(service, metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

func (ch *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(pathPrefix)
	group.Use(httpmiddlewares.MetricsLogger(ch.Metrics))

	group.GET("", ch.ListChapters)
	group.POST("", ch.CreateChapter)
	group.GET("/:id", ch.GetChapterByID)
	group.PUT("/:id", ch.UpdateChapter)
	group.DELETE("/:id", ch.DeleteChapter)
	group.GET("/book/:id", ch.GetChapterByBookID)
}
