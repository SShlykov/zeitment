package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
)

func NewMinioClient() (*minio.Client, error) {
	host, accessKeyID, secretKey := GetMinioConfig()

	return minio.New(host, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretKey, ""),
		Secure: true,
	})
}

func GetObject(ctx context.Context, minioClient *minio.Client) {
	bucketName, objectName := "test", "micheile-henderson-lZ_4nPFKcV8-unsplash.jpg"

	reader, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	localFile, err := os.Create("testfile.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer localFile.Close()

	stat, err := reader.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err = io.CopyN(localFile, reader, stat.Size); err != nil {
		log.Fatalln(err)
	}
}
