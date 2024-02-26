package bookevents

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	service "github.com/SShlykov/zeitment/bookback/internal/services/bookevents"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (bec *Controller) RegisterRoutes(e *echo.Echo, ctx context.Context) {
	e.GET("/api/v1/bookevents/:id", func(c echo.Context) error { return bec.GetBookEventByID(c, ctx) })
	e.PUT("/api/v1/bookevents/:id", func(c echo.Context) error { return bec.UpdateBookEvent(c, ctx) })
	e.DELETE("/api/v1/bookevents/:id", func(c echo.Context) error { return bec.DeleteBookEvent(c, ctx) })
	e.POST("/api/v1/bookevents", func(c echo.Context) error { return bec.CreateBookEvent(c, ctx) })
	e.GET("/api/v1/bookevents/book/:id", func(c echo.Context) error { return bec.GetBookEventsByBookID(c, ctx) })
	e.GET("/api/v1/bookevents/chapter/:id", func(c echo.Context) error { return bec.GetBookEventsByChapterID(c, ctx) })
	e.GET("/api/v1/bookevents/page/:id", func(c echo.Context) error { return bec.GetBookEventsByPageID(c, ctx) })
	e.GET("/api/v1/bookevents/paragraph/:id", func(c echo.Context) error { return bec.GetBookEventsByParagraphID(c, ctx) })
}

// GetBookEventByID обрабатывает запросы на получение события книги по идентификатору.
// @router /bookevents/{id} [get]
// @summary Получить событие книги по идентификатору
// @description Извлекает событие книги по идентификатору
// @tags События книги
// @produce  application/json
// @param id path string true "ID события книги"
// @success 200 {object} models.BookEvent
// @failure 404 {object} config.HTTPError
func (bec *Controller) GetBookEventByID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	event, err := bec.service.GetBookEventByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}

	return c.JSON(http.StatusOK, event)
}

// UpdateBookEvent обрабатывает запросы на обновление события книги.
// @router /bookevents/{id} [put]
// @summary Обновить событие книги
// @description Обновляет событие книги
// @tags События книги
// @accept application/json
// @produce application/json
// @param id path string true "ID события книги"
// @param event body models.BookEvent true "BookEvent object"
// @success 200 {object} models.BookEvent
// @failure 400 {object} config.HTTPError
func (bec *Controller) UpdateBookEvent(c echo.Context, ctx context.Context) error {
	var event models.BookEvent
	if err := c.Bind(&event); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}

	id := c.Param("id")
	updatedEvent, err := bec.service.UpdateBookEvent(ctx, id, &event)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, config.ErrorForbidden)
	}

	return c.JSON(http.StatusOK, updatedEvent)
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
func (bec *Controller) DeleteBookEvent(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	if _, err := bec.service.DeleteBookEvent(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}

// CreateBookEvent обрабатывает запросы на создание события книги.
// @router /bookevents [post]
// @summary Создать событие книги
// @description Создает событие книги
// @tags События книги
// @accept application/json
// @produce application/json
// @param event body models.BookEvent true "BookEvent object"
// @success 201 {object} models.BookEvent
// @failure 400 {object} config.HTTPError
func (bec *Controller) CreateBookEvent(c echo.Context, ctx context.Context) error {
	var event models.BookEvent
	if err := c.Bind(&event); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}
	createdEvent, err := bec.service.CreateBookEvent(ctx, &event)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, config.ErrorNotCreated)
	}
	return c.JSON(http.StatusCreated, createdEvent)
}

// GetBookEventsByBookID обрабатывает запросы на получение событий книги по ID книги.
// @router /bookevents/book/{id} [get]
// @summary Получить события книги по ID книги
// @description Извлекает события книги по ID книги
// @tags События книги
// @produce application/json
// @param id path string true "ID книги"
// @success 200 {object} []models.BookEvent
// @failure 404 {object} config.HTTPError
func (bec *Controller) GetBookEventsByBookID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	events, err := bec.service.GetBookEventsByBookID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}
	return c.JSON(http.StatusOK, events)
}

// GetBookEventsByChapterID обрабатывает запросы на получение событий книги по ID главы.
// @router /bookevents/chapter/{id} [get]
// @summary Получить события книги по ID главы
// @description Извлекает события книги по ID главы
// @tags События книги
// @produce application/json
// @param id path string true "ID главы"
// @success 200 {object} []models.BookEvent
// @failure 404 {object} config.HTTPError
func (bec *Controller) GetBookEventsByChapterID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	events, err := bec.service.GetBookEventsByChapterID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}
	return c.JSON(http.StatusOK, events)
}

// GetBookEventsByPageID обрабатывает запросы на получение событий книги по ID страницы.
// @router /bookevents/page/{id} [get]
// @summary Получить события книги по ID страницы
// @description Извлекает события книги по ID страницы
// @tags События книги
// @produce application/json
// @param id path string true "ID страницы"
// @success 200 {object} []models.BookEvent
// @failure 404 {object} config.HTTPError
func (bec *Controller) GetBookEventsByPageID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	events, err := bec.service.GetBookEventsByPageID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}
	return c.JSON(http.StatusOK, events)
}

// GetBookEventsByParagraphID обрабатывает запросы на получение событий книги по ID параграфа.
// @router /bookevents/paragraph/{id} [get]
// @summary Получить события книги по ID параграфа
// @description Извлекает события книги по ID параграфа
// @tags События книги
// @produce application/json
// @param id path string true "ID параграфа"
// @success 200 {object} []models.BookEvent
// @failure 404 {object} config.HTTPError
func (bec *Controller) GetBookEventsByParagraphID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	events, err := bec.service.GetBookEventsByParagraphID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}
	return c.JSON(http.StatusOK, events)
}
