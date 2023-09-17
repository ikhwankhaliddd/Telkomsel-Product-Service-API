package publicrepo_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety/entity"
	publicrepo "github.com/ikhwankhaliddd/product-service/internal/components/variety/public_repo"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestGetVariety(t *testing.T) {
	ctx := context.Background()

	db, mock, _ := sqlmock.Newx()

	type args struct {
		expQuery  string
		productID uint64
	}

	type mockData struct {
		wantRes []entity.Variety
		wantErr error
	}

	repo := publicrepo.NewVarietyGetter(db)
	tests := []struct {
		name     string
		args     args
		mockData mockData
		doMock   func(args args, mocmockData mockData)
	}{
		{
			name: "it_success_get_variety",
			args: args{
				expQuery: `
				SELECT id, name, price, stock, image, created_at
				FROM variety WHERE product_id = $1 
				`,
				productID: 3,
			},
			mockData: mockData{
				wantRes: []entity.Variety{
					{
						ID:        1,
						Name:      "variety 1",
						ProductID: 3,
						Price:     100000,
						Stock:     1,
						Image:     "image.jpg",
						CreatedAt: time.Time{},
					},
				},
			},
			doMock: func(args args, mocmockData mockData) {
				rows := mock.NewRows([]string{
					"id", "name", "price", "stock", "image", "created_at",
				}).AddRow(1, "variety 1", 100000, 1, "image.jpg", time.Time{})

				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WithArgs(args.productID).WillReturnRows(rows)
			},
		},
		{
			name: "it_return_error_when_get_variety",
			args: args{
				expQuery: `
				SELECT id, name, price, stock, image, created_at
				FROM variety WHERE product_id = $1 
				`,
				productID: 1,
			},
			mockData: mockData{
				wantErr: entity.ErrMockVariety,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WithArgs(args.productID).WillReturnError(entity.ErrMockVariety)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.mockData)
			}

			variety, err := repo.GetByProductID(ctx, tt.args.productID)

			assert.Equal(t, tt.mockData.wantRes, variety)
			assert.Equal(t, tt.mockData.wantErr, err)
		})
	}
}
