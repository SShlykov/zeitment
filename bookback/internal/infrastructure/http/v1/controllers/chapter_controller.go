package controllers

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/errors"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type chapterService interface {
	CreateChapter(ctx context.Context, request models.CreateChapterRequest) (*models.Chapter, error)
	GetChapterByID(ctx context.Context, id string) (*models.Chapter, error)
	UpdateChapter(ctx context.Context, id string, request models.UpdateChapterRequest) (*models.Chapter, error)
	DeleteChapter(ctx context.Context, id string) (*models.Chapter, error)
	ListChapters(ctx context.Context, request models.RequestChapter) ([]*models.Chapter, error)
	GetChapterByBookID(ctx context.Context, bookID string, request models.RequestChapter) ([]*models.Chapter, error)
}

type ChapterController struct {
	Service chapterService
	Metrics metrics.Metrics
	Logger  *slog.Logger
	Ctx     context.Context
}

// NewChapterController создает новый экземпляр ChapterController.
func NewChapterController(srv chapterService, metric metrics.Metrics, logger *slog.Logger, ctx context.Context) *ChapterController {
	return &ChapterController{Service: srv, Metrics: metric, Logger: logger, Ctx: ctx}
}

func (ch *ChapterController) ListChapters(c echo.Context) error {
	var request models.RequestChapter
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	chapters, err := ch.Service.ListChapters(ch.Ctx, request)
	if err != nil {
		ch.Logger.Info("error", slog.String("err", err.Error()))
		ch.Metrics.IncCounter("controller.chapter.ListChapters.error", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}
	return c.JSON(http.StatusOK, models.WebResponse[[]*models.Chapter]{Data: chapters, Status: "ok"})
}

func (ch *ChapterController) CreateChapter(c echo.Context) error {
	var chap models.CreateChapterRequest
	if err := c.Bind(&chap); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	createdChapter, err := ch.Service.CreateChapter(ch.Ctx, chap)
	if err != nil {
		ch.Logger.Info("error", slog.String("err", err.Error()))
		ch.Metrics.IncCounter("controller.chapter.CreateChapter.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.ChapterNotCreated)
	}
	return c.JSON(http.StatusCreated, models.WebResponse[*models.Chapter]{Data: createdChapter, Status: "created"})
}

func (ch *ChapterController) GetChapterByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	chapter, err := ch.Service.GetChapterByID(ch.Ctx, id)
	if err != nil {
		ch.Logger.Info("error", slog.String("err", err.Error()))
		ch.Metrics.IncCounter("controller.chapter.GetChapterByID.error", err.Error())
		return echo.NewHTTPError(http.StatusNotFound, errors.ChapterNotFound)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Chapter]{Data: chapter, Status: "ok"})
}

func (ch *ChapterController) UpdateChapter(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.UpdateChapterRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed))
	}

	chapter, err := ch.Service.UpdateChapter(ch.Ctx, id, request)
	if err != nil {
		ch.Logger.Info("error", slog.String("err", err.Error()))
		ch.Metrics.IncCounter("controller.chapter.UpdateChapter.error", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Chapter]{Data: chapter, Status: "updated"})
}

func (ch *ChapterController) DeleteChapter(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	chapter, err := ch.Service.DeleteChapter(ch.Ctx, id)
	if err != nil {
		ch.Logger.Info("error", slog.String("err", err.Error()))
		ch.Metrics.IncCounter("controller.chapter.DeleteChapter.error", err.Error())
		return echo.NewHTTPError(http.StatusNotAcceptable, errors.ChapterNotDeleted)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Chapter]{Data: chapter, Status: "deleted"})
}

func (ch *ChapterController) GetChapterByBookID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.RequestChapter
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	chapters, err := ch.Service.GetChapterByBookID(ch.Ctx, id, request)
	if err != nil {
		ch.Logger.Info("error", slog.String("err", err.Error()))
		ch.Metrics.IncCounter("controller.chapter.DeleteChapter.error", err.Error())
		return echo.NewHTTPError(http.StatusNotFound, errors.ChapterNotFound)
	}
	return c.JSON(http.StatusOK, models.WebResponse[[]*models.Chapter]{Data: chapters, Status: "ok"})
}
