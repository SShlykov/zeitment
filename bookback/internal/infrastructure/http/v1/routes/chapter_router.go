package routes

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/controllers"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/SShlykov/zeitment/postgres"
	"github.com/labstack/echo/v4"
)

func Chapter(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger loggerPkg.Logger, ctx context.Context) {
	chapterRepo := pgrepo.NewChapterRepository(database)
	service := services.NewChapterService(chapterRepo)
	cnt := controllers.NewChapterController(service, metrics, logger, ctx)

	group := e.Group(v1.ChaptersPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.POST(v1.ListSubPath, cnt.ListChapters)
	group.POST("", cnt.CreateChapter)
	group.GET(v1.IDVar, cnt.GetChapterByID)
	group.PUT(v1.IDVar, cnt.UpdateChapter)
	group.DELETE(v1.IDVar, cnt.DeleteChapter)
	group.POST("/book"+v1.IDVar, cnt.GetChapterByBookID)
}
