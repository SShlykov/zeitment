package router

import (
	"context"
	_ "github.com/SShlykov/zeitment/bookback/docs"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/book"
	"github.com/SShlykov/zeitment/bookback/internal/servers/http/controllers/health"
	bookrepo "github.com/SShlykov/zeitment/bookback/internal/services/book"
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

func SetSwagger(e *echo.Echo, swaggerEnabled bool) {
	if swaggerEnabled {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
}
