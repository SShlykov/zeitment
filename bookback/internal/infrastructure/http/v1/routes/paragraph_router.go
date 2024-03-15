package routes

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/controllers"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/pgrepo"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/SShlykov/zeitment/metrics"
	"github.com/SShlykov/zeitment/postgres"
	"github.com/labstack/echo/v4"
)

func Paragraph(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger loggerPkg.Logger, ctx context.Context) {
	repo := pgrepo.NewParagraphRepository(database)
	service := services.NewParagraphService(repo)
	cnt := controllers.NewParagraphController(service, metrics, logger, ctx)

	group := e.Group(v1.ParagraphsPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.POST(v1.ListSubPath, cnt.ListParagraphs)
	group.POST("", cnt.CreateParagraph)
	group.GET(v1.IDVar, cnt.GetParagraphByID)
	group.PUT(v1.IDVar, cnt.UpdateParagraph)
	group.DELETE(v1.IDVar, cnt.DeleteParagraph)
	group.POST("/pages"+v1.IDVar, cnt.GetParagraphsByPageID)
}
