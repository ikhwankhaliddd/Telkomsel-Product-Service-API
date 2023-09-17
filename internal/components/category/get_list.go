package category

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/entity"
)

type IGetCategoryList interface {
	GetCategoryList(ctx context.Context) (res []entity.Category, err error)
}

type categoryListGetter struct {
	repo repo.IGetCategoryList
}

func NewCategoryListGetter(repo repo.IGetCategoryList) *categoryListGetter {
	return &categoryListGetter{
		repo: repo,
	}
}
func (uc *categoryListGetter) GetCategoryList(ctx context.Context) (res []entity.Category, err error) {
	return uc.repo.GetCategoryList(ctx)
}
