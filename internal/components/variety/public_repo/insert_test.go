package publicrepo_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety/entity"
	publicrepo "github.com/ikhwankhaliddd/product-service/internal/components/variety/public_repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/variety/valuetype"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestInsertVariety(t *testing.T) {
	ctx := context.Background()
	db, mock, _ := sqlmock.Newx()

	type args struct {
		expQuery  string
		productID uint64
		data      []valuetype.CreateVarietyIn
	}

	type mockData struct {
		wantErr error
	}

	repo := publicrepo.NewInsertVariety(db)

	tests := []struct {
		name     string
		args     args
		mockData mockData
		doMock   func(args args, mocmockData mockData)
	}{
		{
			name: "it_success_create_variety",
			args: args{
				expQuery: `INSERT INTO variety (
					product_id,
					name,
					price,
					stock,
					created_at
				) VALUES(
					$1,
					$2,
					$3,
					$4,
					NOW()
				)`,
				productID: 1,
				data: []valuetype.CreateVarietyIn{
					{
						Name:  "test variety 1",
						Price: 10.0,
						Stock: 50,
					},
				},
			},
			mockData: mockData{
				wantErr: nil,
			},
			doMock: func(args args, mockData mockData) {
				for _, val := range args.data {
					mock.ExpectExec(regexp.QuoteMeta(args.expQuery)).WithArgs(
						args.productID, val.Name, val.Price, val.Stock,
					).WillReturnResult(sqlmock.NewResult(1, 1))
				}
			},
		},
		{
			name: "it_will_return_error_when_create_variety",
			args: args{
				expQuery: `INSERT INTO variety (
					product_id,
					name,
					price,
					stock,
					created_at
				) VALUES(
					$1,
					$2,
					$3,
					$4,
					NOW()
				)`,
				productID: 1,
				data: []valuetype.CreateVarietyIn{
					{
						Name:  "test variety 1",
						Price: 10.0,
						Stock: 50,
					},
				},
			},
			mockData: mockData{
				wantErr: entity.ErrMockVariety,
			},
			doMock: func(args args, mockData mockData) {
				for _, val := range args.data {
					mock.ExpectExec(regexp.QuoteMeta(args.expQuery)).WithArgs(
						args.productID, val.Name, val.Price, val.Stock,
					).WillReturnError(entity.ErrMockVariety)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args, tt.mockData)
			}

			err := repo.Insert(ctx, tt.args.productID, tt.args.data)
			assert.Equal(t, tt.mockData.wantErr, err)
		})
	}

}
