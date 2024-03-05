package bookevents

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Возможные ошибки при работе с событиями книг.

var (
	ErrorValidationFailed = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка валидации полей ввода! Проверьте введенные данные и попробуйте снова.",
	)
	ErrorBookEventNotFound = echo.NewHTTPError(
		http.StatusNotFound,
		"Событие книги не найдено",
	)
	ErrorBookEventNotCreated = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка создания события книги. Событие с такими параметрами уже существует.",
	)
	ErrorDeleteBookEvent = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка удаления события книги",
	)
	ErrorUnknown = echo.NewHTTPError(
		http.StatusInternalServerError,
		"Неизвестная ошибка",
	)
)
