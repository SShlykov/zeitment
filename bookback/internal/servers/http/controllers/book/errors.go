package book

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Возможные ошибки при работе с книгами.

var (
	ErrorValidationFailed = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка валидации полей ввода! Проверьте введенные данные и попробуйте снова.",
	)
	ErrorBookNotFound = echo.NewHTTPError(
		http.StatusNotFound,
		"Книга не найдена",
	)
	ErrorBookNotCreated = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка создания книги. Книга с такими параметрами уже существует.",
	)
	ErrorDeleteBook = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка удаления книги",
	)
	ErrorUnknown = echo.NewHTTPError(
		http.StatusInternalServerError,
		"Неизвестная ошибка",
	)
)
