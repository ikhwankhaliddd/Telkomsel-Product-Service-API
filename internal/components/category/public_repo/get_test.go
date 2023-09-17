package publicrepo_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/entity"
	publicrepo "github.com/ikhwankhaliddd/product-service/internal/components/category/public_repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/valuetype"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestGetCategory(t *testing.T) {
	ctx := context.Background()
	db, mock, _ := sqlmock.Newx()

	type args struct {
		expQuery string
		data     valuetype.GetCategoryIn
	}

	type mockData struct {
		wantRes uint64
		wantErr error
	}

	repo := publicrepo.NewCategoryGetter(db)

	tests := []struct {
		name     string
		args     args
		mockData mockData
		doMock   func(args args, mockData mockData)
	}{
		{
			name: "it_success_get_category",
			args: args{
				expQuery: `SELECT id FROM category WHERE name = $1`,
				data: valuetype.GetCategoryIn{
					Name: "Makanan",
				},
			},
			mockData: mockData{
				wantRes: 1,
			},
			doMock: func(args args, mockData mockData) {
				rows := mock.NewRows([]string{
					"id",
				}).AddRow(1)

				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WithArgs(args.data.Name).WillReturnRows(rows)
			},
		},
		{
			name: "it_return_error_when_get_category",
			args: args{
				expQuery: `SELECT id FROM category WHERE name = $1`,
				data: valuetype.GetCategoryIn{
					Name: "",
				},
			},
			mockData: mockData{
				wantErr: entity.ErrMockCategory,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WithArgs(args.data.Name).WillReturnError(entity.ErrMockCategory)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.mockData)
			}

			category, err := repo.GetByName(ctx, tt.args.data)

			assert.Equal(t, tt.mockData.wantRes, category.ID)
			assert.Equal(t, tt.mockData.wantErr, err)
		})
	}
}
