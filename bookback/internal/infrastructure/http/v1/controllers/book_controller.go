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

type bookService interface {
	CreateBook(ctx context.Context, request models.CreateBookRequest) (*models.Book, error)
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	UpdateBook(ctx context.Context, id string, request models.UpdateBookRequest) (*models.Book, error)
	DeleteBook(ctx context.Context, id string) (*models.Book, error)
	ListBooks(ctx context.Context, request models.RequestBook) ([]*models.Book, error)
}

// BookController структура для HTTP-контроллера книг.
type BookController struct {
	Service bookService
	Metrics metrics.Metrics
	Logger  loggerPkg.Logger
	Ctx     context.Context
}

// NewBookController создает новый экземпляр Controller.
func NewBookController(srv bookService, metric metrics.Metrics, logger loggerPkg.Logger, ctx context.Context) *BookController {
	return &BookController{Service: srv, Metrics: metric, Logger: logger, Ctx: ctx}
}

func (bc *BookController) ListBooks(c echo.Context) error {
	var request models.RequestBook
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	books, err := bc.Service.ListBooks(bc.Ctx, request)
	if err != nil {
		bc.Logger.Info("error", slog.Group("err", err))
		bc.Metrics.IncCounter("controller.book.ListBooks.error", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}
	return c.JSON(http.StatusOK, models.WebResponse[[]*models.Book]{Data: books, Status: "ok"})
}

func (bc *BookController) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	book, err := bc.Service.GetBookByID(bc.Ctx, id)
	if err != nil {
		bc.Logger.Info("error", slog.String("id", id), slog.String("err", err.Error()))
		bc.Metrics.IncCounter("controller.book.GetBookByID.error", err.Error())
		return echo.NewHTTPError(http.StatusNotFound, errors.BookNotFound)
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.Book]{Data: book, Status: "ok"})
}

func (bc *BookController) CreateBook(c echo.Context) error {
	var request models.CreateBookRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	createdBook, err := bc.Service.CreateBook(bc.Ctx, request)
	if err != nil {
		bc.Logger.Info("error", slog.String("err", err.Error()))
		bc.Metrics.IncCounter("controller.book.CreateBook.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.BookNotCreated)
	}
	return c.JSON(http.StatusCreated, models.WebResponse[*models.Book]{Data: createdBook, Status: "created"})
}

func (bc *BookController) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	var request models.UpdateBookRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	updatedBook, err := bc.Service.UpdateBook(bc.Ctx, id, request)
	if err != nil {
		bc.Logger.Info("error", slog.String("err", err.Error()))
		bc.Metrics.IncCounter("controller.book.UpdateBook.error", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Book]{Data: updatedBook, Status: "updated"})
}

func (bc *BookController) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	book, err := bc.Service.DeleteBook(bc.Ctx, id)
	if err != nil {
		bc.Logger.Info("error", slog.String("err", err.Error()))
		bc.Metrics.IncCounter("controller.book.DeleteBook.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.BookNotDeleted)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Book]{Data: book, Status: "deleted"})
}
