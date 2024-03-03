package page

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/httpmiddlewares"
	pagerepo "github.com/SShlykov/zeitment/bookback/internal/services/page"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func SetPageController(e *echo.Echo, database db.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	service := pagerepo.NewService(pagerepo.NewRepository(database))
	controller := NewController(service, metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

func (p *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(PathPrefix)
	group.Use(httpmiddlewares.MetricsLogger(p.Metrics))

	group.GET("", p.ListPages)
	group.POST("", p.CreatePage)
	group.GET("/:id", p.GetPageByID)
	group.PUT("/:id", p.UpdatePage)
	group.DELETE("/:id", p.DeletePage)
	group.GET("/chapters/:id", p.GetPagesByChapterID)
}
