package domain

import (
	"context"
	"mime/multipart"
)

type AssetUsecase interface {
	UploadAsset(ctx context.Context, fh *multipart.FileHeader) error
	UploadMultipleAsset(ctx context.Context, fh []*multipart.FileHeader) error
	DeleteAsset(ctx context.Context, fileName string) error
}
