package mapvariables

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/httpmiddlewares"
	mapvariablesrepo "github.com/SShlykov/zeitment/bookback/internal/services/mapvariables"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func SetMapVariablesController(e *echo.Echo, database db.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	service := mapvariablesrepo.NewService(mapvariablesrepo.NewRepository(database))
	controller := NewController(service, metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

func (mvc *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(PathPrefix)
	group.Use(httpmiddlewares.MetricsLogger(mvc.Metrics))

	group.GET("/:id", mvc.GetMapVariableByID)
	group.PUT("/:id", mvc.UpdateMapVariable)
	group.DELETE("/:id", mvc.DeleteMapVariable)
	group.POST("", mvc.CreateMapVariable)
	group.GET("/book/:id", mvc.GetMapVariablesByBookID)
	group.GET("/chapter/:id", mvc.GetMapVariablesByChapterID)
	group.GET("/page/:id", mvc.GetMapVariablesByPageID)
	group.GET("/paragraph/:id", mvc.GetMapVariablesByParagraphID)
}
