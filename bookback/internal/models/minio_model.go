package models

type RequestMinioObject struct {
	BucketName string `json:"bucket_name"`
	ObjectName string `json:"object_name"`
}

type MinioResp struct {
	ContentType string
	Name        string
	Content     []byte
}
