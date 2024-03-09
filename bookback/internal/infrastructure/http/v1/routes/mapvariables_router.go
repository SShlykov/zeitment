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

func MapVariables(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewMapVariablesRepository(database)
	service := services.NewMapVariablesService(repo)
	cnt := controllers.NewMapVariablesController(service, metrics, logger, ctx)

	group := e.Group(MapVariablesPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.GET("/:id", cnt.GetMapVariableByID)
	group.PUT("/:id", cnt.UpdateMapVariable)
	group.DELETE("/:id", cnt.DeleteMapVariable)
	group.POST("", cnt.CreateMapVariable)
	group.POST("/book/:id", cnt.GetMapVariablesByBookID)
	group.POST("/chapter/:id", cnt.GetMapVariablesByChapterID)
	group.POST("/page/:id", cnt.GetMapVariablesByPageID)
	group.POST("/paragraph/:id", cnt.GetMapVariablesByParagraphID)
}
