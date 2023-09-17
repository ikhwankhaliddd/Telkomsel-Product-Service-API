package publicrepo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/entity"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/internal/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/valuetype"
	"github.com/jmoiron/sqlx"
)

type IGetCategory interface {
	GetByName(ctx context.Context, in valuetype.GetCategoryIn) (res entity.Category, err error)
}

type categoryGetter struct {
	db *sqlx.DB
}

func NewCategoryGetter(db *sqlx.DB) *categoryGetter {
	return &categoryGetter{db: db}
}

func (repo *categoryGetter) GetByName(ctx context.Context, in valuetype.GetCategoryIn) (res entity.Category, err error) {
	categoryModel := model.Category{
		Name: in.Name,
	}

	query := `SELECT id FROM categories WHERE name = $1`

	row := repo.db.QueryRowxContext(ctx, query, categoryModel.Name)
	err = row.Scan(&categoryModel.ID)
	if err != nil {
		return res, err
	}

	res.ID = categoryModel.ID

	return res, nil
}
