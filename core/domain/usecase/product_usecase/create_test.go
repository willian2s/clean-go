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

func TestCreate(t *testing.T) {
	fakeRequestProduct := dto.CreateProductRequest{}
	fakeDBProduct := domain.Product{}
	faker.FakeData(&fakeRequestProduct)
	faker.FakeData(&fakeDBProduct)

	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockController)
	mockProductRepository.EXPECT().Create(&fakeRequestProduct).Return(&fakeDBProduct, nil)

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Create(&fakeRequestProduct)

	require.Nil(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, fakeDBProduct.Name)
	require.Equal(t, product.Price, fakeDBProduct.Price)
	require.Equal(t, product.Description, fakeDBProduct.Description)
}

func TestCreate_Error(t *testing.T) {
	fakeRequestProduct := dto.CreateProductRequest{}
	faker.FakeData(&fakeRequestProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Create(&fakeRequestProduct).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Create(&fakeRequestProduct)

	require.NotNil(t, err)
	require.Nil(t, product)
}
