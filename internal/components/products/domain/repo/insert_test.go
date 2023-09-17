package repo_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/entity"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"
	varietyValueType "github.com/ikhwankhaliddd/product-service/internal/components/variety/valuetype"
	"github.com/stretchr/testify/assert"

	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestInsertProduct(t *testing.T) {
	ctx := context.Background()
	db, mock, _ := sqlmock.Newx()
	type args struct {
		expQuery string
		data     valuetype.CreateProductIn
	}

	type mockData struct {
		wantRes uint64
		wantErr error
	}

	repo := repo.NewInsertProduct(db)

	tests := []struct {
		name     string
		args     args
		mockData mockData
		doMock   func(args args, mockData mockData)
	}{
		{
			name: "it_success_create_product",
			args: args{
				expQuery: `INSERT INTO "products" (
					"name",
					"category_id",
					"description",
					"created_at"
				) VALUES (
					$1,
					$2,
					$3,
					NOW()
				) RETURNING "products"."id"`,
				data: valuetype.CreateProductIn{
					Name:        "test product",
					Description: "test product description",
					CategoryID:  1,
					Variety: []varietyValueType.CreateVarietyIn{
						{
							Name:  "test product variety",
							Price: 1000.0,
							Stock: 10,
							Image: "product.jpg",
						},
					},
				},
			},
			mockData: mockData{
				wantRes: 1,
			},
			doMock: func(args args, mockData mockData) {
				rows := mock.NewRows([]string{
					"id",
				}).AddRow(1)

				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WithArgs(
					args.data.Name, args.data.CategoryID, args.data.Description).
					WillReturnRows(rows)
			},
		},
		{
			name: "it_return_error_when_create_product",
			args: args{
				expQuery: `INSERT INTO "products" (
					"name",
					"category_id",
					"description",
					"created_at"
				) VALUES (
					$1,
					$2,
					$3,
					NOW()
				) RETURNING "products"."id"`,
				data: valuetype.CreateProductIn{
					Name:        "test product",
					Description: "test product description",
					Variety: []varietyValueType.CreateVarietyIn{
						{
							Name:  "test product variety",
							Price: 1000.0,
							Stock: 10,
							Image: "product.jpg",
						},
					},
				},
			},
			mockData: mockData{
				wantErr: entity.ErrMockProducts,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WithArgs(
					args.data.Name, args.data.CategoryID, args.data.Description).
					WillReturnError(entity.ErrMockProducts)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.mockData)
			}

			res, err := repo.Insert(ctx, tt.args.data)
			assert.Equal(t, tt.mockData.wantRes, res.ID)
			assert.Equal(t, tt.mockData.wantErr, err)
		})
	}

}
