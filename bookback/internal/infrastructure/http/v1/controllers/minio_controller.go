package controllers

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1/errors"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type MinioUseCase interface {
	CreateMinioObject(ctx context.Context, request models.RequestMinioObject) (string, error)
	GetMinioObject(ctx context.Context, request models.RequestMinioObject) (*models.MinioResp, error)
}

type MinioController interface {
	CreateMinioObject(c echo.Context) error
	GetMinioObject(c echo.Context) error
	DownloadMinioObject(c echo.Context) error
}

type minioController struct {
	Service MinioUseCase
	Metrics metrics.Metrics
	Logger  loggerPkg.Logger
	Ctx     context.Context
}

func NewMinioController(srv MinioUseCase, metric metrics.Metrics, logger loggerPkg.Logger, ctx context.Context) MinioController {
	return &minioController{Service: srv, Metrics: metric, Logger: logger, Ctx: ctx}
}

func (mc *minioController) CreateMinioObject(c echo.Context) error {
	var request models.RequestMinioObject
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ValidationFailed)
	}

	object, err := mc.Service.CreateMinioObject(mc.Ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.MinioObjectNotCreated)
	}
	return c.JSON(http.StatusCreated, object)
}

func (mc *minioController) GetMinioObject(c echo.Context) error {
	var request models.RequestMinioObject
	request.ObjectName = c.Param("object")
	request.BucketName = c.Param("bucket")

	object, err := mc.Service.GetMinioObject(mc.Ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.MinioObjectNotFound)
	}

	return c.Blob(http.StatusOK, object.ContentType, object.Content)
}

func (mc *minioController) DownloadMinioObject(c echo.Context) error {
	var request models.RequestMinioObject
	request.ObjectName = c.Param("object")
	request.BucketName = c.Param("bucket")

	object, err := mc.Service.GetMinioObject(mc.Ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.MinioObjectNotFound)
	}
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+"test"+object.Name)

	return c.Blob(http.StatusOK, object.ContentType, object.Content)
}
