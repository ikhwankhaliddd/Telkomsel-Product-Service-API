package model

import (
	"database/sql"
	"time"
)

type Variety struct {
	ID        uint64
	ProductID uint64
	Name      string
	Price     float64
	Stock     int
	Image     sql.NullString
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
