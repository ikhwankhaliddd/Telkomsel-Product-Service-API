package repo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"
	"github.com/ikhwankhaliddd/product-service/internal/helper/converter"
	"github.com/jmoiron/sqlx"
)

type IInsertRating interface {
	InsertRating(ctx context.Context, in valuetype.PostRatingIn) (err error)
}

type insertRating struct {
	db *sqlx.DB
}

func NewInsertRating(db *sqlx.DB) *insertRating {
	return &insertRating{
		db: db,
	}
}

func (repo *insertRating) InsertRating(ctx context.Context, in valuetype.PostRatingIn) (err error) {
	convertedRating := converter.Int64ToNullInt64(int64(in.Rating))
	productModel := model.Product{
		ID:     in.ID,
		Rating: convertedRating,
	}

	query := `UPDATE products SET rating = $1 WHERE id = $2`
	_, err = repo.db.QueryxContext(ctx, query, productModel.Rating, productModel.ID)
	if err != nil {
		return err
	}

	return nil
}
