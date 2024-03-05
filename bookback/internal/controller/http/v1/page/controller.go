package page

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	service "github.com/SShlykov/zeitment/bookback/internal/domain/services/page"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Controller struct {
	Service service.Service
	Metrics metrics.Metrics
	Logger  *slog.Logger
	Ctx     context.Context
}

func NewController(srv service.Service, metric metrics.Metrics, logger *slog.Logger, ctx context.Context) *Controller {
	return &Controller{Service: srv, Metrics: metric, Logger: logger, Ctx: ctx}
}

// ListPages список страниц
// @router /pages [get]
// @summary Получить список страниц
// @description Извлекает список всех страниц
// @tags Страницы
// @produce  application/json
// @success 200 {array} entity.Page
// @failure 500 {object} config.HTTPError
func (p *Controller) ListPages(c echo.Context) error {
	pages, err := p.Service.ListPages(p.Ctx)
	if err != nil {
		return ErrorUnknown
	}
	return c.JSON(http.StatusOK, responseListModel{Status: "ok", Pages: pages})
}

// CreatePage создание новой страницы
// @router /pages [post]
// @summary Создать страницу
// @description Создает новую страницу
// @tags Страницы
// @accept application/json
// @produce application/json
// @param page body entity.Page true "Page object"
// @success 201 {object} entity.Page
// @failure 400 {object} config.HTTPError
func (p *Controller) CreatePage(c echo.Context) error {
	var request requestModel
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	createdPage, err := p.Service.CreatePage(p.Ctx, request.Page)
	if err != nil {
		return ErrorPageNotCreated
	}
	return c.JSON(http.StatusCreated, responseSingleModel{Status: "created", Page: createdPage})
}

// GetPageByID получение страницы по ID
// @router /pages/{id} [get]
// @summary Получить страницу по ID
// @description Извлекает страницу по ее ID
// @tags Страницы
// @produce  application/json
// @param id path string true "ID страницы"
// @success 200 {object} entity.Page
// @failure 404 {object} config.HTTPError
func (p *Controller) GetPageByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	page, err := p.Service.GetPageByID(p.Ctx, id)
	if err != nil {
		return ErrorPageNotFound
	}

	return c.JSON(http.StatusOK, responseSingleModel{Status: "ok", Page: page})
}

// UpdatePage обновление страницы
// @router /pages/{id} [put]
// @summary Обновить страницу
// @description Обновляет страницу
// @tags Страницы
// @accept application/json
// @produce application/json
// @param id path string true "ID страницы"
// @param page body entity.Page true "Page object"
// @success 200 {object} entity.Page
// @failure 400 {object} config.HTTPError
func (p *Controller) UpdatePage(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	var request requestModel
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	updatedPage, err := p.Service.UpdatePage(p.Ctx, id, request.Page)
	if err != nil {
		if errors.Is(err, config.ErrorNotFound) {
			return ErrorPageNotFound
		}
		return ErrorUnknown
	}
	return c.JSON(http.StatusOK, responseSingleModel{Status: "updated", Page: updatedPage})
}

// DeletePage удаление страницы
// @router /pages/{id} [delete]
// @summary Удалить страницу
// @description Удаляет страницу
// @tags Страницы
// @param id path string true "ID страницы"
// @produce application/json
// @success 200 {object} entity.Page
// @failure 500 {object} config.HTTPError
func (p *Controller) DeletePage(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	deletedPage, err := p.Service.DeletePage(p.Ctx, id)
	if err != nil {
		return ErrorDeletePage
	}
	return c.JSON(http.StatusOK, responseSingleModel{Status: "deleted", Page: deletedPage})
}

// GetPagesByChapterID получение страниц по ID главы
// @router /chapters/{id}/pages [get]
// @summary Получить страницы по ID главы
// @description Извлекает страницы по ID главы
// @tags Страницы
// @produce  application/json
// @param id path string true "ID главы"
// @success 200 {array} entity.Page
// @failure 404 {object} config.HTTPError
func (p *Controller) GetPagesByChapterID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	pages, err := p.Service.GetPagesByChapterID(p.Ctx, id)
	if err != nil {
		return ErrorPageNotFound
	}
	return c.JSON(http.StatusOK, responseListModel{Status: "ok", Pages: pages})
}
