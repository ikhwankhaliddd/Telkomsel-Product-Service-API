package products

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
	varietyRepo "github.com/ikhwankhaliddd/product-service/internal/components/variety/public_repo"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/entity"
)

type IGetProduct interface {
	GetByID(ctx context.Context, id uint64) (res entity.Product, err error)
}

type productGetter struct {
	repo        repo.IGetProduct
	varietyRepo varietyRepo.IGetVariety
}

func NewProductGetter(
	repo repo.IGetProduct,
	varietyRepo varietyRepo.IGetVariety,
) *productGetter {
	return &productGetter{
		repo:        repo,
		varietyRepo: varietyRepo,
	}
}

func (uc *productGetter) GetByID(ctx context.Context, id uint64) (res entity.Product, err error) {
	variety, err := uc.varietyRepo.GetByProductID(ctx, id)
	if err != nil {
		return res, err
	}

	product, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return res, err
	}

	res = entity.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Varieties:   variety,
		Rating:      product.Rating,
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	return res, nil
}
