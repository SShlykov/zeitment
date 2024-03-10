package minio

import (
	"github.com/SShlykov/zeitment/bookback/pkg/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient() (*minio.Client, error) {
	options := config.GetMinioCreds()

	return minio.New(options.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(options.AccessKeyID, options.SecretKey, ""),
		Secure: options.Secure,
	})
}
