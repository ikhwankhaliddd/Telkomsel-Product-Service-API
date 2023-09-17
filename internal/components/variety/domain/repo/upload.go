package repo

import (
	"context"
	"errors"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/variety/valuetype"
	"github.com/ikhwankhaliddd/product-service/internal/helper/converter"
	"github.com/jmoiron/sqlx"
)

type IUploadImage interface {
	UploadImage(ctx context.Context, in valuetype.UploadImageIn) (err error)
}

type imageUploader struct {
	db *sqlx.DB
}

func NewImageUploader(db *sqlx.DB) *imageUploader {
	return &imageUploader{
		db: db,
	}
}

func (repo *imageUploader) UploadImage(ctx context.Context, in valuetype.UploadImageIn) (err error) {
	convertedFilename := converter.StringToNullString(in.File.Filename)
	varietyModel := model.Variety{
		ID:    uint64(in.ID),
		Image: convertedFilename,
	}

	query := `UPDATE variety SET image = $1 WHERE id = $2`

	result, err := repo.db.ExecContext(ctx, query, varietyModel.Image, varietyModel.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("failed to update variety. No rows affected")
	}

	return nil
}
