package model

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          uint64
	CategoryID  uint64
	Name        string
	Description sql.NullString
	Rating      sql.NullInt64
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
	DeletedAt   sql.NullTime
}
