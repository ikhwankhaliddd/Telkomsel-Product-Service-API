package products

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/entity"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"

	categoryRepo "github.com/ikhwankhaliddd/product-service/internal/components/category/public_repo"
	varietyRepo "github.com/ikhwankhaliddd/product-service/internal/components/variety/public_repo"
)

type IGetListProduct interface {
	GetListProduct(ctx context.Context, in valuetype.GetProductListIn) (res []entity.Product, count int, err error)
}

type listProductGetter struct {
	repo         repo.IGetListProduct
	varietyRepo  varietyRepo.IGetVariety
	categoryRepo categoryRepo.IGetCategoryByID
}

func NewProductListGetter(
	repo repo.IGetListProduct,
	varietyRepo varietyRepo.IGetVariety,
	categoryRepo categoryRepo.IGetCategoryByID,
) *listProductGetter {
	return &listProductGetter{
		repo:         repo,
		varietyRepo:  varietyRepo,
		categoryRepo: categoryRepo,
	}
}

func (uc *listProductGetter) GetListProduct(ctx context.Context, in valuetype.GetProductListIn) (res []entity.Product, count int, err error) {
	productList, count, err := uc.repo.GetListProduct(ctx, in)

	for _, product := range productList {

		res = append(res, entity.Product{
			ID:          product.ID,
			CategoryID:  product.CategoryID,
			Name:        product.Name,
			Description: product.Description,
			Rating:      product.Rating,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}

	return res, count, err
}
