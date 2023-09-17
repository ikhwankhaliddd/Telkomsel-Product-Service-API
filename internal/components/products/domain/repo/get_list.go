package repo

import (
	"context"

	sqlb "github.com/elgris/sqrl"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/entity"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"
	"github.com/jmoiron/sqlx"
)

type IGetListProduct interface {
	GetListProduct(ctx context.Context, in valuetype.GetProductListIn) (res []entity.Product, count int, err error)
}

type productListGetter struct {
	db *sqlx.DB
}

func NewProductListGetter(db *sqlx.DB) *productListGetter {
	return &productListGetter{
		db: db,
	}
}

func (repo *productListGetter) GetListProduct(ctx context.Context, in valuetype.GetProductListIn) (res []entity.Product, count int, err error) {
	queryBuilder := sqlb.Select(`id`, `name`, `description`, `rating`, `created_at`, `updated_at`, `category_id`).From("products").Where("deleted_at IS NULL")

	repo.applyFilter(queryBuilder, in)
	repo.applyPagination(queryBuilder, in)
	repo.applySort(queryBuilder)

	queryString, queryArgs, err := queryBuilder.PlaceholderFormat(sqlb.Dollar).ToSql()
	if err != nil {
		return nil, 0, err
	}

	rows, err := repo.db.QueryxContext(ctx, queryString, queryArgs...)
	if err != nil {
		return nil, 0, err
	}

	defer func() { _ = rows.Close() }()

	for rows.Next() {
		productModel := model.Product{}
		err := rows.Scan(
			&productModel.ID,
			&productModel.Name,
			&productModel.Description,
			&productModel.Rating,
			&productModel.CreatedAt,
			&productModel.UpdatedAt,
			&productModel.CategoryID,
		)
		if err != nil {
			return nil, 0, err
		}

		res = append(res, entity.Product{
			ID:          productModel.ID,
			Name:        productModel.Name,
			Description: productModel.Description.String,
			Rating:      int(productModel.Rating.Int64),
			CreatedAt:   productModel.CreatedAt,
			UpdatedAt:   productModel.UpdatedAt.Time,
			CategoryID:  productModel.CategoryID,
		})
	}

	query := "SELECT count(id) FROM products"
	rowsCount := repo.db.QueryRowContext(ctx, query)
	err = rowsCount.Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}

func (repo *productListGetter) applyFilter(queryBuilder *sqlb.SelectBuilder, filter valuetype.GetProductListIn) {
	if filter.Search != "" {
		queryBuilder = queryBuilder.Where(sqlb.Or{
			sqlb.Expr("description LIKE ?", "%"+filter.Search+"%"),
			sqlb.Expr("name LIKE ?", "%"+filter.Search+"%"),
		})
	}
}

func (repo *productListGetter) applyPagination(queryBuilder *sqlb.SelectBuilder, offsetPagination valuetype.GetProductListIn) {
	if offsetPagination.Offset != 0 {
		queryBuilder = queryBuilder.Offset(uint64(offsetPagination.Offset))
	}
	if offsetPagination.Limit != 0 {
		queryBuilder = queryBuilder.Limit(uint64(offsetPagination.Limit))
	}
}

func (repo *productListGetter) applySort(queryBuilder *sqlb.SelectBuilder) {
	// NOTE: for now only by created_at (ADD MORE LATER)
	queryBuilder = queryBuilder.OrderBy("created_at")
}
