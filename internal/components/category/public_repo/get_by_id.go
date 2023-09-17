package publicrepo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/entity"
	"github.com/jmoiron/sqlx"
)

type IGetCategoryByID interface {
	GetByID(ctx context.Context, id uint64) (res entity.Category, err error)
}

type categoryByIDGetter struct {
	db *sqlx.DB
}

func NewCategoryByIDGetter(db *sqlx.DB) *categoryByIDGetter {
	return &categoryByIDGetter{
		db: db,
	}
}

func (repo *categoryByIDGetter) GetByID(ctx context.Context, id uint64) (res entity.Category, err error) {
	categoryModel := model.Category{}

	query := `SELECT name FROM categories WHERE id = $1`

	row := repo.db.QueryRowxContext(ctx, query, id)

	err = row.Scan(
		&categoryModel.Name,
	)
	if err != nil {
		return res, err
	}

	res.Name = categoryModel.Name
	return res, nil
}
