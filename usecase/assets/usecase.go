package assets

import (
	"context"
	"geoproperty_be/config"
	"geoproperty_be/domain"
	"mime/multipart"

	"github.com/asnur/minio-go"
)

type ServiceAssets struct {
	minioClient *minio.Client
}

func NewUseCase(minioClient *minio.Client) domain.AssetUsecase {
	return &ServiceAssets{minioClient: minioClient}
}

func (s *ServiceAssets) UploadAsset(ctx context.Context, fh *multipart.FileHeader) (err error) {
	var buf multipart.File
	buf, err = fh.Open()
	if err != nil {
		return err
	}

	defer buf.Close()

	contentType := fh.Header["Content-Type"][0]
	fileSize := fh.Size

	_, err = s.minioClient.PutObject(ctx, config.MinioBucketName, fh.Filename, buf, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceAssets) UploadMultipleAsset(ctx context.Context, fh []*multipart.FileHeader) (err error) {
	for _, file := range fh {
		var buf multipart.File
		buf, err = file.Open()
		if err != nil {
			continue
		}

		defer buf.Close()

		contentType := file.Header["Content-Type"][0]
		fileSize := file.Size

		_, err = s.minioClient.PutObject(ctx, config.MinioBucketName, file.Filename, buf, fileSize, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ServiceAssets) DeleteAsset(ctx context.Context, fileName string) error {
	err := s.minioClient.RemoveObject(ctx, config.MinioBucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
