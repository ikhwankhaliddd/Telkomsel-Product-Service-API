package repo_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/entity"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestTestGetCategoryList(t *testing.T) {
	ctx := context.Background()

	db, mock, _ := sqlmock.Newx()

	type args struct {
		expQuery string
	}

	type mockData struct {
		wantRes []entity.Category
		wantErr error
	}

	repo := repo.NewCategoryListGetter(db)

	tests := []struct {
		name     string
		args     args
		mockData mockData
		doMock   func(args args, mockData mockData)
	}{
		{
			name: "it_success_return_category_list",
			args: args{
				expQuery: `SELECT id, name FROM categories WHERE deleted_at IS NULL`,
			},
			mockData: mockData{
				wantRes: []entity.Category{
					{
						ID:   1,
						Name: "Makanan",
					},
					{
						ID:   2,
						Name: "Elektronik",
					},
				},
			},
			doMock: func(args args, mockData mockData) {
				rows := mock.NewRows([]string{
					"id", "name",
				}).AddRow(1, "Makanan").AddRow(2, "Elektronik")

				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WillReturnRows(rows)
			},
		},
		{
			name: "it_return_error_when_get_category_list",
			args: args{
				expQuery: `SELECT id, name FROM categories WHERE deleted_at IS NULL`,
			},
			mockData: mockData{
				wantErr: entity.ErrMockCategory,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WillReturnError(entity.ErrMockCategory)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.mockData)
			}

			result, err := repo.GetCategoryList(ctx)
			assert.Equal(t, tt.mockData.wantRes, result)
			assert.Equal(t, tt.mockData.wantErr, err)
		})
	}
}
