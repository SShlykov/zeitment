package page

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Возможные ошибки, которые могут возникнуть при работе с контроллером страниц.

var (
	ErrorValidationFailed = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка валидации полей ввода! Проверьте введенные данные и попробуйте снова.",
	)
	ErrorPageNotFound = echo.NewHTTPError(
		http.StatusNotFound,
		"Страница не найдена",
	)
	ErrorPageNotCreated = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка создания страницы. Страница с такими параметрами уже существует.",
	)
	ErrorDeletePage = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка удаления страницы",
	)
	ErrorUnknown = echo.NewHTTPError(
		http.StatusInternalServerError,
		"Неизвестная ошибка",
	)
)
