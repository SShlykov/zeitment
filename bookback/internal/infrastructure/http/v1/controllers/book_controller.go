package controllers

import (
	"context"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/errors"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/SShlykov/zeitment/metrics"
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

	GetTableOfContentsByBookID(ctx context.Context, request models.RequestTOC) (*models.TableOfContents, error)
	TogglePublic(ctx context.Context, request models.ToggleBookRequest) (*models.Book, error)
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

func (bc *BookController) TogglePublic(c echo.Context) error {
	var request models.ToggleBookRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	book, err := bc.Service.TogglePublic(bc.Ctx, request)
	if err != nil {
		bc.Logger.Info("error", loggerPkg.Err(err))
		bc.Metrics.IncCounter("controller.book.TogglePublic.error", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}

	fmt.Println("book", book)
	return c.JSON(http.StatusOK, models.WebResponse[*models.Book]{Data: book, Status: "ok"})
}

func (bc *BookController) GetTableOfContentsByBookID(c echo.Context) error {
	var request models.RequestTOC
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	toc, err := bc.Service.GetTableOfContentsByBookID(bc.Ctx, request)
	if err != nil {
		bc.Logger.Info("error", loggerPkg.Err(err))
		bc.Metrics.IncCounter("controller.book.GetTableOfContentsByBookID.error", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Unknown)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.TableOfContents]{Data: toc, Status: "ok"})
}

func (bc *BookController) ListBooks(c echo.Context) error {
	var request models.RequestBook
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	books, err := bc.Service.ListBooks(bc.Ctx, request)
	if err != nil {
		bc.Logger.Info("error", loggerPkg.Err(err))
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
		bc.Logger.Info("error", slog.String("id", id), loggerPkg.Err(err))
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
		bc.Logger.Info("error", loggerPkg.Err(err))
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
		bc.Logger.Info("error", loggerPkg.Err(err))
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
		bc.Logger.Info("error", loggerPkg.Err(err))
		bc.Metrics.IncCounter("controller.book.DeleteBook.error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, errors.BookNotDeleted)
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Book]{Data: book, Status: "deleted"})
}
