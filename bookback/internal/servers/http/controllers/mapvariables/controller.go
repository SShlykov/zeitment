package mapvariables

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	service "github.com/SShlykov/zeitment/bookback/internal/services/mapvariables"
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

func (mvc *Controller) RegisterRoutes(e *echo.Echo, ctx context.Context) {
	e.GET("/api/v1/mapvariables/:id", func(c echo.Context) error { return mvc.GetMapVariableByID(c, ctx) })
	e.PUT("/api/v1/mapvariables/:id", func(c echo.Context) error { return mvc.UpdateMapVariable(c, ctx) })
	e.DELETE("/api/v1/mapvariables/:id", func(c echo.Context) error { return mvc.DeleteMapVariable(c, ctx) })
	e.POST("/api/v1/mapvariables", func(c echo.Context) error { return mvc.CreateMapVariable(c, ctx) })
	e.GET("/api/v1/mapvariables/book/:id", func(c echo.Context) error { return mvc.GetMapVariablesByBookID(c, ctx) })
	e.GET("/api/v1/mapvariables/chapter/:id", func(c echo.Context) error { return mvc.GetMapVariablesByChapterID(c, ctx) })
	e.GET("/api/v1/mapvariables/page/:id", func(c echo.Context) error { return mvc.GetMapVariablesByPageID(c, ctx) })
	e.GET("/api/v1/mapvariables/paragraph/:id", func(c echo.Context) error { return mvc.GetMapVariablesByParagraphID(c, ctx) })
}

// GetMapVariableByID обрабатывает запросы на получение переменной карты по идентификатору.
// @router /mapvariables/{id} [get]
// @summary Получить переменную карты по идентификатору
// @description Извлекает переменную карты по идентификатору
// @tags Переменные карты
// @produce  application/json
// @param id path string true "ID переменной карты"
// @success 200 {object} models.MapVariable
// @failure 404 {object} config.HTTPError
func (mvc *Controller) GetMapVariableByID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	variable, err := mvc.service.GetMapVariableByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}

	return c.JSON(http.StatusOK, variable)
}

// UpdateMapVariable обрабатывает запросы на обновление переменной карты.
// @router /mapvariables/{id} [put]
// @summary Обновить переменную карты
// @description Обновляет переменную карты
// @tags Переменные карты
// @accept application/json
// @produce application/json
// @param id path string true "ID переменной карты"
// @param variable body models.MapVariable true "MapVariable object"
// @success 200 {object} models.MapVariable
// @failure 400 {object} config.HTTPError
func (mvc *Controller) UpdateMapVariable(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	var variable models.MapVariable
	if err := c.Bind(&variable); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}

	updatedVariable, err := mvc.service.UpdateMapVariable(ctx, id, &variable)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, config.ErrorForbidden)
	}

	return c.JSON(http.StatusOK, updatedVariable)
}

// DeleteMapVariable обрабатывает запросы на удаление переменной карты.
// @router /mapvariables/{id} [delete]
// @summary Удалить переменную карты
// @description Удаляет переменную карты
// @tags Переменные карты
// @param id path string true "ID переменной карты"
// @success 204
// @failure 404 {object} config.HTTPError
func (mvc *Controller) DeleteMapVariable(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	if _, err := mvc.service.DeleteMapVariable(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}

// GetMapVariablesByBookID обрабатывает запросы на получение переменных карты по идентификатору книги.
// @router /mapvariables/book/{id} [get]
// @summary Получить переменные карты по идентификатору книги
// @description Извлекает переменные карты по идентификатору книги
// @tags Переменные карты
// @produce application/json
// @param id path string true "ID книги"
// @success 200 {array} models.MapVariable
// @failure 404 {object} config.HTTPError
func (mvc *Controller) GetMapVariablesByBookID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	variables, err := mvc.service.GetMapVariablesByBookID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}

	return c.JSON(http.StatusOK, variables)
}

// GetMapVariablesByChapterID обрабатывает запросы на получение переменных карты по идентификатору главы.
// @router /mapvariables/chapter/{id} [get]
// @summary Получить переменные карты по идентификатору главы
// @description Извлекает переменные карты по идентификатору главы
// @tags Переменные карты
// @produce application/json
// @param id path string true "ID главы"
// @success 200 {array} models.MapVariable
// @failure 404 {object} config.HTTPError
func (mvc *Controller) GetMapVariablesByChapterID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	variables, err := mvc.service.GetMapVariablesByChapterID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}

	return c.JSON(http.StatusOK, variables)
}

// GetMapVariablesByPageID обрабатывает запросы на получение переменных карты по идентификатору страницы.
// @router /mapvariables/page/{id} [get]
// @summary Получить переменные карты по идентификатору страницы
// @description Извлекает переменные карты по идентификатору страницы
// @tags Переменные карты
// @produce application/json
// @param id path string true "ID страницы"
// @success 200 {array} models.MapVariable
// @failure 404 {object} config.HTTPError
func (mvc *Controller) GetMapVariablesByPageID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	variables, err := mvc.service.GetMapVariablesByPageID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}

	return c.JSON(http.StatusOK, variables)
}

// GetMapVariablesByParagraphID обрабатывает запросы на получение переменных карты по идентификатору параграфа.
// @router /mapvariables/paragraph/{id} [get]
// @summary Получить переменные карты по идентификатору параграфа
// @description Извлекает переменные карты по идентификатору параграфа
// @tags Переменные карты
// @produce application/json
// @param id path string true "ID параграфа"
// @success 200 {array} models.MapVariable
// @failure 404 {object} config.HTTPError
func (mvc *Controller) GetMapVariablesByParagraphID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	variables, err := mvc.service.GetMapVariablesByParagraphID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}

	return c.JSON(http.StatusOK, variables)
}

// CreateMapVariable обрабатывает создание новой переменной карты.
// @router /mapvariables [post]
// @summary Создать переменную карты
// @description Создает новую переменную карты
// @tags Переменные карты
// @accept application/json
// @produce application/json
// @param variable body models.MapVariable true "MapVariable object"
// @success 201 {object} models.MapVariable
// @failure 400 {object} config.HTTPError
func (mvc *Controller) CreateMapVariable(c echo.Context, ctx context.Context) error {
	var variable models.MapVariable
	if err := c.Bind(&variable); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}

	createdVariable, err := mvc.service.CreateMapVariable(ctx, &variable)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, config.ErrorForbidden)
	}

	return c.JSON(http.StatusCreated, createdVariable)
}
