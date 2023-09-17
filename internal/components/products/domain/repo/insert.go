package repo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/entity"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"

	"github.com/ikhwankhaliddd/product-service/internal/helper/converter"
	"github.com/jmoiron/sqlx"
)

type IInsertProduct interface {
	Insert(ctx context.Context, in valuetype.CreateProductIn) (res entity.Product, err error)
}

type insertProduct struct {
	db *sqlx.DB
}

func NewInsertProduct(db *sqlx.DB) *insertProduct {
	return &insertProduct{
		db: db,
	}
}

func (repo *insertProduct) Insert(ctx context.Context, in valuetype.CreateProductIn) (res entity.Product, err error) {
	convertedDescription := converter.StringToNullString(in.Description)
	productModel := model.Product{
		Name:        in.Name,
		CategoryID:  in.CategoryID,
		Description: convertedDescription,
	}

	query := `INSERT INTO "products" (
		"name",
		"category_id",
		"description",
		"created_at"
	) VALUES (
		$1,
		$2,
		$3,
		NOW()
	) RETURNING "products"."id"`

	row := repo.db.QueryRowx(query,
		productModel.Name,
		productModel.CategoryID,
		productModel.Description)

	var id uint64
	if err := row.Scan(&id); err != nil {
		return res, err
	}

	res.ID = id

	return res, nil
}
