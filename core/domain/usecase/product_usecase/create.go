package productusecase

import (
	"github.com/willian2s/clean-go/core/domain"
	"github.com/willian2s/clean-go/core/dto"
)

// Create creates a new product.
//
// It takes a pointer to a CreateProductRequest struct as a parameter and returns a pointer to a Product struct and an error.
func (uc usecase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product, err := uc.repository.Create(productRequest)

	if err != nil {
		return nil, err
	}

	return product, nil
}
