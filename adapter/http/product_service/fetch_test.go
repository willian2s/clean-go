package productservice_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	productservice "github.com/willian2s/clean-go/adapter/http/product_service"
	"github.com/willian2s/clean-go/core/domain"
	"github.com/willian2s/clean-go/core/domain/mocks"
	"github.com/willian2s/clean-go/core/dto"
	"go.uber.org/mock/gomock"
)

func setupFetch(t *testing.T) (dto.PaginationRequestParams, domain.Product, *gomock.Controller) {
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         []string{""},
		Descending:   []string{""},
		Search:       "",
	}
	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProduct)

	mockCtrl := gomock.NewController(t)

	return fakePaginationRequestParams, fakeProduct, mockCtrl
}

func TestFetch(t *testing.T) {
	fakePaginationRequestParams, fakeProduct, mock := setupFetch(t)
	defer mock.Finish()

	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Fetch(&fakePaginationRequestParams).Return(&domain.Pagination{
		Items: []domain.Product{fakeProduct},
		Total: 1,
	}, nil)

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product", nil)
	r.Header.Set("Content-Type", "application/json")

	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")

	r.URL.RawQuery = queryStringParams.Encode()
	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestFetch_PorductError(t *testing.T) {
	fakePaginationRequestParams, _, mock := setupFetch(t)
	defer mock.Finish()

	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Fetch(&fakePaginationRequestParams).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product", nil)
	r.Header.Set("Content-Type", "application/json")

	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")

	r.URL.RawQuery = queryStringParams.Encode()
	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
