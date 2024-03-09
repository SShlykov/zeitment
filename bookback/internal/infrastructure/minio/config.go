package minio

import "os"

func GetMinioConfig() (string, string, string) {
	host := os.Getenv("MINIO_HOST")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")

	return host, accessKeyID, secretKey
}
