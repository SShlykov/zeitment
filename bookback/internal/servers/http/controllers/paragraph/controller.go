package paragraph

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	service "github.com/SShlykov/zeitment/bookback/internal/services/paragraph"
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
	e.GET("/api/v1/paragraphs", func(c echo.Context) error { return p.ListParagraphs(c, ctx) })
	e.POST("/api/v1/paragraphs", func(c echo.Context) error { return p.CreateParagraph(c, ctx) })
	e.GET("/api/v1/paragraphs/:id", func(c echo.Context) error { return p.GetParagraphByID(c, ctx) })
	e.PUT("/api/v1/paragraphs/:id", func(c echo.Context) error { return p.UpdateParagraph(c, ctx) })
	e.DELETE("/api/v1/paragraphs/:id", func(c echo.Context) error { return p.DeleteParagraph(c, ctx) })
}

// ListParagraphs список параграфов
// @router /paragraphs [get]
// @summary Получить список параграфов
// @description Извлекает список всех параграфов
// @tags Параграфы
// @produce  application/json
// @success 200 {array} models.Paragraph
// @failure 500 {object} config.HTTPError
func (p *Controller) ListParagraphs(c echo.Context, ctx context.Context) error {
	paragraphs, err := p.Service.ListParagraphs(ctx, "")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, config.ErrorForbidden)
	}
	return c.JSON(http.StatusOK, paragraphs)
}

// CreateParagraph создание нового параграфа
// @router /paragraphs [post]
// @summary Создать параграф
// @description Создает новый параграф
// @tags Параграфы
// @accept application/json
// @produce application/json
// @param paragraph body models.Paragraph true "Paragraph object"
// @success 201 {object} models.Paragraph
// @failure 400 {object} config.HTTPError
func (p *Controller) CreateParagraph(c echo.Context, ctx context.Context) error {
	var paragraph models.Paragraph
	if err := c.Bind(&paragraph); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}
	createdParagraph, err := p.Service.CreateParagraph(ctx, &paragraph)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, config.ErrorNotCreated)
	}
	return c.JSON(http.StatusCreated, createdParagraph)
}

// GetParagraphByID получение параграфа по идентификатору
// @router /paragraphs/{id} [get]
// @summary Получить параграф по ID
// @description Извлекает параграф по его ID
// @tags Параграфы
// @param id path string true "ID параграфа"
// @produce application/json
// @success 200 {object} models.Paragraph
// @failure 404 {object} config.HTTPError
func (p *Controller) GetParagraphByID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	paragraph, err := p.Service.GetParagraphByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}
	return c.JSON(http.StatusOK, paragraph)
}

// UpdateParagraph обновление параграфа
// @router /paragraphs/{id} [put]
// @summary Обновить параграф
// @description Обновляет параграф по его ID
// @tags Параграфы
// @accept application/json
// @produce application/json
// @param id path string true "ID параграфа"
// @param paragraph body models.Paragraph true "Paragraph object"
// @success 200 {object} models.Paragraph
// @failure 400 {object} config.HTTPError
func (p *Controller) UpdateParagraph(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	var paragraph models.Paragraph
	if err := c.Bind(&paragraph); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}
	updatedParagraph, err := p.Service.UpdateParagraph(ctx, id, &paragraph)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, config.ErrorNotUpdated)
	}
	return c.JSON(http.StatusOK, updatedParagraph)
}

// DeleteParagraph удаление параграфа
// @router /paragraphs/{id} [delete]
// @summary Удалить параграф
// @description Удаляет параграф по его ID
// @tags Параграфы
// @param id path string true "ID параграфа"
// @success 200 {object} models.Paragraph
// @failure 500 {object} config.HTTPError
func (p *Controller) DeleteParagraph(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	deletedParagraph, err := p.Service.DeleteParagraph(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, config.ErrorNotDeleted)
	}
	return c.JSON(http.StatusOK, deletedParagraph)
}
