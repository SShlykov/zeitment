package mapvariables

import (
	"context"
	"errors"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	service "github.com/SShlykov/zeitment/bookback/internal/services/mapvariables"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

// Controller структура для HTTP-контроллера книг.
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

// GetMapVariableByID обрабатывает запросы на получение переменной карты по идентификатору.
// @router /mapvariables/{id} [get]
// @summary Получить переменную карты по идентификатору
// @description Извлекает переменную карты по идентификатору
// @tags Переменные карты
// @produce  application/json
// @param id path string true "ID переменной карты"
// @success 200 {object} models.MapVariable
// @failure 404 {object} config.HTTPError
func (mvc *Controller) GetMapVariableByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	variable, err := mvc.Service.GetMapVariableByID(mvc.Ctx, id)
	if err != nil {
		fmt.Println(err)
		return ErrorMapVariableNotFound
	}

	return c.JSON(http.StatusOK, responseSingleModel{Status: "ok", MapVariable: variable})
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
func (mvc *Controller) UpdateMapVariable(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	var request requestModel
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	updatedVariable, err := mvc.Service.UpdateMapVariable(mvc.Ctx, id, request.MapVariables)
	if err != nil {
		if errors.Is(err, config.ErrorNotFound) {
			return ErrorMapVariableNotFound
		}
		return ErrorUnknown
	}

	return c.JSON(http.StatusOK, responseSingleModel{Status: "updated", MapVariable: updatedVariable})
}

// DeleteMapVariable обрабатывает запросы на удаление переменной карты.
// @router /mapvariables/{id} [delete]
// @summary Удалить переменную карты
// @description Удаляет переменную карты
// @tags Переменные карты
// @param id path string true "ID переменной карты"
// @success 204
// @failure 404 {object} config.HTTPError
func (mvc *Controller) DeleteMapVariable(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	mapVariable, err := mvc.Service.DeleteMapVariable(mvc.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}

	return c.JSON(http.StatusOK, responseSingleModel{Status: "deleted", MapVariable: mapVariable})
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
func (mvc *Controller) GetMapVariablesByBookID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	variables, err := mvc.Service.GetMapVariablesByBookID(mvc.Ctx, id)
	if err != nil {
		return ErrorMapVariableNotFound
	}

	return c.JSON(http.StatusOK, responseListModel{Status: "ok", MapVariables: variables})
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
func (mvc *Controller) GetMapVariablesByChapterID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	variables, err := mvc.Service.GetMapVariablesByChapterID(mvc.Ctx, id)
	if err != nil {
		return ErrorMapVariableNotFound
	}

	return c.JSON(http.StatusOK, responseListModel{Status: "ok", MapVariables: variables})
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
func (mvc *Controller) GetMapVariablesByPageID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	variables, err := mvc.Service.GetMapVariablesByPageID(mvc.Ctx, id)
	if err != nil {
		return ErrorMapVariableNotFound
	}

	return c.JSON(http.StatusOK, responseListModel{Status: "ok", MapVariables: variables})
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
func (mvc *Controller) GetMapVariablesByParagraphID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	variables, err := mvc.Service.GetMapVariablesByParagraphID(mvc.Ctx, id)
	if err != nil {
		return ErrorMapVariableNotFound
	}

	return c.JSON(http.StatusOK, responseListModel{Status: "ok", MapVariables: variables})
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
func (mvc *Controller) CreateMapVariable(c echo.Context) error {
	var request requestModel
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	createdVariable, err := mvc.Service.CreateMapVariable(mvc.Ctx, request.MapVariables)
	if err != nil {
		return ErrorMapVariableNotCreated
	}

	return c.JSON(http.StatusCreated, responseSingleModel{Status: "created", MapVariable: createdVariable})
}
