package repo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/entity"
	"github.com/jmoiron/sqlx"
)

type IGetCategoryList interface {
	GetCategoryList(ctx context.Context) (res []entity.Category, err error)
}

type categoryListGetter struct {
	db *sqlx.DB
}

func NewCategoryListGetter(db *sqlx.DB) *categoryListGetter {
	return &categoryListGetter{
		db: db,
	}
}

func (repo *categoryListGetter) GetCategoryList(ctx context.Context) (res []entity.Category, err error) {
	query := `SELECT id, name FROM categories WHERE deleted_at IS NULL`

	rows, err := repo.db.QueryxContext(ctx, query)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		categoryModel := model.Category{}
		err := rows.Scan(
			&categoryModel.ID,
			&categoryModel.Name,
		)
		if err != nil {
			return res, err
		}

		res = append(res, entity.Category{
			ID:   categoryModel.ID,
			Name: categoryModel.Name,
		})
	}
	return res, nil
}
