package domain

import (
	"net/http"

	"github.com/willian2s/clean-go/core/dto"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// ProductService is a contract of http adapter layer
type ProductService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

// ProductUseCase is a contract of business rule layer
type ProductUseCase interface {
	Create(productRequest *dto.CreateProductRequest) (*Product, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Product], error)
}

// ProductRepository is a contract of database connection adapter layer
type ProductRepository interface {
	Create(productRequest *dto.CreateProductRequest) (*Product, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Product], error)
}
