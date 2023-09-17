package repo_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/entity"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestDeleteProduct(t *testing.T) {
	ctx := context.Background()
	db, mock, _ := sqlmock.Newx()

	type args struct {
		expQuery string
		id       uint64
	}

	type mockData struct {
		wantErr error
	}

	repo := repo.NewDeleteProduct(db)

	tests := []struct {
		name     string
		args     args
		mockData mockData
		doMock   func(args args, mockData mockData)
	}{
		{
			name: "it_success_delete_product",
			args: args{
				expQuery: `UPDATE products SET deleted_at = NOW() WHERE id = $1`,
				id:       1,
			},
			mockData: mockData{
				wantErr: nil,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectExec(regexp.QuoteMeta(args.expQuery)).WithArgs(args.id).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "it_error_when_delete_product",
			args: args{
				expQuery: `UPDATE products SET deleted_at = NOW() WHERE id = $1`,
				id:       1,
			},
			mockData: mockData{
				wantErr: entity.ErrMockProducts,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectExec(regexp.QuoteMeta(args.expQuery)).WithArgs(args.id).WillReturnError(entity.ErrMockProducts)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.mockData)
			}

			err := repo.DeleteProduct(ctx, tt.args.id)

			assert.Equal(t, tt.mockData.wantErr, err)
		})
	}
}
