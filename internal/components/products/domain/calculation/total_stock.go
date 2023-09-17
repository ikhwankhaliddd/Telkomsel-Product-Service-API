package calculation

import (
	"context"

	varietyRepo "github.com/ikhwankhaliddd/product-service/internal/components/variety/public_repo"
)

type ICalcuateTotalStock interface {
	CalculateTotalStock(ctx context.Context, id uint64) (result int, err error)
}

type totalStockCalculator struct {
	varietyRepo varietyRepo.IGetVariety
}

func NewTotalStockCalculator(varietyRepo varietyRepo.IGetVariety) *totalStockCalculator {
	return &totalStockCalculator{
		varietyRepo: varietyRepo,
	}
}

func (calc *totalStockCalculator) CalculateTotalStock(ctx context.Context, id uint64) (result int, err error) {
	variety, err := calc.varietyRepo.GetByProductID(ctx, id)
	if err != nil {
		return result, err
	}

	result = 0

	for _, val := range variety {
		result += val.Stock
	}

	return result, nil
}
