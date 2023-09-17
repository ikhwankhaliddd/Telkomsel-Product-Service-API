package variety

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety/domain/repo"
)

type IDeleteVariety interface {
	DeleteVariety(ctx context.Context, id uint64) (err error)
}

type deleteVariety struct {
	repo repo.IDeleteVariety
}

func NewDeleteVariety(repo repo.IDeleteVariety) *deleteVariety {
	return &deleteVariety{
		repo: repo,
	}
}

func (uc *deleteVariety) DeleteVariety(ctx context.Context, id uint64) (err error) {
	return uc.repo.DeleteVariety(ctx, id)
}
