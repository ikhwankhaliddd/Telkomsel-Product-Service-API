package variety

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/variety/valuetype"
)

type IUpdateVariety interface {
	Update(ctx context.Context, id uint64, in valuetype.UpdateVarietyIn) (err error)
}

type varietyUpdatter struct {
	repo repo.IUpdateVariety
}

func NewVarietyUpdatter(repo repo.IUpdateVariety) *varietyUpdatter {
	return &varietyUpdatter{repo: repo}
}

func (uc *varietyUpdatter) Update(ctx context.Context, id uint64, in valuetype.UpdateVarietyIn) (err error) {
	return uc.repo.Update(ctx, id, in)
}
