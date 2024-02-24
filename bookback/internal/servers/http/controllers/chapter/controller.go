package chapter

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/services/chapter"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	Service chapter.Service
}

func NewController(srv chapter.Service) *Controller {
	return &Controller{Service: srv}
}

func (ch *Controller) RegisterRoutes(e *echo.Echo, ctx context.Context) {
	e.GET("/api/v1/chapters", func(c echo.Context) error { return ch.ListChapters(c, ctx) })
	e.POST("/api/v1/chapters", func(c echo.Context) error { return ch.CreateChapter(c, ctx) })
	e.GET("/api/v1/chapters/:id", func(c echo.Context) error { return ch.GetChapterByID(c, ctx) })
	e.PUT("/api/v1/chapters/:id", func(c echo.Context) error { return ch.UpdateChapter(c, ctx) })
	e.DELETE("/api/v1/chapters/:id", func(c echo.Context) error { return ch.DeleteChapter(c, ctx) })

	e.GET("/api/v1/chapters/book/:id", func(c echo.Context) error { return ch.GetChapterByBookID(c, ctx) })
}

// ListChapters список глав
// @router /chapters [get]
// @summary Получить список глав
// @description Извлекает список всех глав
// @tags Главы
// @produce  application/json
// @success 200 {array} models.Chapter
// @failure 500 {object} config.HTTPError
func (ch *Controller) ListChapters(c echo.Context, ctx context.Context) error {
	chapters, err := ch.Service.ListChapters(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, config.ErrorForbidden)
	}
	return c.JSON(http.StatusOK, chapters)
}

// CreateChapter создание новой главы
// @router /chapters [post]
// @summary Создать главу
// @description Создает новую главу
// @tags Главы
// @accept application/json
// @produce application/json
// @param chapter body models.Chapter true "Chapter object"
// @success 201 {object} models.Chapter
// @failure 400 {object} config.HTTPError
func (ch *Controller) CreateChapter(c echo.Context, ctx context.Context) error {
	var chap models.Chapter
	if err := c.Bind(&chap); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}
	createdChapter, err := ch.Service.CreateChapter(ctx, &chap)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, config.ErrorNotCreated)
	}
	return c.JSON(http.StatusCreated, createdChapter)
}

// GetChapterByID получение главы по ID
// @router /chapters/{id} [get]
// @summary Получить главу по ID
// @description Извлекает главу по ее ID
// @tags Главы
// @param id path string true "ID главы"
// @produce application/json
// @success 200 {object} models.Chapter
// @failure 404 {object} config.HTTPError
func (ch *Controller) GetChapterByID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	chapt, err := ch.Service.GetChapterByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}
	return c.JSON(http.StatusOK, chapt)
}

// UpdateChapter обновление главы
// @router /chapters/{id} [put]
// @summary Обновить главу
// @description Обновляет главу по ее ID
// @tags Главы
// @accept application/json
// @produce application/json
// @param id path string true "ID главы"
// @param chapter body models.Chapter true "Chapter object"
// @success 200 {object} models.Chapter
// @failure 400 {object} config.HTTPError
func (ch *Controller) UpdateChapter(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	var chap models.Chapter
	if err := c.Bind(&chap); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}
	updatedChapter, err := ch.Service.UpdateChapter(ctx, id, &chap)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, config.ErrorNotUpdated)
	}
	return c.JSON(http.StatusOK, updatedChapter)
}

// DeleteChapter удаление главы
// @router /chapters/{id} [delete]
// @summary Удалить главу
// @description Удаляет главу по ее ID
// @tags Главы
// @param id path string true "ID главы"
// @produce application/json
// @success 200 {object} models.Chapter
// @failure 406 {object} config.HTTPError
func (ch *Controller) DeleteChapter(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	chapt, err := ch.Service.DeleteChapter(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, config.ErrorNotDeleted)
	}
	return c.JSON(http.StatusOK, chapt)
}

// GetChapterByBookID получение глав по ID книги
// @router /chapters/book/{id} [get]
// @summary Получить главы по ID книги
// @description Извлекает главы по ID книги
// @tags Главы
// @param id path string true "ID книги"
// @produce application/json
// @success 200 {array} models.Chapter
// @failure 404 {object} config.HTTPError
func (ch *Controller) GetChapterByBookID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	chapters, err := ch.Service.GetChapterByBookID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}
	return c.JSON(http.StatusOK, chapters)
}
