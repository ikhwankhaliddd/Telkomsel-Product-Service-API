package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/variety/valuetype"
	"github.com/jmoiron/sqlx"
)

type IUpdateVariety interface {
	Update(ctx context.Context, id uint64, in valuetype.UpdateVarietyIn) (err error)
}

type varietyUpdatter struct {
	db *sqlx.DB
}

func NewVarietyUpdater(db *sqlx.DB) *varietyUpdatter {
	return &varietyUpdatter{db: db}
}

func (repo *varietyUpdatter) Update(ctx context.Context, id uint64, in valuetype.UpdateVarietyIn) (err error) {
	varietyModel := model.Variety{
		ID:    id,
		Name:  in.Name,
		Price: in.Price,
		Stock: in.Stock,
	}

	var queryBuilder strings.Builder
	queryBuilder.WriteString("UPDATE variety SET ")

	var queryParams []interface{}
	var paramCount int

	if in.Name != "" {
		paramCount++
		queryBuilder.WriteString(fmt.Sprintf(" name = $%d, ", paramCount))
		queryParams = append(queryParams, varietyModel.Name)
	}

	if in.Price != 0 {
		paramCount++
		queryBuilder.WriteString(fmt.Sprintf(" price = $%d, ", paramCount))
		queryParams = append(queryParams, varietyModel.Price)
	}

	if in.Stock != 0 {
		paramCount++
		queryBuilder.WriteString(fmt.Sprintf(" stock = $%d, ", paramCount))
		queryParams = append(queryParams, varietyModel.Stock)
	}

	paramCount++
	queryBuilder.WriteString(fmt.Sprintf("updated_at = NOW() WHERE id = $%d", paramCount))
	queryParams = append(queryParams, varietyModel.ID)

	query := queryBuilder.String()

	result, err := repo.db.ExecContext(ctx, query, queryParams...)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("failed to update variety. No rows affected")
	}

	return nil
}
