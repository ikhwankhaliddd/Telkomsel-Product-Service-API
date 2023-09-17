package products

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"
)

type IUpdateProduct interface {
	Update(ctx context.Context, id uint64, in valuetype.UpdateProductIn) (err error)
}

type productUpdatter struct {
	repo repo.IUpdateProduct
}

func NewProductUpdatter(repo repo.IUpdateProduct) *productUpdatter {
	return &productUpdatter{
		repo: repo,
	}
}

func (uc *productUpdatter) Update(ctx context.Context, id uint64, in valuetype.UpdateProductIn) (err error) {
	return uc.repo.Update(ctx, id, in)
}
