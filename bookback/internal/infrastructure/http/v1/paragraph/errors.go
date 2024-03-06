package paragraph

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
	ErrorParagraphNotFound = echo.NewHTTPError(
		http.StatusNotFound,
		"Параграф не найдена",
	)
	ErrorParagraphNotCreated = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка создания параграфа. Параграф с такими параметрами уже существует.",
	)
	ErrorDeleteParagraph = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка удаления параграфа",
	)
	ErrorUnknown = echo.NewHTTPError(
		http.StatusInternalServerError,
		"Неизвестная ошибка",
	)
)
