package products

import (
	"context"

	categoryRepo "github.com/ikhwankhaliddd/product-service/internal/components/category/public_repo"
	categoryValueType "github.com/ikhwankhaliddd/product-service/internal/components/category/valuetype"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/entity"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"
	varietyRepo "github.com/ikhwankhaliddd/product-service/internal/components/variety/public_repo"
)

type ICreateProduct interface {
	CreateProduct(ctx context.Context, in valuetype.CreateProductIn, categoryIn categoryValueType.GetCategoryIn) (res entity.Product, err error)
}

type productCreator struct {
	repo         repo.IInsertProduct
	categoryRepo categoryRepo.IGetCategory
	varietyRepo  varietyRepo.IInsertVariety
}

func NewProductCreator(repo repo.IInsertProduct, categoryRepo categoryRepo.IGetCategory, varietyRepo varietyRepo.IInsertVariety) *productCreator {
	return &productCreator{
		repo:         repo,
		categoryRepo: categoryRepo,
		varietyRepo:  varietyRepo,
	}
}

func (creator *productCreator) CreateProduct(ctx context.Context, in valuetype.CreateProductIn, categoryIn categoryValueType.GetCategoryIn) (res entity.Product, err error) {

	category, err := creator.categoryRepo.GetByName(ctx, categoryIn)
	if err != nil {
		return res, err
	}

	input := valuetype.CreateProductIn{
		Name:        in.Name,
		Description: in.Description,
		CategoryID:  category.ID,
	}

	varietyInput := in.Variety

	product, err := creator.repo.Insert(ctx, input)
	if err != nil {
		return res, err
	}

	err = creator.varietyRepo.Insert(ctx, product.ID, varietyInput)
	if err != nil {
		return
	}

	res.ID = product.ID

	return res, nil
}
