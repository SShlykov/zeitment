package routes

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/application/usecase"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/middleware"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/controllers"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

// Minio регистрирует контроллер minio в маршрутизаторе.
func Minio(e *echo.Echo, client *minio.Client, metrics metrics.Metrics, logger loggerPkg.Logger, ctx context.Context) {
	srv := usecase.NewMinioUseCase(client)
	cntr := controllers.NewMinioController(srv, metrics, logger, ctx)

	group := e.Group(v1.MinioPath)
	group.Use(middleware.MetricsLogger(metrics))

	group.GET("/get-object/:bucket/:object", cntr.GetMinioObject)
	group.GET("/download-object/:bucket/:object", cntr.DownloadMinioObject)
}
