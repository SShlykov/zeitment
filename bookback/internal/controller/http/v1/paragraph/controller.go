package paragraph

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	service "github.com/SShlykov/zeitment/bookback/internal/domain/services/paragraph"
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

// NewController создает новый экземпляр Controller.
func NewController(srv service.Service, metric metrics.Metrics, logger *slog.Logger, ctx context.Context) *Controller {
	return &Controller{Service: srv, Metrics: metric, Logger: logger, Ctx: ctx}
}

// ListParagraphs список параграфов
// @router /paragraphs [get]
// @summary Получить список параграфов
// @description Извлекает список всех параграфов
// @tags Параграфы
// @produce  application/json
// @success 200 {array} entity.Paragraph
// @failure 500 {object} config.HTTPError
func (p *Controller) ListParagraphs(c echo.Context) error {
	paragraphs, err := p.Service.ListParagraphs(p.Ctx)
	if err != nil {
		return ErrorUnknown
	}
	return c.JSON(http.StatusOK, responseListModel{Status: "ok", Paragraphs: paragraphs})
}

// CreateParagraph создание нового параграфа
// @router /paragraphs [post]
// @summary Создать параграф
// @description Создает новый параграф
// @tags Параграфы
// @accept application/json
// @produce application/json
// @param paragraph body entity.Paragraph true "Paragraph object"
// @success 201 {object} entity.Paragraph
// @failure 400 {object} config.HTTPError
func (p *Controller) CreateParagraph(c echo.Context) error {
	var request requestModel
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	createdParagraph, err := p.Service.CreateParagraph(p.Ctx, request.Paragraph)
	if err != nil {
		return ErrorParagraphNotFound
	}
	return c.JSON(http.StatusCreated, responseSingleModel{Status: "created", Paragraph: createdParagraph})
}

// GetParagraphByID получение параграфа по идентификатору
// @router /paragraphs/{id} [get]
// @summary Получить параграф по ID
// @description Извлекает параграф по его ID
// @tags Параграфы
// @param id path string true "ID параграфа"
// @produce application/json
// @success 200 {object} entity.Paragraph
// @failure 404 {object} config.HTTPError
func (p *Controller) GetParagraphByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	paragraph, err := p.Service.GetParagraphByID(p.Ctx, id)
	if err != nil {
		return ErrorParagraphNotFound
	}

	return c.JSON(http.StatusOK, responseSingleModel{Status: "ok", Paragraph: paragraph})
}

// UpdateParagraph обновление параграфа
// @router /paragraphs/{id} [put]
// @summary Обновить параграф
// @description Обновляет параграф по его ID
// @tags Параграфы
// @accept application/json
// @produce application/json
// @param id path string true "ID параграфа"
// @param paragraph body entity.Paragraph true "Paragraph object"
// @success 200 {object} entity.Paragraph
// @failure 400 {object} config.HTTPError
func (p *Controller) UpdateParagraph(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	var request requestModel
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	updatedParagraph, err := p.Service.UpdateParagraph(p.Ctx, id, request.Paragraph)
	if err != nil {
		if errors.Is(err, config.ErrorNotFound) {
			return ErrorParagraphNotFound
		}
		return ErrorUnknown
	}
	return c.JSON(http.StatusOK, responseSingleModel{Status: "updated", Paragraph: updatedParagraph})
}

// DeleteParagraph удаление параграфа
// @router /paragraphs/{id} [delete]
// @summary Удалить параграф
// @description Удаляет параграф по его ID
// @tags Параграфы
// @param id path string true "ID параграфа"
// @success 200 {object} entity.Paragraph
// @failure 500 {object} config.HTTPError
func (p *Controller) DeleteParagraph(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	deletedParagraph, err := p.Service.DeleteParagraph(p.Ctx, id)
	if err != nil {
		return ErrorDeleteParagraph
	}
	return c.JSON(http.StatusOK, responseSingleModel{Status: "deleted", Paragraph: deletedParagraph})
}

// GetParagraphsByPageID получение параграфов по ID страницы
// @router /paragraphs/pages/{id} [get]
// @summary Получить параграфы по ID страницы
// @description Извлекает параграфы по ID страницы
// @tags Параграфы
// @param id path string true "ID страницы"
// @produce application/json
// @success 200 {array} entity.Paragraph
// @failure 404 {object} config.HTTPError
func (p *Controller) GetParagraphsByPageID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	paragraphs, err := p.Service.GetParagraphsByPageID(p.Ctx, id)
	if err != nil {
		return ErrorParagraphNotFound
	}
	return c.JSON(http.StatusOK, responseListModel{Status: "ok", Paragraphs: paragraphs})
}
