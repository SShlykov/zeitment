package paragraph

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/adapters/db/postgres/bookrepo"
	"github.com/SShlykov/zeitment/bookback/internal/controller/http/httpmiddlewares"
	paragraphrepo "github.com/SShlykov/zeitment/bookback/internal/domain/services/paragraph"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func SetParagraphController(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	service := paragraphrepo.NewService(bookrepo.NewRepository(database))
	controller := NewController(service, metrics, logger, ctx)

	controller.RegisterRoutes(e)
}

func (p *Controller) RegisterRoutes(e *echo.Echo) {
	group := e.Group(PathPrefix)
	group.Use(httpmiddlewares.MetricsLogger(p.Metrics))

	group.GET("", p.ListParagraphs)
	group.POST("", p.CreateParagraph)
	group.GET("/:id", p.GetParagraphByID)
	group.PUT("/:id", p.UpdateParagraph)
	group.DELETE("/:id", p.DeleteParagraph)
	group.GET("/pages/:id", p.GetParagraphsByPageID)
}
