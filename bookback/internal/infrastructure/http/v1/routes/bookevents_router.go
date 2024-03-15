package routes

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/controllers"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/pgrepo"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/SShlykov/zeitment/postgres"
	"github.com/labstack/echo/v4"
)

func BookEvent(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger loggerPkg.Logger, ctx context.Context) {
	repo := pgrepo.NewBookEventsRepository(database)
	service := services.NewBookEventsService(repo)
	cntr := controllers.NewBookEventController(service, metrics, logger, ctx)

	group := e.Group(v1.BookEventsPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.POST("", cntr.CreateBookEvent)
	group.GET(v1.IDVar, cntr.GetBookEventByID)
	group.PUT(v1.IDVar, cntr.UpdateBookEvent)
	group.DELETE(v1.IDVar, cntr.DeleteBookEvent)
	group.POST(v1.BookSubPath+v1.IDVar, cntr.GetBookEventsByBookID)
	group.POST(v1.ChapterSubPath+v1.IDVar, cntr.GetBookEventsByChapterID)
	group.POST("/page"+v1.IDVar, cntr.GetBookEventsByPageID)
	group.POST("/paragraph"+v1.IDVar, cntr.GetBookEventsByParagraphID)
}
