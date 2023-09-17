package variety

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/variety/valuetype"
	"github.com/ikhwankhaliddd/product-service/internal/helper/uploader"
	"github.com/spf13/viper"
)

type IUploadImage interface {
	UploadImage(ctx context.Context, in valuetype.UploadImageIn) (err error)
}

type imageUploader struct {
	repo       repo.IUploadImage
	s3Uploader uploader.IUploadHelper
}

func NewImageUploader(repo repo.IUploadImage, s3Uploader uploader.IUploadHelper) *imageUploader {
	return &imageUploader{
		repo:       repo,
		s3Uploader: s3Uploader,
	}
}

func (uc *imageUploader) UploadImage(ctx context.Context, in valuetype.UploadImageIn) (err error) {
	fileUpload, err := in.File.Open()
	if err != nil {
		return err
	}
	defer fileUpload.Close()

	bucketName := viper.GetString("bucketName")

	err = uc.s3Uploader.UploadFile(
		ctx,
		bucketName,
		in.File.Filename,
		in.File.Header.Get("Content-Type"),
		fileUpload,
	)
	if err != nil {
		return err
	}

	err = uc.repo.UploadImage(ctx, in)
	if err != nil {
		return err
	}
	return nil

}
