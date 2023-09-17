package products

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"
)

type IPostRating interface {
	PostRating(ctx context.Context, in valuetype.PostRatingIn) (err error)
}

type postRating struct {
	repo repo.IInsertRating
}

func NewPostRating(repo repo.IInsertRating) *postRating {
	return &postRating{
		repo: repo,
	}
}

func (uc *postRating) PostRating(ctx context.Context, in valuetype.PostRatingIn) (err error) {
	return uc.repo.InsertRating(ctx, in)
}
