package controllers

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/errors"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type pageService interface {
	CreatePage(ctx context.Context, request models.CreatePageRequest) (*models.Page, error)
	GetPageByID(ctx context.Context, id string) (*models.Page, error)
	UpdatePage(ctx context.Context, id string, request models.UpdatePageRequest) (*models.Page, error)
	DeletePage(ctx context.Context, id string) (*models.Page, error)
	ListPages(ctx context.Context, page models.RequestPage) ([]*models.Page, error)
	GetPagesByChapterID(ctx context.Context, chapterID string, page models.RequestPage) ([]*models.Page, error)
}

type PageController struct {
	Service pageService
	Metrics metrics.Metrics
	Logger  *slog.Logger
	Ctx     context.Context
}

func NewPageController(srv pageService, metric metrics.Metrics, logger *slog.Logger, ctx context.Context) *PageController {
	return &PageController{Service: srv, Metrics: metric, Logger: logger, Ctx: ctx}
}

func (p *PageController) ListPages(c echo.Context) error {
	var request models.RequestPage
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	pages, err := p.Service.ListPages(p.Ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}
	return c.JSON(http.StatusOK, models.WebResponse[[]*models.Page]{Data: pages, Status: "ok"})
}

func (p *PageController) CreatePage(c echo.Context) error {
	var request models.CreatePageRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	page, err := p.Service.CreatePage(p.Ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.PageNotCreated)
	}
	return c.JSON(http.StatusCreated, models.WebResponse[*models.Page]{Data: page, Status: "created"})
}

func (p *PageController) GetPageByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	page, err := p.Service.GetPageByID(p.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.BookNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.Page]{Data: page, Status: "ok"})
}

func (p *PageController) UpdatePage(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.UpdatePageRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	page, err := p.Service.UpdatePage(p.Ctx, id, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Page]{Data: page, Status: "updated"})
}

func (p *PageController) DeletePage(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	page, err := p.Service.DeletePage(p.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.PageNotDeleted)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Page]{Data: page, Status: "deleted"})
}

func (p *PageController) GetPagesByChapterID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.RequestPage
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	pages, err := p.Service.GetPagesByChapterID(p.Ctx, id, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.BookNotFound)
	}
	return c.JSON(http.StatusOK, models.WebResponse[[]*models.Page]{Data: pages, Status: "ok"})
}
