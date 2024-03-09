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

func Page(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewPageRepository(database)
	service := services.NewPageService(repo)
	cnt := controllers.NewPageController(service, metrics, logger, ctx)

	group := e.Group(PagesPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.POST("/list", cnt.ListPages)
	group.POST("", cnt.CreatePage)
	group.GET("/:id", cnt.GetPageByID)
	group.PUT("/:id", cnt.UpdatePage)
	group.DELETE("/:id", cnt.DeletePage)
	group.POST("/chapters/:id", cnt.GetPagesByChapterID)
}
