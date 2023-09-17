package repo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety/domain/model"
	"github.com/jmoiron/sqlx"
)

type IDeleteVariety interface {
	DeleteVariety(ctx context.Context, id uint64) (err error)
}

type deleteVariety struct {
	db *sqlx.DB
}

func NewDeleteVariety(db *sqlx.DB) *deleteVariety {
	return &deleteVariety{
		db: db,
	}
}

func (repo *deleteVariety) DeleteVariety(ctx context.Context, id uint64) (err error) {
	varietyModel := model.Variety{
		ID: id,
	}

	query := `UPDATE variety SET deleted_at = NOW() WHERE id = $1`

	_, err = repo.db.QueryxContext(ctx, query, varietyModel.ID)
	if err != nil {
		return err
	}

	return nil
}
