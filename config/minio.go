package config

import (
	"context"
	"os"

	"github.com/asnur/minio-go"
	"github.com/asnur/minio-go/pkg/credentials"
)

var MinioBucketName = ""

// Connect to Bucket Chat
func ConnectBucket(ctx context.Context) (*minio.Client, error) {
	minioClient, err := minio.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIOUSER"), os.Getenv("MINIOPASSWORD"), ""),
		Secure: false,
	})

	MinioBucketName = os.Getenv("BUCKETNAME")

	if err != nil {
		return nil, err
	}

	_, err_exist := minioClient.BucketExists(ctx, os.Getenv("BUCKETNAME"))

	if err_exist != nil {
		return nil, err_exist
	}

	return minioClient, nil
}
