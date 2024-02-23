package page

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	service "github.com/SShlykov/zeitment/bookback/internal/services/page"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	Service service.Service
}

func NewController(srv service.Service) *Controller {
	return &Controller{Service: srv}
}

func (p *Controller) RegisterRoutes(e *echo.Echo, ctx context.Context) {
	e.GET("/api/v1/pages", func(c echo.Context) error { return p.ListPages(c, ctx) })
	e.POST("/api/v1/pages", func(c echo.Context) error { return p.CreatePage(c, ctx) })
	e.GET("/api/v1/pages/:id", func(c echo.Context) error { return p.GetPageByID(c, ctx) })
	e.PUT("/api/v1/pages/:id", func(c echo.Context) error { return p.UpdatePage(c, ctx) })
	e.DELETE("/api/v1/pages/:id", func(c echo.Context) error { return p.DeletePage(c, ctx) })
}

// ListPages список страниц
// @router /pages [get]
// @summary Получить список страниц
// @description Извлекает список всех страниц
// @tags Страницы
// @produce  application/json
// @success 200 {array} models.Page
// @failure 500 {object} config.HTTPError
func (p *Controller) ListPages(c echo.Context, ctx context.Context) error {
	pages, err := p.Service.ListPages(ctx, "")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, config.ErrorForbidden)
	}
	return c.JSON(http.StatusOK, pages)
}

// CreatePage создание новой страницы
// @router /pages [post]
// @summary Создать страницу
// @description Создает новую страницу
// @tags Страницы
// @accept application/json
// @produce application/json
// @param page body models.Page true "Page object"
// @success 201 {object} models.Page
// @failure 400 {object} config.HTTPError
func (p *Controller) CreatePage(c echo.Context, ctx context.Context) error {
	var page models.Page
	if err := c.Bind(&page); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}
	createdPage, err := p.Service.CreatePage(ctx, &page)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, config.ErrorNotCreated)
	}
	return c.JSON(http.StatusCreated, createdPage)
}

// GetPageByID получение страницы по ID
// @router /pages/{id} [get]
// @summary Получить страницу по ID
// @description Извлекает страницу по ее ID
// @tags Страницы
// @produce  application/json
// @param id path string true "ID страницы"
// @success 200 {object} models.Page
// @failure 404 {object} config.HTTPError
func (p *Controller) GetPageByID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	page, err := p.Service.GetPageByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}
	return c.JSON(http.StatusOK, page)
}

func (p *Controller) UpdatePage(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	var page models.Page
	if err := c.Bind(&page); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}
	updatedPage, err := p.Service.UpdatePage(ctx, id, &page)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, config.ErrorNotUpdated)
	}
	return c.JSON(http.StatusOK, updatedPage)
}

func (p *Controller) DeletePage(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	deletedPage, err := p.Service.DeletePage(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, config.ErrorNotDeleted)
	}
	return c.JSON(http.StatusOK, deletedPage)
}
