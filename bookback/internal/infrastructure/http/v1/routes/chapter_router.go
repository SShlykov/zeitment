package routes

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/controllers"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func Chapter(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	chapterRepo := pgrepo.NewChapterRepository(database)
	service := services.NewChapterService(chapterRepo)
	cnt := controllers.NewChapterController(service, metrics, logger, ctx)

	group := e.Group(v1.ChaptersPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.POST("/list", cnt.ListChapters)
	group.POST("", cnt.CreateChapter)
	group.GET("/:id", cnt.GetChapterByID)
	group.PUT("/:id", cnt.UpdateChapter)
	group.DELETE("/:id", cnt.DeleteChapter)
	group.POST("/book/:id", cnt.GetChapterByBookID)
}
