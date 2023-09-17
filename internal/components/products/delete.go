package products

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
)

type IDeleteProduct interface {
	DeleteProduct(ctx context.Context, id uint64) (err error)
}

type deleteProduct struct {
	repo repo.IDeleteProduct
}

func NewDeleteProduct(repo repo.IDeleteProduct) *deleteProduct {
	return &deleteProduct{
		repo: repo,
	}
}

func (uc *deleteProduct) DeleteProduct(ctx context.Context, id uint64) (err error) {
	return uc.repo.DeleteProduct(ctx, id)
}
