package bookevents

import (
	"context"
	service "github.com/SShlykov/zeitment/bookback/internal/domain/services/bookevents"
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

// GetBookEventByID обрабатывает запросы на получение события книги по идентификатору.
// @router /bookevents/{id} [get]
// @summary Получить событие книги по идентификатору
// @description Извлекает событие книги по идентификатору
// @tags События книги
// @produce  application/json
// @param id path string true "ID события книги"
// @success 200 {object} entity.BookEvent
// @failure 404 {object} config.HTTPError
func (bec *Controller) GetBookEventByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	event, err := bec.Service.GetBookEventByID(bec.Ctx, id)
	if err != nil {
		return ErrorBookEventNotFound
	}

	return c.JSON(http.StatusOK, responseSingleModel{Status: "ok", BookEvent: event})
}

// UpdateBookEvent обрабатывает запросы на обновление события книги.
// @router /bookevents/{id} [put]
// @summary Обновить событие книги
// @description Обновляет событие книги
// @tags События книги
// @accept application/json
// @produce application/json
// @param id path string true "ID события книги"
// @param event body entity.BookEvent true "BookEvent object"
// @success 200 {object} entity.BookEvent
// @failure 400 {object} config.HTTPError
func (bec *Controller) UpdateBookEvent(c echo.Context) error {
	var request requestModel
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	updatedEvent, err := bec.Service.UpdateBookEvent(bec.Ctx, id, request.BookEvents)
	if err != nil {
		return ErrorUnknown
	}

	return c.JSON(http.StatusOK, responseSingleModel{Status: "updated", BookEvent: updatedEvent})
}

// DeleteBookEvent обрабатывает запросы на удаление события книги.
// @router /bookevents/{id} [delete]
// @summary Удалить событие книги
// @description Удаляет событие книги
// @tags События книги
// @produce application/json
// @param id path string true "ID события книги"
// @success 204
// @failure 404 {object} config.HTTPError
func (bec *Controller) DeleteBookEvent(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	deletedEvent, err := bec.Service.DeleteBookEvent(bec.Ctx, id)
	if err != nil {
		return ErrorDeleteBookEvent
	}

	return c.JSON(http.StatusOK, responseSingleModel{Status: "deleted", BookEvent: deletedEvent})
}

// CreateBookEvent обрабатывает запросы на создание события книги.
// @router /bookevents [post]
// @summary Создать событие книги
// @description Создает событие книги
// @tags События книги
// @accept application/json
// @produce application/json
// @param event body entity.BookEvent true "BookEvent object"
// @success 201 {object} entity.BookEvent
// @failure 400 {object} config.HTTPError
func (bec *Controller) CreateBookEvent(c echo.Context) error {
	var request requestModel
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	createdEvent, err := bec.Service.CreateBookEvent(bec.Ctx, request.BookEvents)
	if err != nil {
		return ErrorBookEventNotCreated
	}
	return c.JSON(http.StatusCreated, responseSingleModel{Status: "created", BookEvent: createdEvent})
}

// GetBookEventsByBookID обрабатывает запросы на получение событий книги по ID книги.
// @router /bookevents/book/{id} [get]
// @summary Получить события книги по ID книги
// @description Извлекает события книги по ID книги
// @tags События книги
// @produce application/json
// @param id path string true "ID книги"
// @success 200 {object} []entity.BookEvent
// @failure 404 {object} config.HTTPError
func (bec *Controller) GetBookEventsByBookID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	events, err := bec.Service.GetBookEventsByBookID(bec.Ctx, id)
	if err != nil {
		return ErrorBookEventNotFound
	}
	return c.JSON(http.StatusOK, responseListModel{Status: "ok", BookEvents: events})
}

// GetBookEventsByChapterID обрабатывает запросы на получение событий книги по ID главы.
// @router /bookevents/chapter/{id} [get]
// @summary Получить события книги по ID главы
// @description Извлекает события книги по ID главы
// @tags События книги
// @produce application/json
// @param id path string true "ID главы"
// @success 200 {object} []entity.BookEvent
// @failure 404 {object} config.HTTPError
func (bec *Controller) GetBookEventsByChapterID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	events, err := bec.Service.GetBookEventsByChapterID(bec.Ctx, id)
	if err != nil {
		return ErrorBookEventNotFound
	}

	return c.JSON(http.StatusOK, responseListModel{Status: "ok", BookEvents: events})
}

// GetBookEventsByPageID обрабатывает запросы на получение событий книги по ID страницы.
// @router /bookevents/page/{id} [get]
// @summary Получить события книги по ID страницы
// @description Извлекает события книги по ID страницы
// @tags События книги
// @produce application/json
// @param id path string true "ID страницы"
// @success 200 {object} []entity.BookEvent
// @failure 404 {object} config.HTTPError
func (bec *Controller) GetBookEventsByPageID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	events, err := bec.Service.GetBookEventsByPageID(bec.Ctx, id)
	if err != nil {
		return ErrorBookEventNotFound
	}

	return c.JSON(http.StatusOK, responseListModel{Status: "ok", BookEvents: events})
}

// GetBookEventsByParagraphID обрабатывает запросы на получение событий книги по ID параграфа.
// @router /bookevents/paragraph/{id} [get]
// @summary Получить события книги по ID параграфа
// @description Извлекает события книги по ID параграфа
// @tags События книги
// @produce application/json
// @param id path string true "ID параграфа"
// @success 200 {object} []entity.BookEvent
// @failure 404 {object} config.HTTPError
func (bec *Controller) GetBookEventsByParagraphID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	events, err := bec.Service.GetBookEventsByParagraphID(bec.Ctx, id)
	if err != nil {
		return ErrorBookEventNotFound
	}

	return c.JSON(http.StatusOK, responseListModel{Status: "ok", BookEvents: events})
}
