package productusecase_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
	"github.com/willian2s/clean-go/core/domain"
	"github.com/willian2s/clean-go/core/domain/mocks"
	productusecase "github.com/willian2s/clean-go/core/domain/usecase/product_usecase"
	"github.com/willian2s/clean-go/core/dto"
	"go.uber.org/mock/gomock"
)

func TestFetch(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}
	fakeDBProduct := domain.Product{}

	faker.FakeData(&fakeDBProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Fetch(&fakePaginationRequestParams).Return(&domain.Pagination{
		Items: []domain.Product{fakeDBProduct},
		Total: 1,
	}, nil)

	sut := productusecase.New(mockProductRepository)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.Nil(t, err)

	for _, product := range products.Items.([]domain.Product) {
		require.Nil(t, err)
		require.NotEmpty(t, product.ID)
		require.Equal(t, product.Name, fakeDBProduct.Name)
		require.Equal(t, product.Price, fakeDBProduct.Price)
		require.Equal(t, product.Description, fakeDBProduct.Description)
	}
}

func TestFetch_Error(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Fetch(&fakePaginationRequestParams).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)
	require.Nil(t, product)
}
