package productrepository

import (
	"context"

	"github.com/google/uuid"
	"github.com/willian2s/clean-go/core/domain"
	"github.com/willian2s/clean-go/core/dto"
)

// Create inserts a new product into the database.
//
// It takes a productRequest pointer as a parameter and returns a product pointer and an error.
func (repository repository) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	ctx := context.Background()
	product := domain.Product{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO products (id, name, price, description) VALUES ($1, $2, $3, $4) RETURNING *",
		uuid.New().String(),
		productRequest.Name,
		productRequest.Price,
		productRequest.Description,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Description,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
