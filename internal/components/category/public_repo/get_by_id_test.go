package publicrepo_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/entity"
	publicrepo "github.com/ikhwankhaliddd/product-service/internal/components/category/public_repo"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestGetCategoryNameByID(t *testing.T) {
	ctx := context.Background()
	db, mock, _ := sqlmock.Newx()

	type args struct {
		expQuery string
		id       uint64
	}

	type mockData struct {
		wantRes entity.Category
		wantErr error
	}

	repo := publicrepo.NewCategoryByIDGetter(db)

	tests := []struct {
		name     string
		args     args
		mockData mockData
		doMock   func(args args, mockData mockData)
	}{
		{
			name: "it_success_get_category_name",
			args: args{
				expQuery: `SELECT name FROM categories WHERE id = $1`,
				id:       1,
			},
			mockData: mockData{
				wantRes: entity.Category{
					Name: "Makanan",
				},
			},
			doMock: func(args args, mockData mockData) {
				rows := mock.NewRows([]string{
					"name",
				}).AddRow("Makanan")

				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WithArgs(args.id).WillReturnRows(rows)
			},
		},
		{
			name: "it_return_error_when_get_category_name",
			args: args{
				expQuery: `SELECT name FROM categories WHERE id = $1`,
				id:       1,
			},
			mockData: mockData{
				wantErr: entity.ErrMockCategory,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectQuery(regexp.QuoteMeta(args.expQuery)).WithArgs(args.id).WillReturnError(entity.ErrMockCategory)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.mockData)
			}

			category, err := repo.GetByID(ctx, tt.args.id)

			assert.Equal(t, tt.mockData.wantRes, category)
			assert.Equal(t, tt.mockData.wantErr, err)
		})
	}
}
