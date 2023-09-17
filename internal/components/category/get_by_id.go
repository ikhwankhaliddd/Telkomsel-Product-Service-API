package category

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/entity"
	publicrepo "github.com/ikhwankhaliddd/product-service/internal/components/category/public_repo"
)

type IGetCategoryByID interface {
	GetByID(ctx context.Context, id uint64) (res entity.Category, err error)
}

type categoryByIDGetter struct {
	repo publicrepo.IGetCategoryByID
}

func NewCategoryGetter(repo publicrepo.IGetCategoryByID) *categoryByIDGetter {
	return &categoryByIDGetter{
		repo: repo,
	}
}

func (uc *categoryByIDGetter) GetByID(ctx context.Context, id uint64) (res entity.Category, err error) {
	return uc.repo.GetByID(ctx, id)
}
