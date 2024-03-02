package router

import (
	"context"
	_ "github.com/SShlykov/zeitment/bookback/docs"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/book"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/bookevents"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/chapter"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/health"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/mapvariables"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/page"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/paragraph"
	bookrepo "github.com/SShlykov/zeitment/bookback/internal/services/book"
	bookeventsrepo "github.com/SShlykov/zeitment/bookback/internal/services/bookevents"
	chapterrepo "github.com/SShlykov/zeitment/bookback/internal/services/chapter"
	mapvariablesrepo "github.com/SShlykov/zeitment/bookback/internal/services/mapvariables"
	pagerepo "github.com/SShlykov/zeitment/bookback/internal/services/page"
	paragraphrepo "github.com/SShlykov/zeitment/bookback/internal/services/paragraph"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func SetCORSConfig(e *echo.Echo, corsEnabled bool) {
	if corsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials:                         true,
			UnsafeWildcardOriginWithAllowCredentials: true,
			AllowOrigins:                             []string{"*"},
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}))
	}
}

func SetHealthController(e *echo.Echo, ctx context.Context) {
	health.NewController().RegisterRoutes(e, ctx)
}

func SetBookController(e *echo.Echo, database db.Client, ctx context.Context) {
	book.NewController(bookrepo.NewService(bookrepo.NewRepository(database))).RegisterRoutes(e, ctx)
}

func SetPageController(e *echo.Echo, database db.Client, ctx context.Context) {
	page.NewController(pagerepo.NewService(pagerepo.NewRepository(database))).RegisterRoutes(e, ctx)
}

func SetChapterController(e *echo.Echo, database db.Client, ctx context.Context) {
	chapterRepo := chapterrepo.NewRepository(database)
	bookRepo := bookrepo.NewRepository(database)
	chapter.NewController(chapterrepo.NewService(chapterRepo, bookRepo)).RegisterRoutes(e, ctx)
}

func SetParagraphController(e *echo.Echo, database db.Client, ctx context.Context) {
	paragraph.NewController(paragraphrepo.NewService(paragraphrepo.NewRepository(database))).RegisterRoutes(e, ctx)
}

func SetMapVariablesController(e *echo.Echo, database db.Client, ctx context.Context) {
	mapvariables.NewController(mapvariablesrepo.NewService(mapvariablesrepo.NewRepository(database))).RegisterRoutes(e, ctx)
}

func SetBookEventController(e *echo.Echo, database db.Client, ctx context.Context) {
	bookevents.NewController(bookeventsrepo.NewService(bookeventsrepo.NewRepository(database))).RegisterRoutes(e, ctx)
}

func SetSwagger(e *echo.Echo, swaggerEnabled bool) {
	if swaggerEnabled {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
}