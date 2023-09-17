package category

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/valuetype"
)

type ICreateCategory interface {
	CreateCategory(ctx context.Context, in valuetype.InsertCategoryIn) (err error)
}

type createCategory struct {
	repo repo.IInsertCategory
}

func NewCreateCategory(repo repo.IInsertCategory) *createCategory {
	return &createCategory{
		repo: repo,
	}
}

func (uc *createCategory) CreateCategory(ctx context.Context, in valuetype.InsertCategoryIn) (err error) {
	return uc.repo.Insert(ctx, in)
}
