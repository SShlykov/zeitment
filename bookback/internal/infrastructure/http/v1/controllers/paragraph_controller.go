package controllers

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/errors"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type paragraphService interface {
	CreateParagraph(ctx context.Context, paragraph models.CreateParagraphRequest) (*models.Paragraph, error)
	GetParagraphByID(ctx context.Context, id string) (*models.Paragraph, error)
	UpdateParagraph(ctx context.Context, id string, paragraph models.UpdateParagraphRequest) (*models.Paragraph, error)
	DeleteParagraph(ctx context.Context, id string) (*models.Paragraph, error)
	ListParagraphs(ctx context.Context, request models.RequestParagraph) ([]*models.Paragraph, error)
	GetParagraphsByPageID(ctx context.Context, pageID string, request models.RequestParagraph) ([]*models.Paragraph, error)
}

type ParagraphController struct {
	Service paragraphService
	Metrics metrics.Metrics
	Logger  loggerPkg.Logger
	Ctx     context.Context
}

// NewParagraphController создает новый экземпляр ParagraphController.
func NewParagraphController(srv paragraphService, metric metrics.Metrics, logger loggerPkg.Logger, ctx context.Context) *ParagraphController {
	return &ParagraphController{Service: srv, Metrics: metric, Logger: logger, Ctx: ctx}
}

func (p *ParagraphController) ListParagraphs(c echo.Context) error {
	var request models.RequestParagraph
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	paragraphs, err := p.Service.ListParagraphs(p.Ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}
	return c.JSON(http.StatusOK, models.WebResponse[[]*models.Paragraph]{Data: paragraphs, Status: "ok"})
}

func (p *ParagraphController) CreateParagraph(c echo.Context) error {
	var request models.CreateParagraphRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	paragraph, err := p.Service.CreateParagraph(p.Ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ParagraphNotCreated)
	}
	return c.JSON(http.StatusCreated, models.WebResponse[*models.Paragraph]{Data: paragraph, Status: "created"})
}

func (p *ParagraphController) GetParagraphByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	paragraph, err := p.Service.GetParagraphByID(p.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ParagraphNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.Paragraph]{Data: paragraph, Status: "ok"})
}

func (p *ParagraphController) UpdateParagraph(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.UpdateParagraphRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	paragraph, err := p.Service.UpdateParagraph(p.Ctx, id, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Paragraph]{Data: paragraph, Status: "updated"})
}

func (p *ParagraphController) DeleteParagraph(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	paragraph, err := p.Service.DeleteParagraph(p.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ParagraphNotDeleted)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Paragraph]{Data: paragraph, Status: "deleted"})
}

func (p *ParagraphController) GetParagraphsByPageID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.RequestParagraph
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	paragraphs, err := p.Service.GetParagraphsByPageID(p.Ctx, id, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}
	return c.JSON(http.StatusOK, models.WebResponse[[]*models.Paragraph]{Data: paragraphs, Status: "ok"})
}
