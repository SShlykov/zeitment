package routes

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/services"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/controllers"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/pgrepo"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/SShlykov/zeitment/metrics"
	"github.com/SShlykov/zeitment/postgres"
	"github.com/labstack/echo/v4"
)

// Book регистрирует контроллер книг в маршрутизаторе.
func Book(e *echo.Echo, database postgres.Client, metrics metrics.Metrics, logger loggerPkg.Logger, ctx context.Context) {
	repo := pgrepo.NewBookRepository(database)
	service := services.NewBookService(repo)
	cntr := controllers.NewBookController(service, metrics, logger, ctx)

	group := e.Group(v1.BooksPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.POST("/table_of_content", cntr.GetTableOfContentsByBookID)
	group.POST(v1.ListSubPath, cntr.ListBooks)
	group.POST("", cntr.CreateBook)
	group.GET(v1.IDVar, cntr.GetBookByID)

	group.PUT(v1.ToggleSubPath, cntr.TogglePublic)
	group.PUT(v1.IDVar, cntr.UpdateBook)
	group.DELETE(v1.IDVar, cntr.DeleteBook)
}
