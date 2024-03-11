package usecase

import (
	"bytes"
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/minio/minio-go/v7"
	"io"
)

//go:generate mockgen -source=minio_usecase.go -destination=../../../tests/mocks/usecase/minio_usecase_mock.go
type MinioUseCase interface {
	CreateMinioObject(ctx context.Context, request models.RequestMinioObject) (string, error)
	GetMinioObject(ctx context.Context, request models.RequestMinioObject) (*models.MinioResp, error)
}

type MinioClient interface {
	GetObject(ctx context.Context, bucketName, objectName string, opts minio.GetObjectOptions) (*minio.Object, error)
}

type minioUseCase struct {
	minioClient MinioClient
}

func NewMinioUseCase(minioClient MinioClient) MinioUseCase {
	return &minioUseCase{minioClient: minioClient}
}

func (ms *minioUseCase) CreateMinioObject(_ context.Context, _ models.RequestMinioObject) (string, error) {
	return "", nil
}

func (ms *minioUseCase) GetMinioObject(ctx context.Context, request models.RequestMinioObject) (*models.MinioResp, error) {
	reader, err := ms.minioClient.GetObject(ctx, request.BucketName, request.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer func(reader *minio.Object) {
		_ = reader.Close()
	}(reader)

	contentBuffer := new(bytes.Buffer)
	if _, err = io.Copy(contentBuffer, reader); err != nil {
		return nil, err
	}
	stat, err := reader.Stat()
	if err != nil {
		return nil, err
	}

	return &models.MinioResp{ContentType: stat.ContentType, Content: contentBuffer.Bytes(), Name: stat.Key}, nil
}
