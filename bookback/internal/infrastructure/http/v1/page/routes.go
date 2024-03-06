package page

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services/page"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func SetPageController(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewPageRepository(database)
	service := page.NewService(repo)
	controller := NewController(service, metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

func (p *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(PathPrefix)
	group.Use(middleware.MetricsLogger(p.Metrics))

	group.GET("", p.ListPages)
	group.POST("", p.CreatePage)
	group.GET("/:id", p.GetPageByID)
	group.PUT("/:id", p.UpdatePage)
	group.DELETE("/:id", p.DeletePage)
	group.GET("/chapters/:id", p.GetPagesByChapterID)
}
