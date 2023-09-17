package repo_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/entity"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestInsertRating(t *testing.T) {
	ctx := context.Background()
	db, mock, _ := sqlmock.Newx()

	type args struct {
		expQuery string
		data     valuetype.PostRatingIn
	}

	type mockData struct {
		wantErr error
	}

	repo := repo.NewInsertRating(db)
	tests := []struct {
		name     string
		args     args
		mockData mockData
		doMock   func(args args, mockData mockData)
	}{
		{
			name: "it_success_insert_rating",
			args: args{
				expQuery: `UPDATE products SET rating = $1 WHERE id = $2`,
				data: valuetype.PostRatingIn{
					Rating: 5,
					ID:     4,
				},
			},
			mockData: mockData{
				wantErr: nil,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectExec(regexp.QuoteMeta(args.expQuery)).WithArgs(args.data.Rating, args.data.ID).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "it_return_error_when_insert_rating",
			args: args{
				expQuery: `UPDATE products SET rating = $1 WHERE id = $2`,
				data: valuetype.PostRatingIn{
					Rating: 5,
					ID:     4,
				},
			},
			mockData: mockData{
				wantErr: entity.ErrMockProducts,
			},
			doMock: func(args args, mockData mockData) {
				mock.ExpectExec(regexp.QuoteMeta(args.expQuery)).WithArgs(args.data.Rating, args.data.ID).WillReturnError(entity.ErrMockProducts)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.mockData)
			}

			err := repo.InsertRating(ctx, tt.args.data)

			assert.Equal(t, tt.mockData.wantErr, err)
		})
	}
}
