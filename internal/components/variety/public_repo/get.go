package publicrepo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/variety/entity"
	"github.com/jmoiron/sqlx"
)

type IGetVariety interface {
	GetByProductID(ctx context.Context, productID uint64) (res []entity.Variety, err error)
}

type varietyGetter struct {
	db *sqlx.DB
}

func NewVarietyGetter(db *sqlx.DB) *varietyGetter {
	return &varietyGetter{
		db: db,
	}
}

func (repo *varietyGetter) GetByProductID(ctx context.Context, productID uint64) (res []entity.Variety, err error) {
	varietyModelList := []model.Variety{}

	query := `
	SELECT id, name, price, stock, image, created_at
	FROM variety WHERE product_id = $1 
	`

	rows, err := repo.db.QueryxContext(ctx, query, productID)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		varietyModel := model.Variety{}

		err := rows.Scan(
			&varietyModel.ID,
			&varietyModel.Name,
			&varietyModel.Price,
			&varietyModel.Stock,
			&varietyModel.Image,
			&varietyModel.CreatedAt,
		)

		if err != nil {
			return res, err
		}

		varietyModelList = append(varietyModelList, varietyModel)
	}

	for _, val := range varietyModelList {
		res = append(res, entity.Variety{
			ID:        val.ID,
			ProductID: productID,
			Name:      val.Name,
			Price:     val.Price,
			Stock:     val.Stock,
			Image:     val.Image.String,
			CreatedAt: val.CreatedAt,
		})
	}

	return res, nil
}
