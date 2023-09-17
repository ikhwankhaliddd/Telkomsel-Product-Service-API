package valuetype

import varietyValueType "github.com/ikhwankhaliddd/product-service/internal/components/variety/valuetype"

type CreateProductIn struct {
	Name        string
	Description string
	CategoryID  uint64
	Variety     []varietyValueType.CreateVarietyIn
}

type GetProductListIn struct {
	Offset int
	Limit  int
	Search string
}

type UpdateProductIn struct {
	Name        string
	Description string
	CategoryID  uint64
	Variety     []varietyValueType.CreateVarietyIn
}

type PostRatingIn struct {
	ID     uint64
	Rating int
}
