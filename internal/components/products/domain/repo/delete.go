package repo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/model"
	"github.com/jmoiron/sqlx"
)

type IDeleteProduct interface {
	DeleteProduct(ctx context.Context, id uint64) (err error)
}

type deleteProduct struct {
	db *sqlx.DB
}

func NewDeleteProduct(db *sqlx.DB) *deleteProduct {
	return &deleteProduct{
		db: db,
	}
}

func (repo *deleteProduct) DeleteProduct(ctx context.Context, id uint64) (err error) {
	productModel := model.Product{
		ID: id,
	}

	query := `UPDATE products SET deleted_at = NOW() WHERE id = $1`

	_, err = repo.db.QueryxContext(ctx, query, productModel.ID)
	if err != nil {
		return err
	}

	return nil
}
