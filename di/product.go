package di

import (
	productservice "github.com/willian2s/clean-go/adapter/http/product_service"
	"github.com/willian2s/clean-go/adapter/postgres"
	productrepository "github.com/willian2s/clean-go/adapter/postgres/product_repository"
	"github.com/willian2s/clean-go/core/domain"
	productusecase "github.com/willian2s/clean-go/core/domain/usecase/product_usecase"
)

// ConfigProductDI configures and returns a ProductService instance.
//
// conn: The connection pool interface to the PostgreSQL database.
// returns: A ProductService instance.
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	prodictRepository := productrepository.New(conn)
	productUseCase := productusecase.New(prodictRepository)
	productService := productservice.New(productUseCase)

	return productService
}
