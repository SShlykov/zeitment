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

func Page(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewPageRepository(database)
	service := services.NewPageService(repo)
	cnt := controllers.NewPageController(service, metrics, logger, ctx)

	group := e.Group(v1.PagesPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.POST(v1.ListSubPath, cnt.ListPages)
	group.POST("", cnt.CreatePage)
	group.GET(v1.IDVar, cnt.GetPageByID)
	group.PUT(v1.IDVar, cnt.UpdatePage)
	group.DELETE(v1.IDVar, cnt.DeletePage)
	group.POST("/chapters"+v1.IDVar, cnt.GetPagesByChapterID)
}
