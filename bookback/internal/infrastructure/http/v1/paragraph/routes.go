package paragraph

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services/paragraph"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func SetParagraphController(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewParagraphRepository(database)
	service := paragraph.NewService(repo)
	controller := NewController(service, metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

func (p *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(PathPrefix)
	group.Use(middleware.MetricsLogger(p.Metrics))

	group.GET("", p.ListParagraphs)
	group.POST("", p.CreateParagraph)
	group.GET("/:id", p.GetParagraphByID)
	group.PUT("/:id", p.UpdateParagraph)
	group.DELETE("/:id", p.DeleteParagraph)
	group.GET("/pages/:id", p.GetParagraphsByPageID)
}
