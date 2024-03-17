package controllers

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/errors"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/SShlykov/zeitment/metrics"
	"github.com/labstack/echo/v4"
	"net/http"
)

type bookEventService interface {
	CreateBookEvent(ctx context.Context, request models.CreateBookEventRequest) (*models.BookEvent, error)
	GetBookEventByID(ctx context.Context, id string) (*models.BookEvent, error)
	UpdateBookEvent(ctx context.Context, id string, request models.UpdateBookEventRequest) (*models.BookEvent, error)
	DeleteBookEvent(ctx context.Context, id string) (*models.BookEvent, error)

	GetBookEventsByBookID(ctx context.Context, bookID string, request models.RequestBookEvent) ([]*models.BookEvent, error)
	GetBookEventsByChapterID(ctx context.Context, chapterID string, request models.RequestBookEvent) ([]*models.BookEvent, error)
	GetBookEventsByPageID(ctx context.Context, pageID string, request models.RequestBookEvent) ([]*models.BookEvent, error)
	GetBookEventsByParagraphID(ctx context.Context, paragraphID string, request models.RequestBookEvent) ([]*models.BookEvent, error)

	TogglePublic(ctx context.Context, request models.ToggleBookEventRequest) (*models.BookEvent, error)
}

// BookEventController структура для HTTP-контроллера событий книги.
type BookEventController struct {
	Service bookEventService
	Metrics metrics.Metrics
	Logger  loggerPkg.Logger
	Ctx     context.Context
}

// NewBookEventController создает новый экземпляр BookEventController.
func NewBookEventController(srv bookEventService, metric metrics.Metrics,
	logger loggerPkg.Logger, ctx context.Context) *BookEventController {
	return &BookEventController{Service: srv, Metrics: metric, Logger: logger, Ctx: ctx}
}

func (bec *BookEventController) TogglePublic(c echo.Context) error {
	var request models.ToggleBookEventRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	event, err := bec.Service.TogglePublic(bec.Ctx, request)
	if err != nil {
		bec.Logger.Info("error", loggerPkg.Err(err))
		bec.Metrics.IncCounter("controller.BookEvent.TogglePublic.error", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.BookEvent]{Data: event, Status: "ok"})
}

func (bec *BookEventController) GetBookEventByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	event, err := bec.Service.GetBookEventByID(bec.Ctx, id)
	if err != nil {
		bec.Logger.Info("error", loggerPkg.Err(err))
		bec.Metrics.IncCounter("controller.BookEvent.GetBookEventByID.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.BookEventNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.BookEvent]{Data: event, Status: "ok"})
}

func (bec *BookEventController) CreateBookEvent(c echo.Context) error {
	var request models.CreateBookEventRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	createdEvent, err := bec.Service.CreateBookEvent(bec.Ctx, request)
	if err != nil {
		bec.Logger.Info("error", loggerPkg.Err(err))
		bec.Metrics.IncCounter("controller.BookEvent.CreateBookEvent.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.BookEventNotCreated)
	}
	return c.JSON(http.StatusCreated, models.WebResponse[*models.BookEvent]{Data: createdEvent, Status: "created"})
}

func (bec *BookEventController) UpdateBookEvent(c echo.Context) error {
	var request models.UpdateBookEventRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	event, err := bec.Service.UpdateBookEvent(bec.Ctx, id, request)
	if err != nil {
		bec.Logger.Info("error", loggerPkg.Err(err))
		bec.Metrics.IncCounter("controller.BookEvent.UpdateBookEvent.error", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.BookEvent]{Data: event, Status: "updated"})
}

func (bec *BookEventController) DeleteBookEvent(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	deletedEvent, err := bec.Service.DeleteBookEvent(bec.Ctx, id)
	if err != nil {
		bec.Logger.Info("error", loggerPkg.Err(err))
		bec.Metrics.IncCounter("controller.BookEvent.DeleteBookEvent.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.BookEventNotDeleted)
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.BookEvent]{Data: deletedEvent, Status: "deleted"})
}

func (bec *BookEventController) GetBookEventsByBookID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.RequestBookEvent
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	events, err := bec.Service.GetBookEventsByBookID(bec.Ctx, id, request)
	if err != nil {
		bec.Logger.Info("error", loggerPkg.Err(err))
		bec.Metrics.IncCounter("controller.BookEvent.GetBookEventsByBookID.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.BookEventNotFound)
	}
	return c.JSON(http.StatusOK, models.WebResponse[[]*models.BookEvent]{Data: events, Status: "ok"})
}

func (bec *BookEventController) GetBookEventsByChapterID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.RequestBookEvent
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	events, err := bec.Service.GetBookEventsByChapterID(bec.Ctx, id, request)
	if err != nil {
		bec.Logger.Info("error", loggerPkg.Err(err))
		bec.Metrics.IncCounter("controller.BookEvent.GetBookEventsByChapterID.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.BookEventNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[[]*models.BookEvent]{Data: events, Status: "ok"})
}

func (bec *BookEventController) GetBookEventsByPageID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.RequestBookEvent
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	events, err := bec.Service.GetBookEventsByPageID(bec.Ctx, id, request)
	if err != nil {
		bec.Logger.Info("error", loggerPkg.Err(err))
		bec.Metrics.IncCounter("controller.BookEvent.GetBookEventsByPageID.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.BookEventNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[[]*models.BookEvent]{Data: events, Status: "ok"})
}

func (bec *BookEventController) GetBookEventsByParagraphID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.RequestBookEvent
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	events, err := bec.Service.GetBookEventsByParagraphID(bec.Ctx, id, request)
	if err != nil {
		bec.Logger.Info("error", loggerPkg.Err(err))
		bec.Metrics.IncCounter("controller.BookEvent.GetBookEventsByPageID.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.BookEventNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[[]*models.BookEvent]{Data: events, Status: "ok"})
}
