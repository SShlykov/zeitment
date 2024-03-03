package chapter

import (
	"context"
	"errors"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/services/chapter"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Controller struct {
	Service chapter.Service
	Metrics metrics.Metrics
	Logger  *slog.Logger
	Ctx     context.Context
}

// NewController создает новый экземпляр Controller.
func NewController(srv chapter.Service, metric metrics.Metrics, logger *slog.Logger, ctx context.Context) *Controller {
	return &Controller{Service: srv, Metrics: metric, Logger: logger, Ctx: ctx}
}

// ListChapters список глав
// @router /chapters [get]
// @summary Получить список глав
// @description Извлекает список всех глав
// @tags Главы
// @produce  application/json
// @success 200 {array} models.Chapter
// @failure 500 {object} config.HTTPError
func (ch *Controller) ListChapters(c echo.Context) error {
	chapters, err := ch.Service.ListChapters(ch.Ctx)
	if err != nil {
		return ErrorUnknown
	}
	return c.JSON(http.StatusOK, responseListModel{Status: "ok", Chapters: chapters})
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
func (ch *Controller) CreateChapter(c echo.Context) error {
	var chap requestModel
	if err := c.Bind(&chap); err != nil {
		return ErrorValidationFailed
	}

	createdChapter, err := ch.Service.CreateChapter(ch.Ctx, chap.Chapter)
	if err != nil {
		return ErrorChapterNotCreated
	}
	return c.JSON(http.StatusCreated, responseSingleModel{Status: "created", Chapter: createdChapter})
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
func (ch *Controller) GetChapterByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	chapt, err := ch.Service.GetChapterByID(ch.Ctx, id)
	if err != nil {
		return ErrorChapterNotFound
	}
	return c.JSON(http.StatusOK, responseSingleModel{Status: "ok", Chapter: chapt})
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
func (ch *Controller) UpdateChapter(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	var request requestModel
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}

	updatedChapter, err := ch.Service.UpdateChapter(ch.Ctx, id, request.Chapter)
	if err != nil {
		if errors.Is(err, config.ErrorNotFound) {
			return ErrorChapterNotFound
		}
		fmt.Println(err)
		return ErrorUnknown
	}
	return c.JSON(http.StatusOK, responseSingleModel{Status: "updated", Chapter: updatedChapter})
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
func (ch *Controller) DeleteChapter(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	chapt, err := ch.Service.DeleteChapter(ch.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, config.ErrorNotDeleted)
	}
	return c.JSON(http.StatusOK, responseSingleModel{Status: "deleted", Chapter: chapt})
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
func (ch *Controller) GetChapterByBookID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	chapters, err := ch.Service.GetChapterByBookID(ch.Ctx, id)
	if err != nil {
		return ErrorDeleteChapter
	}
	return c.JSON(http.StatusOK, responseListModel{Status: "ok", Chapters: chapters})
}
