package publicrepo

import (
	"context"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety/domain/model"
	"github.com/ikhwankhaliddd/product-service/internal/components/variety/valuetype"
	"github.com/jmoiron/sqlx"
)

type IInsertVariety interface {
	Insert(ctx context.Context, productID uint64, in []valuetype.CreateVarietyIn) (err error)
}

type insertVariety struct {
	db *sqlx.DB
}

func NewInsertVariety(db *sqlx.DB) *insertVariety {
	return &insertVariety{db: db}
}

func (repo *insertVariety) Insert(ctx context.Context, productID uint64, in []valuetype.CreateVarietyIn) (err error) {

	varietiesModel := []model.Variety{}
	for _, val := range in {
		varietyModel := model.Variety{
			ProductID: productID,
			Name:      val.Name,
			Price:     val.Price,
			Stock:     val.Stock,
		}
		varietiesModel = append(varietiesModel, varietyModel)
	}

	query := `INSERT INTO variety (
		product_id,
		name,
		price,
		stock,
		created_at
	) VALUES(
		$1,
		$2,
		$3,
		$4,
		NOW()
	)`

	for _, record := range varietiesModel {
		_, err := repo.db.ExecContext(ctx, query, record.ProductID, record.Name, record.Price, record.Stock)
		if err != nil {
			return err
		}
	}
	return nil
}
