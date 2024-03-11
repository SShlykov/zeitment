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

func MapVariables(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewMapVariablesRepository(database)
	service := services.NewMapVariablesService(repo)
	cnt := controllers.NewMapVariablesController(service, metrics, logger, ctx)

	group := e.Group(v1.MapVariablesPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.GET(v1.IDVar, cnt.GetMapVariableByID)
	group.PUT(v1.IDVar, cnt.UpdateMapVariable)
	group.DELETE(v1.IDVar, cnt.DeleteMapVariable)
	group.POST("", cnt.CreateMapVariable)
	group.POST("/book"+v1.IDVar, cnt.GetMapVariablesByBookID)
	group.POST(v1.ChapterSubPath+v1.IDVar, cnt.GetMapVariablesByChapterID)
	group.POST("/page"+v1.IDVar, cnt.GetMapVariablesByPageID)
	group.POST("/paragraph"+v1.IDVar, cnt.GetMapVariablesByParagraphID)
}
