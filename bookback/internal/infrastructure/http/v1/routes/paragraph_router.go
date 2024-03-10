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

func Paragraph(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger *slog.Logger, ctx context.Context) {
	repo := pgrepo.NewParagraphRepository(database)
	service := services.NewParagraphService(repo)
	cnt := controllers.NewParagraphController(service, metrics, logger, ctx)

	group := e.Group(v1.ParagraphsPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.POST("/list", cnt.ListParagraphs)
	group.POST("", cnt.CreateParagraph)
	group.GET("/:id", cnt.GetParagraphByID)
	group.PUT("/:id", cnt.UpdateParagraph)
	group.DELETE("/:id", cnt.DeleteParagraph)
	group.POST("/pages/:id", cnt.GetParagraphsByPageID)
}
