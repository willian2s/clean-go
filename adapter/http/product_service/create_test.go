package productservice_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	productservice "github.com/willian2s/clean-go/adapter/http/product_service"
	"github.com/willian2s/clean-go/core/domain"
	"github.com/willian2s/clean-go/core/domain/mocks"
	"github.com/willian2s/clean-go/core/dto"
	"go.uber.org/mock/gomock"
)

func setupCreate(t *testing.T) (dto.CreateProductRequest, domain.Product, *gomock.Controller) {
	fakeProductRequest := dto.CreateProductRequest{}
	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProductRequest)
	faker.FakeData(&fakeProduct)

	mockController := gomock.NewController(t)

	return fakeProductRequest, fakeProduct, mockController
}

func TestCreate(t *testing.T) {
	fakeProductRequest, fakeProduct, mock := setupCreate(t)
	defer mock.Finish()

	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Create(&fakeProductRequest).Return(&fakeProduct, nil)

	sut := productservice.New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestCreate_JsonErrorFormater(t *testing.T) {
	_, _, mock := setupCreate(t)
	defer mock.Finish()

	mockProductUseCase := mocks.NewMockProductUseCase(mock)

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader("{"))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestCreate_PorductError(t *testing.T) {
	fakeProductRequest, _, mock := setupCreate(t)
	defer mock.Finish()

	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Create(&fakeProductRequest).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productservice.New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}
