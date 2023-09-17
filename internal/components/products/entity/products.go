package entity

import (
	"time"

	varietyEntity "github.com/ikhwankhaliddd/product-service/internal/components/variety/entity"
)

type Product struct {
	ID          uint64
	CategoryID  uint64
	Name        string
	Description string
	Varieties   []varietyEntity.Variety
	TotalStock  int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
