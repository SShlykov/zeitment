package book

import (
	"context"
	service "github.com/SShlykov/zeitment/bookback/internal/domain/services/book"
	"github.com/SShlykov/zeitment/bookback/internal/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/models"
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

// ListBooks обрабатывает запросы на получение списка книг.
// @router /books [post]
// @summary Получить список книг
// @description Извлекает список всех книг
// @tags Книги
// @produce  application/json
// @success 200 {array} entity.Book
// @failure 500 {object} string
func (bc *Controller) ListBooks(c echo.Context) error {
	var request models.RequestBook
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	books, err := bc.Service.ListBooks(bc.Ctx, request.Options.Limit, request.Options.Offset)
	if err != nil {
		return ErrorUnknown
	}
	return c.JSON(http.StatusOK, models.WebResponse[[]*models.Book]{Data: books, Status: "ok"})
}

// CreateBook обрабатывает создание новой книги.
// @router /books [post]
// @summary Создать книгу
// @description Создает новую книгу
// @tags Книги
// @accept application/json
// @produce application/json
// @param book body entity.Book true "Book object"
// @success 201 {object} entity.Book
// @failure 400 {object} string
// @failure 500 {object} string
func (bc *Controller) CreateBook(c echo.Context) error {
	var request models.CreateBookRequest
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	createdBook, err := bc.Service.CreateBook(bc.Ctx, request)
	if err != nil {
		return ErrorBookNotCreated
	}
	return c.JSON(http.StatusCreated, models.WebResponse[*models.Book]{Data: createdBook, Status: "created"})
}

// GetBookByID обрабатывает запросы на получение книги по ID.
// @router /books/{id} [get]
// @summary Получить книгу по ID
// @description Извлекает книгу по ее ID
// @tags Книги
// @param id path string true "Book ID"
// @produce application/json
// @success 200 {object} entity.Book
// @failure 400 {object} string
// @failure 404 {object} string
// @failure 500 {object} string
func (bc *Controller) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	book, err := bc.Service.GetBookByID(bc.Ctx, id)
	if err != nil {
		return ErrorBookNotFound
	}

	return c.JSON(http.StatusOK, models.WebResponse[*models.Book]{Data: book, Status: "ok"})
}

// UpdateBook обрабатывает обновление книги.
// @router /books/{id} [put]
// @summary Обновить книгу
// @description Обновляет книгу по ее ID
// @tags Книги
// @accept application/json
// @produce application/json
// @param id path string true "Book ID"
// @param book body entity.Book true "Book object"
// @success 200 {object} responseSingleModel
// @failure 400 {object} string
// @failure 404 {object} string
// @failure 500 {object} string
func (bc *Controller) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	var request models.UpdateBookRequest
	if err := c.Bind(&request); err != nil {
		return ErrorValidationFailed
	}

	updatedBook, err := bc.Service.UpdateBook(bc.Ctx, id, request)
	if err != nil {
		return ErrorUnknown
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Book]{Data: updatedBook, Status: "updated"})
}

// DeleteBook обрабатывает удаление книги по ID.
// @router /books/{id} [delete]
// @summary Удалить книгу
// @description Удаляет книгу по ее ID
// @tags Книги
// @param id path string true "Book ID"
// @success 204 {object} entity.Book
// @failure 400 {object} string
func (bc *Controller) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrorValidationFailed
	}

	book, err := bc.Service.DeleteBook(bc.Ctx, id)
	if err != nil {
		return ErrorDeleteBook
	}
	return c.JSON(http.StatusOK, models.WebResponse[*models.Book]{Data: book, Status: "deleted"})
}
