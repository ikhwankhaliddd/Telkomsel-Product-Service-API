package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"
	"github.com/ikhwankhaliddd/product-service/internal/helper/converter"
	"github.com/jmoiron/sqlx"
)

type IUpdateProduct interface {
	Update(ctx context.Context, id uint64, in valuetype.UpdateProductIn) (err error)
}

type productUpdatter struct {
	db *sqlx.DB
}

func NewProductUpdatter(db *sqlx.DB) *productUpdatter {
	return &productUpdatter{
		db: db,
	}
}

func (repo *productUpdatter) Update(ctx context.Context, id uint64, in valuetype.UpdateProductIn) (err error) {
	convertedDescription := converter.StringToNullString(in.Description)
	productModel := model.Product{
		ID:          id,
		Name:        in.Name,
		Description: convertedDescription,
	}

	var queryBuilder strings.Builder
	queryBuilder.WriteString("UPDATE products SET ")

	var queryParams []interface{}
	var paramCount int

	if in.Name != "" {
		paramCount++
		queryBuilder.WriteString(fmt.Sprintf(" name = $%d, ", paramCount))
		queryParams = append(queryParams, productModel.Name)
	}

	if in.Description != "" {
		paramCount++
		queryBuilder.WriteString(fmt.Sprintf(" description = $%d, ", paramCount))
		queryParams = append(queryParams, productModel.Description)
	}

	paramCount++
	queryBuilder.WriteString(fmt.Sprintf("updated_at = NOW() WHERE id = $%d", paramCount))
	queryParams = append(queryParams, productModel.ID)

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
		return errors.New("failed to update product. No rows affected")
	}

	return nil
}
