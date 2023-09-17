package repo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/model"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/entity"

	"github.com/jmoiron/sqlx"
)

type IGetProduct interface {
	GetByID(ctx context.Context, id uint64) (res entity.Product, err error)
}

type productGetter struct {
	db *sqlx.DB
}

func NewProductGetter(db *sqlx.DB) *productGetter {
	return &productGetter{
		db: db,
	}
}

func (repo *productGetter) GetByID(ctx context.Context, id uint64) (res entity.Product, err error) {
	productModel := model.Product{}
	query := `
	SELECT id, name, rating, description, created_at, updated_at, category_id
	FROM products
	WHERE id = $1 AND deleted_at IS NULL
	`

	row := repo.db.QueryRowxContext(ctx, query, id)
	err = row.Scan(
		&productModel.ID,
		&productModel.Name,
		&productModel.Rating,
		&productModel.Description,
		&productModel.CreatedAt,
		&productModel.UpdatedAt,
		&productModel.CategoryID,
	)
	if err != nil {
		return res, err
	}

	res = entity.Product{
		ID:          productModel.ID,
		Name:        productModel.Name,
		Rating:      int(productModel.Rating.Int64),
		Description: productModel.Description.String,
		CreatedAt:   productModel.CreatedAt,
		UpdatedAt:   productModel.UpdatedAt.Time,
		CategoryID:  productModel.CategoryID,
	}

	return res, nil
}
