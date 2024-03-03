package chapter

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Возможные ошибки при работе с главами книги.

var (
	ErrorValidationFailed = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка валидации полей ввода! Проверьте введенные данные и попробуйте снова.",
	)
	ErrorChapterNotFound = echo.NewHTTPError(
		http.StatusNotFound,
		"Глава не найдена",
	)
	ErrorChapterNotCreated = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка создания главы. Глава с такими параметрами уже существует.",
	)
	ErrorDeleteChapter = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка удаления главы",
	)
	ErrorUnknown = echo.NewHTTPError(
		http.StatusInternalServerError,
		"Неизвестная ошибка",
	)
)
