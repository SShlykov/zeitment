package book

import (
	"context"
	"errors"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	service "github.com/SShlykov/zeitment/bookback/internal/services/book"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Controller структура для HTTP-контроллера книг.
type Controller struct {
	Service service.Service
}

// NewController создает новый экземпляр Controller.
func NewController(srv service.Service) *Controller {
	return &Controller{Service: srv}
}

// RegisterRoutes регистрирует маршруты для обработки запросов к книгам.
func (bc *Controller) RegisterRoutes(e *echo.Echo, ctx context.Context) {
	e.GET("/api/v1/books", func(c echo.Context) error { return bc.ListBooks(c, ctx) })
	e.POST("/api/v1/books", func(c echo.Context) error { return bc.CreateBook(c, ctx) })
	e.GET("/api/v1/books/:id", func(c echo.Context) error { return bc.GetBookByID(c, ctx) })
	e.PUT("/api/v1/books/:id", func(c echo.Context) error { return bc.UpdateBook(c, ctx) })
	e.DELETE("/api/v1/books/:id", func(c echo.Context) error { return bc.DeleteBook(c, ctx) })
}

// ListBooks обрабатывает запросы на получение списка книг.
// @router /books [get]
// @summary Получить список книг
// @description Извлекает список всех книг
// @tags Книги
// @produce  application/json
// @success 200 {array} models.Book
// @failure 500 {object} config.HTTPError
func (bc *Controller) ListBooks(c echo.Context, ctx context.Context) error {
	books, err := bc.Service.ListBooks(ctx)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadGateway, config.ErrorForbidden)
	}
	return c.JSON(http.StatusOK, books)
}

// CreateBook обрабатывает создание новой книги.
// @router /books [post]
// @summary Создать книгу
// @description Создает новую книгу
// @tags Книги
// @accept application/json
// @produce application/json
// @param book body models.Book true "Book object"
// @success 201 {object} models.Book
// @failure 400 {object} config.HTTPError
func (bc *Controller) CreateBook(c echo.Context, ctx context.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}
	createdBook, err := bc.Service.CreateBook(ctx, &book)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, config.ErrorNotCreated)
	}
	return c.JSON(http.StatusCreated, createdBook)
}

// GetBookByID обрабатывает запросы на получение книги по ID.
// @router /books/{id} [get]
// @summary Получить книгу по ID
// @description Извлекает книгу по ее ID
// @tags Книги
// @param id path string true "Book ID"
// @produce application/json
// @success 200 {object} models.Book
// @failure 404 {object} config.HTTPError
func (bc *Controller) GetBookByID(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	book, err := bc.Service.GetBookByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}
	return c.JSON(http.StatusOK, book)
}

// UpdateBook обрабатывает обновление книги.
// @router /books/{id} [put]
// @summary Обновить книгу
// @description Обновляет книгу по ее ID
// @tags Книги
// @accept application/json
// @produce application/json
// @param id path string true "Book ID"
// @param book body models.Book true "Book object"
// @success 200 {object} models.Book
// @failure 400 {object} config.HTTPError
func (bc *Controller) UpdateBook(c echo.Context, ctx context.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, config.ErrorBadInput)
	}
	paramID := c.Param("id")
	updatedBook, err := bc.Service.UpdateBook(ctx, paramID, &book)
	if err != nil {
		if errors.Is(err, config.ErrorNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
		}
		return echo.NewHTTPError(http.StatusNotAcceptable, config.ErrorNotUpdated)
	}
	return c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook обрабатывает удаление книги по ID.
// @router /books/{id} [delete]
// @summary Удалить книгу
// @description Удаляет книгу по ее ID
// @tags Книги
// @param id path string true "Book ID"
// @success 204
// @failure 404 {object} config.HTTPError
func (bc *Controller) DeleteBook(c echo.Context, ctx context.Context) error {
	id := c.Param("id")
	book, err := bc.Service.DeleteBook(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ErrorNotFound)
	}
	return c.JSON(http.StatusOK, book)
}
