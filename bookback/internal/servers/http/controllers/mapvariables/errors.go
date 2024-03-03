package mapvariables

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrorValidationFailed = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка валидации полей ввода! Проверьте введенные данные и попробуйте снова.",
	)
	ErrorMapVariableNotFound = echo.NewHTTPError(
		http.StatusNotFound,
		"Переменная не найдена",
	)
	ErrorMapVariableNotCreated = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка создания переменной. Переменная с такими параметрами уже существует.",
	)
	ErrorDeleteMapVariable = echo.NewHTTPError(
		http.StatusBadRequest,
		"Ошибка удаления переменной",
	)
	ErrorUnknown = echo.NewHTTPError(
		http.StatusInternalServerError,
		"Неизвестная ошибка",
	)
)
