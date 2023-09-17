package repo_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/ikhwankhaliddd/product-service/internal/components/category/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/entity"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/valuetype"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestInsertCategory(t *testing.T) {
	ctx := context.Background()

	db, mock, _ := sqlmock.Newx()

	type args struct {
		expQuery string
		data     valuetype.InsertCategoryIn
	}

	type mockData struct {
		wantErr error
	}

	repo := repo.NewInsertCategory(db)

	tests := []struct {
		name     string
		args     args
		mockData mockData
		doMock   func(args args, mockData mockData)
	}{
		{
			name: "it_success_insert_category",
			args: args{
				expQuery: `INSERT INTO categories (name, created_at) VALUES($1, NOW())`,
				data: valuetype.InsertCategoryIn{
					Name: "Minuman",
				},
			},
			mockData: mockData{
				wantErr: nil,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectExec(regexp.QuoteMeta(args.expQuery)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "it_return_error_when_insert_category",
			args: args{
				expQuery: `INSERT INTO categories (name, created_at) VALUES($1, NOW())`,
				data: valuetype.InsertCategoryIn{
					Name: "Minuman",
				},
			},
			mockData: mockData{
				wantErr: entity.ErrMockCategory,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectExec(regexp.QuoteMeta(args.expQuery)).WillReturnError(entity.ErrMockCategory)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.mockData)
			}

			err := repo.Insert(ctx, tt.args.data)

			assert.Equal(t, tt.mockData.wantErr, err)
		})
	}
}
