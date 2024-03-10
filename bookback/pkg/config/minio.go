package config

import (
	"os"
)

type MinioOptions struct {
	Address     string
	AccessKeyID string
	SecretKey   string
	Secure      bool
}

func GetMinioCreds() MinioOptions {
	host := os.Getenv("MINIO_HOST")
	port := os.Getenv("MINIO_PORT")
	addr := host + ":" + port

	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	secureString := os.Getenv("MINIO_SECURE")
	minioSecure := secureString == "true"

	return MinioOptions{Address: addr, AccessKeyID: accessKeyID, SecretKey: secretKey, Secure: minioSecure}
}
