package controllers

import (
	"context"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/errors"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type mapVariablesService interface {
	CreateMapVariable(ctx context.Context, request models.CreateMapVariableRequest) (*models.MapVariable, error)
	GetMapVariableByID(ctx context.Context, id string) (*models.MapVariable, error)
	UpdateMapVariable(ctx context.Context, id string, request models.UpdateMapVariableRequest) (*models.MapVariable, error)
	DeleteMapVariable(ctx context.Context, id string) (*models.MapVariable, error)
	GetMapVariablesByBookID(ctx context.Context, mapID string) ([]*models.MapVariable, error)
	GetMapVariablesByChapterID(ctx context.Context, chapterID string) ([]*models.MapVariable, error)
	GetMapVariablesByPageID(ctx context.Context, pageID string) ([]*models.MapVariable, error)
	GetMapVariablesByParagraphID(ctx context.Context, paragraphID string) ([]*models.MapVariable, error)
}

// MapVariablesController структура для HTTP-контроллера книг.
type MapVariablesController struct {
	Service mapVariablesService
	Metrics metrics.Metrics
	Logger  *slog.Logger
	Ctx     context.Context
}

// NewMapVariablesController создает новый экземпляр MapVariablesController.
func NewMapVariablesController(srv mapVariablesService, metric metrics.Metrics, logger *slog.Logger, ctx context.Context) *MapVariablesController {
	return &MapVariablesController{Service: srv, Metrics: metric, Logger: logger, Ctx: ctx}
}

func (mvc *MapVariablesController) GetMapVariableByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	variable, err := mvc.Service.GetMapVariableByID(mvc.Ctx, id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, errors.MapVariablesNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.MapVariable]{Data: variable, Status: "ok"})
}

func (mvc *MapVariablesController) UpdateMapVariable(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.UpdateMapVariableRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	variable, err := mvc.Service.UpdateMapVariable(mvc.Ctx, id, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.MapVariable]{Data: variable, Status: "ok"})
}

func (mvc *MapVariablesController) DeleteMapVariable(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	variable, err := mvc.Service.DeleteMapVariable(mvc.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.MapVariablesNotDeleted)
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.MapVariable]{Data: variable, Status: "ok"})
}

func (mvc *MapVariablesController) GetMapVariablesByBookID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	variables, err := mvc.Service.GetMapVariablesByBookID(mvc.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.MapVariablesNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[[]*models.MapVariable]{Data: variables, Status: "ok"})
}

func (mvc *MapVariablesController) GetMapVariablesByChapterID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	variables, err := mvc.Service.GetMapVariablesByChapterID(mvc.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.MapVariablesNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[[]*models.MapVariable]{Data: variables, Status: "ok"})
}

func (mvc *MapVariablesController) GetMapVariablesByPageID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	variables, err := mvc.Service.GetMapVariablesByPageID(mvc.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.MapVariablesNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[[]*models.MapVariable]{Data: variables, Status: "ok"})
}

func (mvc *MapVariablesController) GetMapVariablesByParagraphID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	variables, err := mvc.Service.GetMapVariablesByParagraphID(mvc.Ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.MapVariablesNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[[]*models.MapVariable]{Data: variables, Status: "ok"})
}

func (mvc *MapVariablesController) CreateMapVariable(c echo.Context) error {
	var request models.CreateMapVariableRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	variable, err := mvc.Service.CreateMapVariable(mvc.Ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.MapVariablesNotCreated)
	}

	return c.JSON(http.StatusCreated, models.WebResponse[*models.MapVariable]{Data: variable, Status: "ok"})
}
