package repo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/valuetype"
	"github.com/jmoiron/sqlx"
)

type IInsertCategory interface {
	Insert(ctx context.Context, in valuetype.InsertCategoryIn) (err error)
}

type insertCategory struct {
	db *sqlx.DB
}

func NewInsertCategory(db *sqlx.DB) *insertCategory {
	return &insertCategory{
		db: db,
	}
}

func (repo *insertCategory) Insert(ctx context.Context, in valuetype.InsertCategoryIn) (err error) {
	categoryModel := model.Category{
		Name: in.Name,
	}

	query := `INSERT INTO categories (name, created_at) VALUES($1, NOW())`

	_, err = repo.db.ExecContext(ctx, query, categoryModel.Name)
	if err != nil {
		return err
	}

	return nil
}
