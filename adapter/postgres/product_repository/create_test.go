package productrepository_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
	productrepository "github.com/willian2s/clean-go/adapter/postgres/product_repository"
	"github.com/willian2s/clean-go/core/domain"
	"github.com/willian2s/clean-go/core/dto"
)

func setupCreate() ([]string, dto.CreateProductRequest, domain.Product, pgxmock.PgxPoolIface) {
	cols := []string{"id", "name", "price", "description"}
	fakeProductRequest := dto.CreateProductRequest{}
	fakeProductDBResponse := domain.Product{}
	faker.FakeData(&fakeProductRequest)
	faker.FakeData(&fakeProductDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakeProductRequest, fakeProductDBResponse, mock
}

func TestCreate(t *testing.T) {
	cols, fakeProductRequest, fakeProductDBResponse, mock := setupCreate()
	defer mock.Close()

	mock.ExpectQuery("INSERT INTO products (.+)").WithArgs(
		fakeProductRequest.Name,
		fakeProductRequest.Price,
		fakeProductRequest.Description,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeProductDBResponse.ID,
		fakeProductDBResponse.Name,
		fakeProductDBResponse.Price,
		fakeProductDBResponse.Description,
	))

	sut := productrepository.New(mock)
	product, err := sut.Create(&fakeProductRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, fakeProductDBResponse.Name)
	require.Equal(t, product.Price, fakeProductDBResponse.Price)
	require.Equal(t, product.Description, fakeProductDBResponse.Description)
}

func TestCreate_DBError(t *testing.T) {
	_, fakeProductRequest, _, mock := setupCreate()
	defer mock.Close()

	mock.ExpectQuery("INSERT INTO products (.+)").WithArgs(
		fakeProductRequest.Name,
		fakeProductRequest.Price,
		fakeProductRequest.Description,
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := productrepository.New(mock)
	product, err := sut.Create(&fakeProductRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.NotNil(t, err)
	require.Nil(t, product)
}
