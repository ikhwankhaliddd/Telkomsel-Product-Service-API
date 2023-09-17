package entity

import "time"

type Variety struct {
	ID        uint64
	ProductID uint64
	Name      string
	Price     float64
	Stock     int
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
