package repo_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/entity"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestGetProduct(t *testing.T) {
	ctx := context.Background()
	db, mock, _ := sqlmock.Newx()

	type args struct {
		expQuery string
		id       uint64
	}

	type mockData struct {
		wantRes entity.Product
		wantErr error
	}

	repo := repo.NewProductGetter(db)

	tests := []struct {
		name     string
		args     args
		mockData mockData
		doMock   func(args args, mockData mockData)
	}{
		{
			name: "it_success_get_product",
			args: args{
				expQuery: `
				SELECT id, name, rating, description, created_at, updated_at
				FROM products
				WHERE id = $1 AND deleted_at IS NULL
				`,
				id: 1,
			},
			mockData: mockData{
				wantRes: entity.Product{
					ID:          1,
					Name:        "test product 1",
					Description: "ini description",
					Rating:      5,
					CreatedAt:   time.Time{},
					UpdatedAt:   time.Time{},
				},
			},
			doMock: func(args args, mockData mockData) {
				rows := mock.NewRows([]string{
					"id", "name", "rating", "description", "created_at", "updated_at",
				}).AddRow(1, "test product 1", 5, "ini description", time.Time{}, time.Time{})

				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WithArgs(args.id).WillReturnRows(rows)
			},
		},
		{
			name: "it_return_error_when_get_category",
			args: args{
				expQuery: `
				SELECT id, name, rating, description, created_at, updated_at
				FROM products
				WHERE id = $1 AND deleted_at IS NULL
				`,
				id: 1,
			},
			mockData: mockData{
				wantErr: entity.ErrMockProducts,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WithArgs(args.id).WillReturnError(entity.ErrMockProducts)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.mockData)
			}

			product, err := repo.GetByID(ctx, tt.args.id)

			assert.Equal(t, tt.mockData.wantRes, product)
			assert.Equal(t, tt.mockData.wantErr, err)
		})
	}
}
